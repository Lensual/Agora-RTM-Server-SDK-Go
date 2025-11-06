package utils

import agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"

// NewCompositeRtmEventHandler 返回组合的 [agrtm.RtmEventHandler].
//
// 多个 Handler 将会串行执行.
//
// Example:
//
//	rtmConfig := agrtm.NewRtmConfig()
//	rtmConfig.EventHandler = utils.NewCompositeRtmEventHandler(
//		// Handle A
//		// Handle B
//		// ...
//	)
func NewCompositeRtmEventHandler(handles ...*agrtm.RtmEventHandler) *agrtm.RtmEventHandler {
	return &agrtm.RtmEventHandler{
		OnMessageEvent: func(event *agrtm.MessageEvent) {
			for _, h := range handles {
				if h.OnMessageEvent != nil {
					h.OnMessageEvent(event)
				}
			}
		},
		OnPresenceEvent: func(event *agrtm.PresenceEvent) {
			for _, h := range handles {
				if h.OnPresenceEvent != nil {
					h.OnPresenceEvent(event)
				}
			}
		},
		OnTopicEvent: func(event *agrtm.TopicEvent) {
			for _, h := range handles {
				if h.OnTopicEvent != nil {
					h.OnTopicEvent(event)
				}
			}
		},
		OnLockEvent: func(event *agrtm.LockEvent) {
			for _, h := range handles {
				if h.OnLockEvent != nil {
					h.OnLockEvent(event)
				}
			}
		},
		OnStorageEvent: func(event *agrtm.StorageEvent) {
			for _, h := range handles {
				if h.OnStorageEvent != nil {
					h.OnStorageEvent(event)
				}
			}
		},
		OnJoinResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnJoinResult != nil {
					h.OnJoinResult(requestId, channelName, userId, errorCode)
				}
			}
		},
		OnLeaveResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnLeaveResult != nil {
					h.OnLeaveResult(requestId, channelName, userId, errorCode)
				}
			}
		},
		OnJoinTopicResult: func(requestId uint64, channelName, userId, topic, meta string, errorCode int) {
			for _, h := range handles {
				if h.OnJoinTopicResult != nil {
					h.OnJoinTopicResult(requestId, channelName, userId, topic, meta, errorCode)
				}
			}
		},
		OnLeaveTopicResult: func(requestId uint64, channelName, userId, topic, meta string, errorCode int) {
			for _, h := range handles {
				if h.OnLeaveTopicResult != nil {
					h.OnLeaveTopicResult(requestId, channelName, userId, topic, meta, errorCode)
				}
			}
		},
		OnSubscribeTopicResult: func(requestId uint64, channelName, userId, topic string, succeedUsers, failedUsers agrtm.UserList, errorCode int) {
			for _, h := range handles {
				if h.OnSubscribeTopicResult != nil {
					h.OnSubscribeTopicResult(requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
				}
			}
		},
		// Deprecated: use OnLinkStateEvent instead.
		OnConnectionStateChanged: func(channelName string, state, reason int) {
			for _, h := range handles {
				if h.OnConnectionStateChanged != nil {
					h.OnConnectionStateChanged(channelName, state, reason)
				}
			}
		},
		OnTokenPrivilegeWillExpire: func(channelName string) {
			for _, h := range handles {
				if h.OnTokenPrivilegeWillExpire != nil {
					h.OnTokenPrivilegeWillExpire(channelName)
				}
			}
		},
		OnSubscribeResult: func(requestId uint64, channelName string, errorCode int) {
			for _, h := range handles {
				if h.OnSubscribeResult != nil {
					h.OnSubscribeResult(requestId, channelName, errorCode)
				}
			}
		},
		OnPublishResult: func(requestId uint64, errorCode int) {
			for _, h := range handles {
				if h.OnPublishResult != nil {
					h.OnPublishResult(requestId, errorCode)
				}
			}
		},
		OnLoginResult: func(requestId uint64, errorCode int) {
			for _, h := range handles {
				if h.OnLoginResult != nil {
					h.OnLoginResult(requestId, errorCode)
				}
			}
		},
		OnSetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			for _, h := range handles {
				if h.OnSetChannelMetadataResult != nil {
					h.OnSetChannelMetadataResult(requestId, channelName, channelType, errorCode)
				}
			}
		},
		OnUpdateChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			for _, h := range handles {
				if h.OnUpdateChannelMetadataResult != nil {
					h.OnUpdateChannelMetadataResult(requestId, channelName, channelType, errorCode)
				}
			}
		},
		OnRemoveChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			for _, h := range handles {
				if h.OnRemoveChannelMetadataResult != nil {
					h.OnRemoveChannelMetadataResult(requestId, channelName, channelType, errorCode)
				}
			}
		},
		OnGetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, data *agrtm.IMetadata, errorCode int) {
			for _, h := range handles {
				if h.OnGetChannelMetadataResult != nil {
					h.OnGetChannelMetadataResult(requestId, channelName, channelType, data, errorCode)
				}
			}
		},
		OnSetUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnSetUserMetadataResult != nil {
					h.OnSetUserMetadataResult(requestId, userId, errorCode)
				}
			}
		},
		OnUpdateUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnUpdateUserMetadataResult != nil {
					h.OnUpdateUserMetadataResult(requestId, userId, errorCode)
				}
			}
		},
		OnRemoveUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnRemoveUserMetadataResult != nil {
					h.OnRemoveUserMetadataResult(requestId, userId, errorCode)
				}
			}
		},
		OnGetUserMetadataResult: func(requestId uint64, userId string, data *agrtm.IMetadata, errorCode int) {
			for _, h := range handles {
				if h.OnGetUserMetadataResult != nil {
					h.OnGetUserMetadataResult(requestId, userId, data, errorCode)
				}
			}
		},
		OnSubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnSubscribeUserMetadataResult != nil {
					h.OnSubscribeUserMetadataResult(requestId, userId, errorCode)
				}
			}
		},
		OnSetLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			for _, h := range handles {
				if h.OnSetLockResult != nil {
					h.OnSetLockResult(requestId, channelName, channelType, lockName, errorCode)
				}
			}
		},
		OnRemoveLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			for _, h := range handles {
				if h.OnRemoveLockResult != nil {
					h.OnRemoveLockResult(requestId, channelName, channelType, lockName, errorCode)
				}
			}
		},
		OnReleaseLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			for _, h := range handles {
				if h.OnReleaseLockResult != nil {
					h.OnReleaseLockResult(requestId, channelName, channelType, lockName, errorCode)
				}
			}
		},
		OnAcquireLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int, errorDetails string) {
			for _, h := range handles {
				if h.OnAcquireLockResult != nil {
					h.OnAcquireLockResult(requestId, channelName, channelType, lockName, errorCode, errorDetails)
				}
			}
		},
		OnRevokeLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			for _, h := range handles {
				if h.OnRevokeLockResult != nil {
					h.OnRevokeLockResult(requestId, channelName, channelType, lockName, errorCode)
				}
			}
		},
		OnGetLocksResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockDetailList *agrtm.LockDetail, count uint, errorCode int) {
			for _, h := range handles {
				if h.OnGetLocksResult != nil {
					h.OnGetLocksResult(requestId, channelName, channelType, lockDetailList, count, errorCode)
				}
			}
		},
		OnWhoNowResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			for _, h := range handles {
				if h.OnWhoNowResult != nil {
					h.OnWhoNowResult(requestId, userStateList, count, nextPage, errorCode)
				}
			}
		},
		OnGetOnlineUsersResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			for _, h := range handles {
				if h.OnGetOnlineUsersResult != nil {
					h.OnGetOnlineUsersResult(requestId, userStateList, count, nextPage, errorCode)
				}
			}
		},
		OnWhereNowResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			for _, h := range handles {
				if h.OnWhereNowResult != nil {
					h.OnWhereNowResult(requestId, channels, count, errorCode)
				}
			}
		},
		OnGetUserChannelsResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			for _, h := range handles {
				if h.OnGetUserChannelsResult != nil {
					h.OnGetUserChannelsResult(requestId, channels, count, errorCode)
				}
			}
		},
		OnPresenceSetStateResult: func(requestId uint64, errorCode int) {
			for _, h := range handles {
				if h.OnPresenceSetStateResult != nil {
					h.OnPresenceSetStateResult(requestId, errorCode)
				}
			}
		},
		OnPresenceRemoveStateResult: func(requestId uint64, errorCode int) {
			for _, h := range handles {
				if h.OnPresenceRemoveStateResult != nil {
					h.OnPresenceRemoveStateResult(requestId, errorCode)
				}
			}
		},
		OnPresenceGetStateResult: func(requestId uint64, state *agrtm.UserState, errorCode int) {
			for _, h := range handles {
				if h.OnPresenceGetStateResult != nil {
					h.OnPresenceGetStateResult(requestId, state, errorCode)
				}
			}
		},
		OnLinkStateEvent: func(event *agrtm.LinkStateEvent) {
			for _, h := range handles {
				if h.OnLinkStateEvent != nil {
					h.OnLinkStateEvent(event)
				}
			}
		},
		OnLogoutResult: func(requestId uint64, errorCode int) {
			for _, h := range handles {
				if h.OnLogoutResult != nil {
					h.OnLogoutResult(requestId, errorCode)
				}
			}
		},
		OnRenewTokenResult: func(requestId uint64, serverType agrtm.RtmServiceType, channelName string, errorCode int) {
			for _, h := range handles {
				if h.OnRenewTokenResult != nil {
					h.OnRenewTokenResult(requestId, serverType, channelName, errorCode)
				}
			}
		},
		OnPublishTopicMessageResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			for _, h := range handles {
				if h.OnPublishTopicMessageResult != nil {
					h.OnPublishTopicMessageResult(requestId, channelName, topic, errorCode)
				}
			}
		},
		OnUnsubscribeTopicResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			for _, h := range handles {
				if h.OnUnsubscribeTopicResult != nil {
					h.OnUnsubscribeTopicResult(requestId, channelName, topic, errorCode)
				}
			}
		},
		OnGetSubscribedUserListResult: func(requestId uint64, channelName string, topic string, user *agrtm.UserList, errorCode int) {
			for _, h := range handles {
				if h.OnGetSubscribedUserListResult != nil {
					h.OnGetSubscribedUserListResult(requestId, channelName, topic, user, errorCode)
				}
			}
		},
		OnGetHistoryMessagesResult: func(requestId uint64, messageList []agrtm.HistoryMessage, newStart uint64, errorCode int) {
			for _, h := range handles {
				if h.OnGetHistoryMessagesResult != nil {
					h.OnGetHistoryMessagesResult(requestId, messageList, newStart, errorCode)
				}
			}
		},
		OnUnsubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			for _, h := range handles {
				if h.OnUnsubscribeUserMetadataResult != nil {
					h.OnUnsubscribeUserMetadataResult(requestId, userId, errorCode)
				}
			}
		},
	}
}
