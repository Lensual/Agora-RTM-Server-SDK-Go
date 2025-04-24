package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	_ "net/http/pprof"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/pkg/agora"
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
	} else {
		token = appId
	}
	logWithTime("appId: %s, channelName: %s, userId: %s, token: %s\n", appId, channelName, userId, token)

	// 检查参数
	if appId == "" || channelName == "" || userId == "" || token == "" {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	ret := 0
	// 适用于没有token的情况
	myEventHandler := &MyRtmEventHandler{}
	rtmEventHandler := agrtm.NewRtmEventHandlerBridge(myEventHandler)
	fmt.Printf("NewRtmEventHandlerBridge: %p\n", rtmEventHandler) //DEBUG
	//defer rtmEventHandler.Delete()

	rtmConfig := agrtm.NewRtmConfig()
	//defer rtmConfig.Delete()
	rtmConfig.SetAppId(appId)
	rtmConfig.SetUserId(userId)
	rtmConfig.SetEventHandler(rtmEventHandler.ToAgoraEventHandler())
	fmt.Printf("NewRtmConfig: %+v\n", rtmConfig) //DEBUG

	rtmClient := agrtm.CreateAgoraRtmClient(rtmConfig)
	logWithTime("CreateAgoraRtmClient: %p\n", rtmClient) //DEBUG

	// set user channel info to event handler
	sign := make(chan struct{})
	myEventHandler.ChannelName = channelName
	myEventHandler.UserId = userId	
	myEventHandler.RtmClient = rtmClient
	myEventHandler.Sign = sign
	


	logWithTime("Login Start: %d\n", ret)
	ret = rtmClient.Login(token)
	
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

	var reqId uint64
	opt := agrtm.NewSubscribeOptions()
	

	logWithTime("Subscribe start: %d\n", ret)
	ret = rtmClient.Subscribe(channelName, opt, &reqId)

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

	// release rtmEventHandler
	rtmEventHandler.Delete()
	rtmEventHandler = nil
	// release rtmConfig
	rtmConfig.Delete()
	rtmConfig = nil
}
