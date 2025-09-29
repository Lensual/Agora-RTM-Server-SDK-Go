package main

import (
	"fmt"
	"time"

	agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"
)

type MyRtmEventHandler struct {
	// public member, should be named with a uppercase letter
	RtmClient   *agrtm.IRtmClient
	ChannelName string
	UserId      string
	Sign        chan struct{}
}

func logWithTime(format string, args ...interface{}) {
	fmt.Printf("[%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05.000"),
		fmt.Sprintf(format, args...))
}

func (h *MyRtmEventHandler) IsRtmEventHandler() {}

func (h *MyRtmEventHandler) OnMessageEvent(event *agrtm.MessageEvent) {
	logWithTime("OnMessageEvent: received message event - Channel: %s, Publisher: %s, Message: %s, Message Length: %d",
		event.ChannelName, event.Publisher, event.Message, len(event.Message))

	message := event.Message
	fmt.Printf("messageContent: %v, string:%s\n", message, string(message))
	// simulate a echo server
	if h.RtmClient != nil {
		// Publish(channelName string, message []byte, length uint, option *PublishOptions, requestId *uint64) int {
		requestId := uint64(0)
		ret := int(0)
		// Note1: parse the message, and get properties
		channelType := event.ChannelType
		channelname := event.ChannelName
		publisher := event.Publisher

		pubName := channelname
		// Note2: accoring to the message type, send the message to the channel or user
		if channelType == agrtm.RtmChannelTypeUSER {
			pubName = string(publisher)

			ret, requestId = h.RtmClient.SendUserMessage(pubName, []byte(message))
		} else {
			ret, requestId = h.RtmClient.SendChannelMessage(pubName, []byte(message))
		}
		fmt.Printf("send: %s, ret: %d, requestId: %d\n", pubName, ret, requestId)

		//Note3: send the message to the channel or user like echo server through Publish

		opt := agrtm.NewPublishOptions()
		opt.MessageType = (event.MessageType)
		opt.ChannelType = (channelType)
		opt.CustomType = (event.CustomType)
		ret, requestId = h.RtmClient.Publish(pubName, []byte(message), opt)
		fmt.Printf("publish: %s, ret: %d, requestId: %d\n", pubName, ret, requestId)
	}
}
func (h *MyRtmEventHandler) OnPresenceEvent(event *agrtm.PresenceEvent) {
	logWithTime("OnPresenceEvent event:%v", event)
}
func (h *MyRtmEventHandler) OnTopicEvent(event *agrtm.TopicEvent) {
	logWithTime("OnTopicEvent event:%v", event)
}
func (h *MyRtmEventHandler) OnLockEvent(event *agrtm.LockEvent) {
	logWithTime("OnLockEvent event:%v", event)
}
func (h *MyRtmEventHandler) OnStorageEvent(event *agrtm.StorageEvent) {
	logWithTime("OnStorageEvent event:%v", event)
}
func (h *MyRtmEventHandler) OnJoinResult(requestId uint64, channelName string, userId string, errorCode int) {
	logWithTime("OnJoinResult requestId:%v channelName:%v userId:%v errorCode:%v", requestId, channelName, userId, errorCode)
}
func (h *MyRtmEventHandler) OnLeaveResult(requestId uint64, channelName string, userId string, errorCode int) {
	logWithTime("OnLeaveResult requestId:%v channelName:%v userId:%v errorCode:%v", requestId, channelName, userId, errorCode)
}
func (h *MyRtmEventHandler) OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
	logWithTime("OnJoinTopicResult requestId:%v channelName:%v userId:%v topic:%v meta:%v errorCode:%v", requestId, channelName, userId, topic, meta, errorCode)
}
func (h *MyRtmEventHandler) OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
	logWithTime("OnLeaveTopicResult requestId:%v channelName:%v userId:%v topic:%v meta:%v errorCode:%v", requestId, channelName, userId, topic, meta, errorCode)
}
func (h *MyRtmEventHandler) OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers agrtm.UserList, failedUsers agrtm.UserList, errorCode int) {
	logWithTime("OnSubscribeTopicResult requestId:%v channelName:%v userId:%v topic:%v succeedUsers:%v failedUsers:%v errorCode:%v", requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
}
func (h *MyRtmEventHandler) OnConnectionStateChanged(channelName string, state int, reason int) {
	logWithTime("OnConnectionStateChanged channelName:%s state:%v reason:%v", channelName, state, reason)
}
func (h *MyRtmEventHandler) OnTokenPrivilegeWillExpire(channelName string) {
	logWithTime("OnTokenPrivilegeWillExpire channelName:%v", channelName)
}
func (h *MyRtmEventHandler) OnSubscribeResult(requestId uint64, channelName string, errorCode int) {
	logWithTime("OnSubscribeResult requestId:%v channelName:%v errorCode:%v", requestId, channelName, errorCode)
	if h.Sign != nil {
		h.Sign <- struct{}{}
	}
}
func (h *MyRtmEventHandler) OnPublishResult(requestId uint64, errorCode int) {
	logWithTime("OnPublishResult requestId:%v errorCode:%v", requestId, errorCode)
}
func (h *MyRtmEventHandler) OnLoginResult(requestId uint64, errorCode int) {
	logWithTime("OnLoginResult requestId:%v errorCode:%v", requestId, errorCode)
	if h.Sign != nil {
		h.Sign <- struct{}{}
	}
}
func (h *MyRtmEventHandler) OnSetChannelMetadataResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
	logWithTime("OnSetChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v", requestId, channelName, channelType, errorCode)
}
func (h *MyRtmEventHandler) OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
	logWithTime("OnUpdateChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v", requestId, channelName, channelType, errorCode)
}
func (h *MyRtmEventHandler) OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
	logWithTime("OnRemoveChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v", requestId, channelName, channelType, errorCode)
}
func (h *MyRtmEventHandler) OnGetChannelMetadataResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, data *agrtm.IMetadata, errorCode int) {
	logWithTime("OnGetChannelMetadataResult requestId:%v channelName:%v channelType:%v data:%v errorCode:%v", requestId, channelName, channelType, data, errorCode)
}
func (h *MyRtmEventHandler) OnSetUserMetadataResult(requestId uint64, userId string, errorCode int) {
	logWithTime("OnSetUserMetadataResult requestId:%v userId:%v errorCode:%v", requestId, userId, errorCode)
}
func (h *MyRtmEventHandler) OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode int) {
	logWithTime("OnUpdateUserMetadataResult requestId:%v userId:%v errorCode:%v", requestId, userId, errorCode)
}
func (h *MyRtmEventHandler) OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode int) {
	logWithTime("OnRemoveUserMetadataResult requestId:%v userId:%v errorCode:%v", requestId, userId, errorCode)
}
func (h *MyRtmEventHandler) OnGetUserMetadataResult(requestId uint64, userId string, data *agrtm.IMetadata, errorCode int) {
	logWithTime("OnGetUserMetadataResult requestId:%v userId:%v data:%v errorCode:%v", requestId, userId, data, errorCode)
}
func (h *MyRtmEventHandler) OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode int) {
	logWithTime("OnSubscribeUserMetadataResult requestId:%v userId:%v errorCode:%v", requestId, userId, errorCode)
}
func (h *MyRtmEventHandler) OnSetLockResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
	logWithTime("OnSetLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v", requestId, channelName, channelType, lockName, errorCode)
}
func (h *MyRtmEventHandler) OnRemoveLockResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
	logWithTime("OnRemoveLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v", requestId, channelName, channelType, lockName, errorCode)
}
func (h *MyRtmEventHandler) OnReleaseLockResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
	logWithTime("OnReleaseLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v", requestId, channelName, channelType, lockName, errorCode)
}
func (h *MyRtmEventHandler) OnAcquireLockResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int, errorDetails string) {
	logWithTime("OnAcquireLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v errorDetails:%v", requestId, channelName, channelType, lockName, errorCode, errorDetails)
}
func (h *MyRtmEventHandler) OnRevokeLockResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
	logWithTime("OnRevokeLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v", requestId, channelName, channelType, lockName, errorCode)
}
func (h *MyRtmEventHandler) OnGetLocksResult(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockDetailList *agrtm.LockDetail, count uint, errorCode int) {
	logWithTime("OnGetLocksResult requestId:%v channelName:%v channelType:%v lockDetailList:%v errorCode:%v", requestId, channelName, channelType, lockDetailList, errorCode)
}
func (h *MyRtmEventHandler) OnWhoNowResult(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
	logWithTime("OnWhoNowResult requestId:%v userStateList:%v count:%v nextPage:%v errorCode:%v", requestId, userStateList, count, nextPage, errorCode)
}
func (h *MyRtmEventHandler) OnGetOnlineUsersResult(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
	logWithTime("OnGetOnlineUsersResult requestId:%v userStateList:%v count:%v nextPage:%v errorCode:%v", requestId, userStateList, count, nextPage, errorCode)
}
func (h *MyRtmEventHandler) OnWhereNowResult(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
	logWithTime("OnWhereNowResult requestId:%v channels:%v count:%v errorCode:%v", requestId, channels, count, errorCode)
}
func (h *MyRtmEventHandler) OnGetUserChannelsResult(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
	logWithTime("OnGetUserChannelsResult requestId:%v channels:%v count:%v errorCode:%v", requestId, channels, count, errorCode)
}
func (h *MyRtmEventHandler) OnPresenceSetStateResult(requestId uint64, errorCode int) {
	logWithTime("OnPresenceSetStateResult requestId:%v errorCode:%v", requestId, errorCode)
}
func (h *MyRtmEventHandler) OnPresenceRemoveStateResult(requestId uint64, errorCode int) {
	logWithTime("OnPresenceRemoveStateResult requestId:%v errorCode:%v", requestId, errorCode)
}
func (h *MyRtmEventHandler) OnPresenceGetStateResult(requestId uint64, state *agrtm.UserState, errorCode int) {
	logWithTime("OnPresenceGetStateResult requestId:%v state:%v errorCode:%v", requestId, state, errorCode)
}
func (h *MyRtmEventHandler) OnLinkStateEvent(event *agrtm.LinkStateEvent) {
	logWithTime("OnLinkStateEvent event:%v", event)
}
func (h *MyRtmEventHandler) OnLogoutResult(requestId uint64, errorCode int) {
	logWithTime("OnLogoutResult requestId:%v errorCode:%v", requestId, errorCode)
}
func (h *MyRtmEventHandler) OnRenewTokenResult(requestId uint64, serverType agrtm.RtmServiceType, channelName string, errorCode int) {
	logWithTime("OnRenewTokenResult requestId:%v serverType:%v channelName:%v errorCode:%v", requestId, serverType, channelName, errorCode)
}
func (h *MyRtmEventHandler) OnPublishTopicMessageResult(requestId uint64, channelName string, topic string, errorCode int) {
	logWithTime("OnPublishTopicMessageResult requestId:%v channelName:%v topic:%v errorCode:%v", requestId, channelName, topic, errorCode)
}
func (h *MyRtmEventHandler) OnUnsubscribeTopicResult(requestId uint64, channelName string, topic string, errorCode int) {
	logWithTime("OnUnsubscribeTopicResult requestId:%v channelName:%v topic:%v errorCode:%v", requestId, channelName, topic, errorCode)
}
func (h *MyRtmEventHandler) OnGetSubscribedUserListResult(requestId uint64, channelName string, topic string, user *agrtm.UserList, errorCode int) {
	logWithTime("OnGetSubscribedUserListResult requestId:%v channelName:%v topic:%v user:%v errorCode:%v", requestId, channelName, topic, user, errorCode)
}
func (h *MyRtmEventHandler) OnGetHistoryMessagesResult(requestId uint64, messageList []agrtm.HistoryMessage, newStart uint64, errorCode int) {
	logWithTime("OnGetHistoryMessagesResult requestId:%v messageList:%v newStart:%v errorCode:%v", requestId, messageList, newStart, errorCode)
}
func (h *MyRtmEventHandler) OnUnsubscribeUserMetadataResult(requestId uint64, userId string, errorCode int) {
	logWithTime("OnUnsubscribeUserMetadataResult requestId:%v userId:%v errorCode:%v", requestId, userId, errorCode)
}
