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
	}

	appId = os.Args[1]
	channelName = os.Args[2]
	userId = os.Args[3]

	if lenArgs >= 5 {
		token = os.Args[4]
	} else {
		token = appId
	}
	fmt.Printf("appId: %s, channelName: %s, userId: %s, token: %s\n", appId, channelName, userId, token)

	// 检查参数
	if appId == "" || channelName == "" || userId == "" || token == "" {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	ret := 0
	// 适用于没有token的情况
	rtmEventHandler := agrtm.NewRtmEventHandlerBridge(&MyRtmEventHandler{})
	fmt.Printf("NewRtmEventHandlerBridge: %p\n", rtmEventHandler) //DEBUG
	//defer rtmEventHandler.Delete()

	rtmConfig := agrtm.NewRtmConfig()
	//defer rtmConfig.Delete()
	rtmConfig.SetAppId(appId)
	rtmConfig.SetUserId(userId)
	rtmConfig.SetEventHandler(rtmEventHandler.ToAgoraEventHandler())
	fmt.Printf("NewRtmConfig: %+v\n", rtmConfig) //DEBUG

	rtmClient := agrtm.CreateAgoraRtmClient(rtmConfig)
	fmt.Printf("CreateAgoraRtmClient: %p\n", rtmClient) //DEBUG

	
	



	ret = rtmClient.Login(token)
	fmt.Printf("Login: %d\n", ret)
	if ret != 0 {
		panic(ret)
	}
	time.Sleep(time.Second * 3) //等登录完成，正常情况需要收到OnLoginResult后再Subscribe，这里方便测试用了sleep

	var reqId uint64
	opt := agrtm.NewSubscribeOptions()
	fmt.Printf("NewSubscribeOptions: %p\n", opt) //DEBUG

	ret = rtmClient.Subscribe(channelName, opt, &reqId)
	fmt.Printf("Subscribe: %d\n", ret)
	if ret != 0 {
		panic(ret)
	}

	//阻塞直到有信号传入
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	fmt.Println("启动")

waitSignal:
	for {
		select {
		case signal := <-c:
			if signal == os.Interrupt ||
				signal == os.Kill ||
				signal == syscall.SIGABRT ||
				signal == syscall.SIGTERM {
				fmt.Println("退出信号", signal)
				break waitSignal
			}
		default:
			// errcode = channel.SendMessage(message)
			// fmt.Printf("channel.SendMessage:%v\n", errcode) //DEBUG
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
