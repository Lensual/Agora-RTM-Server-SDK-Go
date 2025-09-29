package agorartm

import "reflect"

// EventHandlerAdapter adapter, convert the user's event handler to the internal interface
type EventHandlerAdapter struct {
	userHandler RtmEventHandler
	userValue   reflect.Value
	userType    reflect.Type
}

// NewEventHandlerAdapter create event handler adapter
func NewEventHandlerAdapter(userHandler RtmEventHandler) *EventHandlerAdapter {
	// security check: if userHandler is nil, create a safe default adapter
	if userHandler == nil {
		return &EventHandlerAdapter{
			userHandler: nil,
			userValue:   reflect.ValueOf(nil),
			userType:    nil,
		}
	}

	return &EventHandlerAdapter{
		userHandler: userHandler,
		userValue:   reflect.ValueOf(userHandler),
		userType:    reflect.TypeOf(userHandler),
	}
}

// callUserMethod use reflection to call the user's method (if exists)
func (adapter *EventHandlerAdapter) callUserMethod(methodName string, args ...interface{}) {
	// security check: if userHandler is nil, return directly
	if adapter.userHandler == nil {
		return
	}

	// security check: if methodName is empty, return directly
	if methodName == "" {
		return
	}

	// first check if it is a functional configuration
	if config, ok := adapter.userHandler.(*RtmEventHandlerConfig); ok {
		adapter.callFunctionConfig(config, methodName, args...)
		return
	}

	// otherwise use reflection to call the user's method
	method := adapter.userValue.MethodByName(methodName)
	if !method.IsValid() {
		// user did not implement this method, return directly
		return
	}

	// security check: check if the parameter quantity matches
	methodType := method.Type()
	numIn := methodType.NumIn()
	if len(args) != numIn {
		// parameter quantity does not match, return directly
		return
	}

	// prepare parameters, and perform type safety check
	values := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValue := reflect.ValueOf(arg)
		expectedType := methodType.In(i)

		// type safety check
		if arg == nil {
			// if the parameter is nil, check if the method accepts nil
			if expectedType.Kind() == reflect.Ptr || expectedType.Kind() == reflect.Interface {
				values[i] = reflect.Zero(expectedType)
			} else {
				// type does not match, return directly
				return
			}
		} else if !argValue.Type().AssignableTo(expectedType) {
			// type does not match, try type conversion
			if argValue.Type().ConvertibleTo(expectedType) {
				values[i] = argValue.Convert(expectedType)
			} else {
				// cannot convert, return directly
				return
			}
		} else {
			values[i] = argValue
		}
	}

	// security call method
	defer func() {
		if r := recover(); r != nil {
			// if the method call panic, record log but do not affect the program running
			// here can add log record
		}
	}()

	method.Call(values)
}

// callFunctionConfig call the callback function in the functional configuration
func (adapter *EventHandlerAdapter) callFunctionConfig(config *RtmEventHandlerConfig, methodName string, args ...interface{}) {
	// security check: if config is nil, return directly
	if config == nil {
		return
	}

	// security call function, prevent panic
	defer func() {
		if r := recover(); r != nil {
			// if the function call panic, record log but do not affect the program running
			// here can add log record
		}
	}()

	switch methodName {
	case "OnMessageEvent":
		if config.OnMessageEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*MessageEvent); ok {
				config.OnMessageEvent(event)
			}
		}
	case "OnPresenceEvent":
		if config.OnPresenceEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*PresenceEvent); ok {
				config.OnPresenceEvent(event)
			}
		}
	case "OnTopicEvent":
		if config.OnTopicEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*TopicEvent); ok {
				config.OnTopicEvent(event)
			}
		}
	case "OnLockEvent":
		if config.OnLockEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*LockEvent); ok {
				config.OnLockEvent(event)
			}
		}
	case "OnStorageEvent":
		if config.OnStorageEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*StorageEvent); ok {
				config.OnStorageEvent(event)
			}
		}
	case "OnJoinResult":
		if config.OnJoinResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if userId, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnJoinResult(requestId, channelName, userId, errorCode)
						}
					}
				}
			}
		}
	case "OnLeaveResult":
		if config.OnLeaveResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if userId, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnLeaveResult(requestId, channelName, userId, errorCode)
						}
					}
				}
			}
		}
	case "OnJoinTopicResult":
		if config.OnJoinTopicResult != nil && len(args) >= 6 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if userId, ok := args[2].(string); ok {
						if topic, ok := args[3].(string); ok {
							if meta, ok := args[4].(string); ok {
								if errorCode, ok := args[5].(int); ok {
									config.OnJoinTopicResult(requestId, channelName, userId, topic, meta, errorCode)
								}
							}
						}
					}
				}
			}
		}
	case "OnLeaveTopicResult":
		if config.OnLeaveTopicResult != nil && len(args) >= 6 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if userId, ok := args[2].(string); ok {
						if topic, ok := args[3].(string); ok {
							if meta, ok := args[4].(string); ok {
								if errorCode, ok := args[5].(int); ok {
									config.OnLeaveTopicResult(requestId, channelName, userId, topic, meta, errorCode)
								}
							}
						}
					}
				}
			}
		}
	case "OnSubscribeTopicResult":
		if config.OnSubscribeTopicResult != nil && len(args) >= 7 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if userId, ok := args[2].(string); ok {
						if topic, ok := args[3].(string); ok {
							if succeedUsers, ok := args[4].(UserList); ok {
								if failedUsers, ok := args[5].(UserList); ok {
									if errorCode, ok := args[6].(int); ok {
										config.OnSubscribeTopicResult(requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
									}
								}
							}
						}
					}
				}
			}
		}
	case "OnConnectionStateChanged":
		if config.OnConnectionStateChanged != nil && len(args) >= 3 {
			if channelName, ok := args[0].(string); ok {
				if state, ok := args[1].(int); ok {
					if reason, ok := args[2].(int); ok {
						config.OnConnectionStateChanged(channelName, state, reason)
					}
				}
			}
		}
	case "OnTokenPrivilegeWillExpire":
		if config.OnTokenPrivilegeWillExpire != nil && len(args) >= 1 {
			if channelName, ok := args[0].(string); ok {
				config.OnTokenPrivilegeWillExpire(channelName)
			}
		}
	case "OnSubscribeResult":
		if config.OnSubscribeResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnSubscribeResult(requestId, channelName, errorCode)
					}
				}
			}
		}
	case "OnPublishResult":
		if config.OnPublishResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(int); ok {
					config.OnPublishResult(requestId, errorCode)
				}
			}
		}
	case "OnLoginResult":
		if config.OnLoginResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(int); ok {
					config.OnLoginResult(requestId, errorCode)
				}
			}
		}
	case "OnSetChannelMetadataResult":
		if config.OnSetChannelMetadataResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnSetChannelMetadataResult(requestId, channelName, channelType, errorCode)
						}
					}
				}
			}
		}
	case "OnUpdateChannelMetadataResult":
		if config.OnUpdateChannelMetadataResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnUpdateChannelMetadataResult(requestId, channelName, channelType, errorCode)
						}
					}
				}
			}
		}
	case "OnRemoveChannelMetadataResult":
		if config.OnRemoveChannelMetadataResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnRemoveChannelMetadataResult(requestId, channelName, channelType, errorCode)
						}
					}
				}
			}
		}
	case "OnGetChannelMetadataResult":
		if config.OnGetChannelMetadataResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if data, ok := args[3].(*IMetadata); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnGetChannelMetadataResult(requestId, channelName, channelType, data, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnSetUserMetadataResult":
		if config.OnSetUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnSetUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnUpdateUserMetadataResult":
		if config.OnUpdateUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnUpdateUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnRemoveUserMetadataResult":
		if config.OnRemoveUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnRemoveUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnGetUserMetadataResult":
		if config.OnGetUserMetadataResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if data, ok := args[2].(*IMetadata); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnGetUserMetadataResult(requestId, userId, data, errorCode)
						}
					}
				}
			}
		}
	case "OnSubscribeUserMetadataResult":
		if config.OnSubscribeUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnSubscribeUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnSetLockResult":
		if config.OnSetLockResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnSetLockResult(requestId, channelName, channelType, lockName, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnRemoveLockResult":
		if config.OnRemoveLockResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnRemoveLockResult(requestId, channelName, channelType, lockName, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnReleaseLockResult":
		if config.OnReleaseLockResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnReleaseLockResult(requestId, channelName, channelType, lockName, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnAcquireLockResult":
		if config.OnAcquireLockResult != nil && len(args) >= 6 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								if errorDetails, ok := args[5].(string); ok {
									config.OnAcquireLockResult(requestId, channelName, channelType, lockName, errorCode, errorDetails)
								}
							}
						}
					}
				}
			}
		}
	case "OnRevokeLockResult":
		if config.OnRevokeLockResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnRevokeLockResult(requestId, channelName, channelType, lockName, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnGetLocksResult":
		if config.OnGetLocksResult != nil && len(args) >= 6 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RtmChannelType); ok {
						if lockDetailList, ok := args[3].(*LockDetail); ok {
							if count, ok := args[4].(uint); ok {
								if errorCode, ok := args[5].(int); ok {
									config.OnGetLocksResult(requestId, channelName, channelType, lockDetailList, count, errorCode)
								}
							}
						}
					}
				}
			}
		}
	case "OnWhoNowResult":
		if config.OnWhoNowResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if userStateList, ok := args[1].(*UserState); ok {
					if count, ok := args[2].(uint); ok {
						if nextPage, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnWhoNowResult(requestId, userStateList, count, nextPage, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnGetOnlineUsersResult":
		if config.OnGetOnlineUsersResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if userStateList, ok := args[1].(*UserState); ok {
					if count, ok := args[2].(uint); ok {
						if nextPage, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnGetOnlineUsersResult(requestId, userStateList, count, nextPage, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnWhereNowResult":
		if config.OnWhereNowResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channels, ok := args[1].(*ChannelInfo); ok {
					if count, ok := args[2].(uint); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnWhereNowResult(requestId, channels, count, errorCode)
						}
					}
				}
			}
		}
	case "OnGetUserChannelsResult":
		if config.OnGetUserChannelsResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channels, ok := args[1].(*ChannelInfo); ok {
					if count, ok := args[2].(uint); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnGetUserChannelsResult(requestId, channels, count, errorCode)
						}
					}
				}
			}
		}
	case "OnPresenceSetStateResult":
		if config.OnPresenceSetStateResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(int); ok {
					config.OnPresenceSetStateResult(requestId, errorCode)
				}
			}
		}
	case "OnPresenceRemoveStateResult":
		if config.OnPresenceRemoveStateResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(int); ok {
					config.OnPresenceRemoveStateResult(requestId, errorCode)
				}
			}
		}
	case "OnPresenceGetStateResult":
		if config.OnPresenceGetStateResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if state, ok := args[1].(*UserState); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnPresenceGetStateResult(requestId, state, errorCode)
					}
				}
			}
		}
	case "OnLinkStateEvent":
		if config.OnLinkStateEvent != nil && len(args) >= 1 {
			if event, ok := args[0].(*LinkStateEvent); ok {
				config.OnLinkStateEvent(event)
			}
		}
	case "OnLogoutResult":
		if config.OnLogoutResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(int); ok {
					config.OnLogoutResult(requestId, errorCode)
				}
			}
		}
	case "OnRenewTokenResult":
		if config.OnRenewTokenResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if serverType, ok := args[1].(RtmServiceType); ok {
					if channelName, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnRenewTokenResult(requestId, serverType, channelName, errorCode)
						}
					}
				}
			}
		}
	case "OnPublishTopicMessageResult":
		if config.OnPublishTopicMessageResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if topic, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnPublishTopicMessageResult(requestId, channelName, topic, errorCode)
						}
					}
				}
			}
		}
	case "OnUnsubscribeTopicResult":
		if config.OnUnsubscribeTopicResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if topic, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnUnsubscribeTopicResult(requestId, channelName, topic, errorCode)
						}
					}
				}
			}
		}
	case "OnGetSubscribedUserListResult":
		if config.OnGetSubscribedUserListResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if topic, ok := args[2].(string); ok {
						if user, ok := args[3].(*UserList); ok {
							if errorCode, ok := args[4].(int); ok {
								config.OnGetSubscribedUserListResult(requestId, channelName, topic, user, errorCode)
							}
						}
					}
				}
			}
		}
	case "OnGetHistoryMessagesResult":
		if config.OnGetHistoryMessagesResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if messageList, ok := args[1].([]HistoryMessage); ok {
					if newStart, ok := args[2].(uint64); ok {
						if errorCode, ok := args[3].(int); ok {
							config.OnGetHistoryMessagesResult(requestId, messageList, newStart, errorCode)
						}
					}
				}
			}
		}
	case "OnUnsubscribeUserMetadataResult":
		if config.OnUnsubscribeUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(int); ok {
						config.OnUnsubscribeUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	}
}

// 实现 IRtmEventHandlerBridgeHandler 接口，使用反射调用用户实现的方法
func (adapter *EventHandlerAdapter) OnMessageEvent(event *MessageEvent) {
	adapter.callUserMethod("OnMessageEvent", event)
}

func (adapter *EventHandlerAdapter) OnPresenceEvent(event *PresenceEvent) {
	adapter.callUserMethod("OnPresenceEvent", event)
}

func (adapter *EventHandlerAdapter) OnTopicEvent(event *TopicEvent) {
	adapter.callUserMethod("OnTopicEvent", event)
}

func (adapter *EventHandlerAdapter) OnLockEvent(event *LockEvent) {
	adapter.callUserMethod("OnLockEvent", event)
}

func (adapter *EventHandlerAdapter) OnStorageEvent(event *StorageEvent) {
	adapter.callUserMethod("OnStorageEvent", event)
}

func (adapter *EventHandlerAdapter) OnJoinResult(requestId uint64, channelName string, userId string, errorCode int) {
	adapter.callUserMethod("OnJoinResult", requestId, channelName, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnLeaveResult(requestId uint64, channelName string, userId string, errorCode int) {
	adapter.callUserMethod("OnLeaveResult", requestId, channelName, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
	adapter.callUserMethod("OnJoinTopicResult", requestId, channelName, userId, topic, meta, errorCode)
}

func (adapter *EventHandlerAdapter) OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode int) {
	adapter.callUserMethod("OnLeaveTopicResult", requestId, channelName, userId, topic, meta, errorCode)
}

func (adapter *EventHandlerAdapter) OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers UserList, failedUsers UserList, errorCode int) {
	adapter.callUserMethod("OnSubscribeTopicResult", requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
}

func (adapter *EventHandlerAdapter) OnConnectionStateChanged(channelName string, state int, reason int) {
	adapter.callUserMethod("OnConnectionStateChanged", channelName, state, reason)
}

func (adapter *EventHandlerAdapter) OnTokenPrivilegeWillExpire(channelName string) {
	adapter.callUserMethod("OnTokenPrivilegeWillExpire", channelName)
}

func (adapter *EventHandlerAdapter) OnSubscribeResult(requestId uint64, channelName string, errorCode int) {
	adapter.callUserMethod("OnSubscribeResult", requestId, channelName, errorCode)
}

func (adapter *EventHandlerAdapter) OnPublishResult(requestId uint64, errorCode int) {
	adapter.callUserMethod("OnPublishResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnLoginResult(requestId uint64, errorCode int) {
	adapter.callUserMethod("OnLoginResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetChannelMetadataResult(requestId uint64, channelName string, channelType RtmChannelType, errorCode int) {
	adapter.callUserMethod("OnSetChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType RtmChannelType, errorCode int) {
	adapter.callUserMethod("OnUpdateChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType RtmChannelType, errorCode int) {
	adapter.callUserMethod("OnRemoveChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetChannelMetadataResult(requestId uint64, channelName string, channelType RtmChannelType, data *IMetadata, errorCode int) {
	adapter.callUserMethod("OnGetChannelMetadataResult", requestId, channelName, channelType, data, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetUserMetadataResult(requestId uint64, userId string, errorCode int) {
	adapter.callUserMethod("OnSetUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode int) {
	adapter.callUserMethod("OnUpdateUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode int) {
	adapter.callUserMethod("OnRemoveUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetUserMetadataResult(requestId uint64, userId string, data *IMetadata, errorCode int) {
	adapter.callUserMethod("OnGetUserMetadataResult", requestId, userId, data, errorCode)
}

func (adapter *EventHandlerAdapter) OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode int) {
	adapter.callUserMethod("OnSubscribeUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetLockResult(requestId uint64, channelName string, channelType RtmChannelType, lockName string, errorCode int) {
	adapter.callUserMethod("OnSetLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveLockResult(requestId uint64, channelName string, channelType RtmChannelType, lockName string, errorCode int) {
	adapter.callUserMethod("OnRemoveLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnReleaseLockResult(requestId uint64, channelName string, channelType RtmChannelType, lockName string, errorCode int) {
	adapter.callUserMethod("OnReleaseLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnAcquireLockResult(requestId uint64, channelName string, channelType RtmChannelType, lockName string, errorCode int, errorDetails string) {
	adapter.callUserMethod("OnAcquireLockResult", requestId, channelName, channelType, lockName, errorCode, errorDetails)
}

func (adapter *EventHandlerAdapter) OnRevokeLockResult(requestId uint64, channelName string, channelType RtmChannelType, lockName string, errorCode int) {
	adapter.callUserMethod("OnRevokeLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetLocksResult(requestId uint64, channelName string, channelType RtmChannelType, lockDetailList *LockDetail, count uint, errorCode int) {
	adapter.callUserMethod("OnGetLocksResult", requestId, channelName, channelType, lockDetailList, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnWhoNowResult(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode int) {
	adapter.callUserMethod("OnWhoNowResult", requestId, userStateList, count, nextPage, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetOnlineUsersResult(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode int) {
	adapter.callUserMethod("OnGetOnlineUsersResult", requestId, userStateList, count, nextPage, errorCode)
}

func (adapter *EventHandlerAdapter) OnWhereNowResult(requestId uint64, channels *ChannelInfo, count uint, errorCode int) {
	adapter.callUserMethod("OnWhereNowResult", requestId, channels, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetUserChannelsResult(requestId uint64, channels *ChannelInfo, count uint, errorCode int) {
	adapter.callUserMethod("OnGetUserChannelsResult", requestId, channels, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceSetStateResult(requestId uint64, errorCode int) {
	adapter.callUserMethod("OnPresenceSetStateResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceRemoveStateResult(requestId uint64, errorCode int) {
	adapter.callUserMethod("OnPresenceRemoveStateResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceGetStateResult(requestId uint64, state *UserState, errorCode int) {
	adapter.callUserMethod("OnPresenceGetStateResult", requestId, state, errorCode)
}

func (adapter *EventHandlerAdapter) OnLinkStateEvent(event *LinkStateEvent) {
	adapter.callUserMethod("OnLinkStateEvent", event)
}

func (adapter *EventHandlerAdapter) OnLogoutResult(requestId uint64, errorCode int) {
	adapter.callUserMethod("OnLogoutResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnRenewTokenResult(requestId uint64, serverType RtmServiceType, channelName string, errorCode int) {
	adapter.callUserMethod("OnRenewTokenResult", requestId, serverType, channelName, errorCode)
}

func (adapter *EventHandlerAdapter) OnPublishTopicMessageResult(requestId uint64, channelName string, topic string, errorCode int) {
	adapter.callUserMethod("OnPublishTopicMessageResult", requestId, channelName, topic, errorCode)
}

func (adapter *EventHandlerAdapter) OnUnsubscribeTopicResult(requestId uint64, channelName string, topic string, errorCode int) {
	adapter.callUserMethod("OnUnsubscribeTopicResult", requestId, channelName, topic, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetSubscribedUserListResult(requestId uint64, channelName string, topic string, user *UserList, errorCode int) {
	adapter.callUserMethod("OnGetSubscribedUserListResult", requestId, channelName, topic, user, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetHistoryMessagesResult(requestId uint64, messageList []HistoryMessage, newStart uint64, errorCode int) {
	adapter.callUserMethod("OnGetHistoryMessagesResult", requestId, messageList, newStart, errorCode)
}

func (adapter *EventHandlerAdapter) OnUnsubscribeUserMetadataResult(requestId uint64, userId string, errorCode int) {
	adapter.callUserMethod("OnUnsubscribeUserMetadataResult", requestId, userId, errorCode)
}
