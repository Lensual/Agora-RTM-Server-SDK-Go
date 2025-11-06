package async_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/async"
	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

var (
	appId  string = os.Getenv("APP_ID")
	userId string = os.Getenv("USER_ID")
	token  string = os.Getenv("TOKEN")
)

var (
	rtmConfig *agrtm.RtmConfig
	rtmClient *agrtm.IRtmClient
)

func TestMain(m *testing.M) {
	rtmConfig = agrtm.NewRtmConfig()
	rtmConfig.AppId = appId
	rtmConfig.UserId = userId
	rtmConfig.EventHandler = async.NewAsyncCallRtmEventHandler() //TODO 需要一个组合事件处理器

	rtmClient = agrtm.NewRtmClient(rtmConfig)
	defer rtmClient.Release()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLogin(t *testing.T) {
	ctx := context.Background()

	// login
	loginCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	_, err := async.Call[int](func() (ret int, reqId uint64) {
		ret, reqId = rtmClient.Login(token)
		log.Printf("Login ret: %d, reqId: %d, token: %s\n", ret, reqId, token)
		return ret, reqId
	}).Await(loginCtx)
	if err != nil {
		panic(err)
	}

	log.Println("login success")

	// logout
	logoutCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	_, err = async.Call[int](func() (ret int, reqId uint64) {
		ret, reqId = rtmClient.Logout()
		log.Printf("Logout ret: %d, reqId: %d\n", ret, reqId)
		return ret, reqId
	}).Await(logoutCtx)
	if err != nil {
		panic(err)
	}

	log.Println("logout success")
}
