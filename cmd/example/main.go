package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

func main() {
	// start pprof
	go func() {
		// listen on all interfaces, if you want to listen on localhost, use http.ListenAndServe("localhost:6060", nil)
		// but local host is not accessible from outside!!
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	// rtm start
	appId := os.Getenv("APPID")
	userId := os.Getenv("USER_ID")
	token := os.Getenv("TOKEN")
	channelName := os.Getenv("CHANNEL_NAME")
	lenArgs := len(os.Args)
	if lenArgs < 4 {
		fmt.Println("Usage: process <appid> <channelname> <userid> <token_optional>")
		os.Exit(1)
		logWithTime("Usage: process <appid> <channelname> <userid> <token_optional>")
	}

	appId = os.Args[1]
	channelName = os.Args[2]
	userId = os.Args[3]

	if lenArgs >= 5 {
		token = os.Args[4]
	}
	logWithTime("appId: %s, channelName: %s, userId: %s, token: %s\n", appId, channelName, userId, token)

	// 检查参数
	if appId == "" || channelName == "" || userId == "" {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	ret := int(0)
	var requestId uint64

	myEventHandler := &MyRtmEventHandler{}

	rtmConfig := agrtm.NewRtmConfig()
	//defer rtmConfig.Delete()
	rtmConfig.AppId = appId
	rtmConfig.UserId = userId
	rtmConfig.EventHandler = myEventHandler

	logConfig := agrtm.NewRtmLogConfig()
	logConfig.FilePath = "./logs/rtm.log"
	logConfig.FileSizeInKB = 1024
	logConfig.Level = agrtm.RtmLogLevelINFO
	rtmConfig.LogConfig = logConfig
	fmt.Printf("NewRtmConfig: %+v\n", rtmConfig) //DEBUG

	rtmClient := agrtm.NewRtmClient(rtmConfig)
	logWithTime("CreateAgoraRtmClient: %p\n", rtmClient) //DEBUG

	// set user channel info to event handler
	sign := make(chan struct{})
	myEventHandler.ChannelName = channelName
	myEventHandler.UserId = userId
	myEventHandler.RtmClient = rtmClient
	myEventHandler.Sign = sign

	logWithTime("Login Start: %d\n", ret)
	if token == "" {
		token = appId
	}
	ret, requestId = rtmClient.Login(token)

	fmt.Printf("Login ret: %d, requestId: %d, token: %s\n", ret, requestId, token)	

	if ret != 0 {
		panic(ret)
	}
	// wait for login result and timedout to 3 seconds
	select {
	case <-sign:
	case <-time.After(time.Second * 3):
		panic("login timeout")
	}
	logWithTime("login success")


	opt := agrtm.NewSubscribeOptions()

	logWithTime("Subscribe start: %d\n", ret)
	ret, requestId = rtmClient.Subscribe(channelName, opt)
	fmt.Printf("Subscribe ret: %d, requestId: %d\n", ret, requestId)

	if ret != 0 {
		panic(ret)
	}
	// wait for subscribe result and timedout to 3 seconds
	select {
	case <-sign:
	case <-time.After(time.Second * 3):
		panic("subscribe timeout")
	}
	logWithTime("subscribe success")

	//阻塞直到有信号传入
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	logWithTime("rtm client start to work")

waitSignal:
	for {
		select {
		case signal := <-c:
			if signal == os.Interrupt ||
				signal == os.Kill ||
				signal == syscall.SIGABRT ||
				signal == syscall.SIGTERM {
				logWithTime("exit signal: %v", signal)
				break waitSignal
			}
		default:
			time.Sleep(time.Second)
		}
	}

	//clean
	rtmClient.Logout()
	// wait for logout
	//	time.Sleep(time.Second * 3)
	//unregister event handler

	//release
	rtmClient.Release()
	rtmClient = nil

	// release myEventHandler
	myEventHandler = nil
	rtmConfig = nil
}
