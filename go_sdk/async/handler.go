package async

import agrtm "github.com/AgoraIO-Extensions/Agora-RTM-Server-SDK-Go/go_sdk/rtm"

func NewAsyncCallRtmEventHandler() *agrtm.RtmEventHandler {
	return &agrtm.RtmEventHandler{
		//#region IRtmClient callback

		OnLoginResult: func(requestId uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnLogoutResult: func(requestId uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnRenewTokenResult: func(requestId uint64, serverType agrtm.RtmServiceType, channelName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetOnlineUsersResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		OnGetUserChannelsResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//#endregion IRtmClient callback

		//#region MessageChannel callback

		OnJoinResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnLeaveResult: func(requestId uint64, channelName string, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnSubscribeResult: func(requestId uint64, channelName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//TODO SDK缺少 OnUnsubscribeResult

		OnPublishResult: func(requestId uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetHistoryMessagesResult: func(requestId uint64, messageList []agrtm.HistoryMessage, newStart uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//#endregion MessageChannel callback

		//#region StreamChannel callback

		OnJoinTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnLeaveTopicResult: func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnSubscribeTopicResult: func(requestId uint64, channelName string, userId string, topic string, succeedUsers agrtm.UserList, failedUsers agrtm.UserList, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnPublishTopicMessageResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnUnsubscribeTopicResult: func(requestId uint64, channelName string, topic string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetSubscribedUserListResult: func(requestId uint64, channelName string, topic string, user *agrtm.UserList, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		// NOTE: OnRenewTokenResult 已归类至 `IRtmClient callback` 中

		//#endregion StreamChannel callback

		//#region Presence callback

		OnWhoNowResult: func(requestId uint64, userStateList *agrtm.UserState, count uint, nextPage string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnWhereNowResult: func(requestId uint64, channels *agrtm.ChannelInfo, count uint, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnPresenceSetStateResult: func(requestId uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnPresenceGetStateResult: func(requestId uint64, state *agrtm.UserState, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnPresenceRemoveStateResult: func(requestId uint64, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//#endregion Presence callback

		//#region Storage callback

		OnSetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, data *agrtm.IMetadata, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnRemoveChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnUpdateChannelMetadataResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		OnSetUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetUserMetadataResult: func(requestId uint64, userId string, data *agrtm.IMetadata, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnRemoveUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnUpdateUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		OnSubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnUnsubscribeUserMetadataResult: func(requestId uint64, userId string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//#endregion Storage callback

		//#region Lock callback

		OnSetLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnAcquireLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int, errorDetails string) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnReleaseLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnRevokeLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnGetLocksResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockDetailList *agrtm.LockDetail, count uint, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},
		OnRemoveLockResult: func(requestId uint64, channelName string, channelType agrtm.RtmChannelType, lockName string, errorCode int) {
			//#region reqIdMapLocker.Lock()
			reqIdMapLocker.Lock()

			iAsyncCall, ok := reqIdMap[requestId]
			delete(reqIdMap, requestId)

			reqIdMapLocker.Unlock()
			//#endregion reqIdMapLocker.Lock()

			if ok {
				asyncCall := (iAsyncCall).(*asyncCallImpl[int])
				if errorCode != 0 {
					asyncCall.error(errorCode)
					return
				}
				asyncCall.complete(errorCode)
			}
		},

		//#endregion Lock callback
	}

}
