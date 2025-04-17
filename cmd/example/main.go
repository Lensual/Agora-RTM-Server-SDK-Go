package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/pkg/agora"
)

func main() {
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
	// 适用于没有token的情况

	rtmClient := agrtm.CreateAgoraRtmClient()
	fmt.Printf("CreateAgoraRtmClient: %p\n", rtmClient) //DEBUG

	rtmEventHandler := agrtm.NewRtmEventHandlerBridge(&MyRtmEventHandler{})
	fmt.Printf("NewRtmEventHandlerBridge: %p\n", rtmEventHandler) //DEBUG

	rtmConfig := agrtm.NewRtmConfig()
	defer rtmConfig.Delete()
	rtmConfig.SetAppId(appId)
	rtmConfig.SetUserId(userId)
	rtmConfig.SetEventHandler(rtmEventHandler.ToAgoraEventHandler())
	fmt.Printf("NewRtmConfig: %+v\n", rtmConfig) //DEBUG

	ret := rtmClient.Initialize(rtmConfig)
	fmt.Printf("Initialize: %d\n", ret)
	if ret != 0 {
		panic(ret)
	}

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
	/*
	message.Release()
	errcode = channel.Leave()
	fmt.Printf("channel.Leave:%v\n", errcode) //DEBUG
	channel.Release()
	channel = nil
	channelEventHandler.Delete()
	errcode = rtmService.Logout()
	fmt.Printf("rtmService.Logout:%v\n", errcode) //DEBUG
	rtmService.Release(true)
	rtmService = nil
	*/
}
