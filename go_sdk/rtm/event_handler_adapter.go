package agorartm

import "reflect"

// EventHandlerAdapter 适配器，将用户的事件处理器转换为内部接口
type EventHandlerAdapter struct {
	userHandler RtmEventHandler
	userValue   reflect.Value
	userType    reflect.Type
}

// NewEventHandlerAdapter 创建事件处理器适配器
func NewEventHandlerAdapter(userHandler RtmEventHandler) *EventHandlerAdapter {
	// 安全检查：如果 userHandler 为 nil，创建一个安全的默认适配器
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

// callUserMethod 使用反射调用用户的方法（如果存在）
func (adapter *EventHandlerAdapter) callUserMethod(methodName string, args ...interface{}) {
	// 安全检查：如果 userHandler 为 nil，直接返回
	if adapter.userHandler == nil {
		return
	}

	// 安全检查：如果方法名为空，直接返回
	if methodName == "" {
		return
	}

	// 首先检查是否是函数式配置
	if config, ok := adapter.userHandler.(*RtmEventHandlerConfig); ok {
		adapter.callFunctionConfig(config, methodName, args...)
		return
	}

	// 否则使用反射调用用户方法
	method := adapter.userValue.MethodByName(methodName)
	if !method.IsValid() {
		// 用户没有实现这个方法，直接返回
		return
	}

	// 安全检查：检查参数数量是否匹配
	methodType := method.Type()
	numIn := methodType.NumIn()
	if len(args) != numIn {
		// 参数数量不匹配，直接返回
		return
	}

	// 准备参数，并进行类型安全检查
	values := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValue := reflect.ValueOf(arg)
		expectedType := methodType.In(i)

		// 类型安全检查
		if arg == nil {
			// 如果参数为 nil，检查方法是否接受 nil
			if expectedType.Kind() == reflect.Ptr || expectedType.Kind() == reflect.Interface {
				values[i] = reflect.Zero(expectedType)
			} else {
				// 类型不匹配，直接返回
				return
			}
		} else if !argValue.Type().AssignableTo(expectedType) {
			// 类型不匹配，尝试类型转换
			if argValue.Type().ConvertibleTo(expectedType) {
				values[i] = argValue.Convert(expectedType)
			} else {
				// 无法转换，直接返回
				return
			}
		} else {
			values[i] = argValue
		}
	}

	// 安全调用方法
	defer func() {
		if r := recover(); r != nil {
			// 如果方法调用 panic，记录日志但不影响程序运行
			// 这里可以添加日志记录
		}
	}()

	method.Call(values)
}

// callFunctionConfig 调用函数式配置中的回调函数
func (adapter *EventHandlerAdapter) callFunctionConfig(config *RtmEventHandlerConfig, methodName string, args ...interface{}) {
	// 安全检查：如果 config 为 nil，直接返回
	if config == nil {
		return
	}

	// 安全调用函数，防止 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果函数调用 panic，记录日志但不影响程序运行
			// 这里可以添加日志记录
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
								if errorCode, ok := args[5].(RTM_ERROR_CODE); ok {
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
								if errorCode, ok := args[5].(RTM_ERROR_CODE); ok {
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
									if errorCode, ok := args[6].(RTM_ERROR_CODE); ok {
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
				if state, ok := args[1].(RTM_CONNECTION_STATE); ok {
					if reason, ok := args[2].(RTM_CONNECTION_CHANGE_REASON); ok {
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
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
						config.OnSubscribeResult(requestId, channelName, errorCode)
					}
				}
			}
		}
	case "OnPublishResult":
		if config.OnPublishResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(RTM_ERROR_CODE); ok {
					config.OnPublishResult(requestId, errorCode)
				}
			}
		}
	case "OnLoginResult":
		if config.OnLoginResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(RTM_ERROR_CODE); ok {
					config.OnLoginResult(requestId, errorCode)
				}
			}
		}
	case "OnSetChannelMetadataResult":
		if config.OnSetChannelMetadataResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if data, ok := args[3].(*IMetadata); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
						config.OnSetUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnUpdateUserMetadataResult":
		if config.OnUpdateUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
						config.OnUpdateUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnRemoveUserMetadataResult":
		if config.OnRemoveUserMetadataResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if userId, ok := args[1].(string); ok {
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
						config.OnSubscribeUserMetadataResult(requestId, userId, errorCode)
					}
				}
			}
		}
	case "OnSetLockResult":
		if config.OnSetLockResult != nil && len(args) >= 5 {
			if requestId, ok := args[0].(uint64); ok {
				if channelName, ok := args[1].(string); ok {
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockName, ok := args[3].(string); ok {
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
					if channelType, ok := args[2].(RTM_CHANNEL_TYPE); ok {
						if lockDetailList, ok := args[3].(*LockDetail); ok {
							if count, ok := args[4].(uint); ok {
								if errorCode, ok := args[5].(RTM_ERROR_CODE); ok {
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
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
							config.OnGetUserChannelsResult(requestId, channels, count, errorCode)
						}
					}
				}
			}
		}
	case "OnPresenceSetStateResult":
		if config.OnPresenceSetStateResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(RTM_ERROR_CODE); ok {
					config.OnPresenceSetStateResult(requestId, errorCode)
				}
			}
		}
	case "OnPresenceRemoveStateResult":
		if config.OnPresenceRemoveStateResult != nil && len(args) >= 2 {
			if requestId, ok := args[0].(uint64); ok {
				if errorCode, ok := args[1].(RTM_ERROR_CODE); ok {
					config.OnPresenceRemoveStateResult(requestId, errorCode)
				}
			}
		}
	case "OnPresenceGetStateResult":
		if config.OnPresenceGetStateResult != nil && len(args) >= 3 {
			if requestId, ok := args[0].(uint64); ok {
				if state, ok := args[1].(*UserState); ok {
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
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
				if errorCode, ok := args[1].(RTM_ERROR_CODE); ok {
					config.OnLogoutResult(requestId, errorCode)
				}
			}
		}
	case "OnRenewTokenResult":
		if config.OnRenewTokenResult != nil && len(args) >= 4 {
			if requestId, ok := args[0].(uint64); ok {
				if serverType, ok := args[1].(RTM_SERVICE_TYPE); ok {
					if channelName, ok := args[2].(string); ok {
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
							if errorCode, ok := args[4].(RTM_ERROR_CODE); ok {
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
						if errorCode, ok := args[3].(RTM_ERROR_CODE); ok {
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
					if errorCode, ok := args[2].(RTM_ERROR_CODE); ok {
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

func (adapter *EventHandlerAdapter) OnJoinResult(requestId uint64, channelName string, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnJoinResult", requestId, channelName, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnLeaveResult(requestId uint64, channelName string, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnLeaveResult", requestId, channelName, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnJoinTopicResult", requestId, channelName, userId, topic, meta, errorCode)
}

func (adapter *EventHandlerAdapter) OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnLeaveTopicResult", requestId, channelName, userId, topic, meta, errorCode)
}

func (adapter *EventHandlerAdapter) OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers UserList, failedUsers UserList, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSubscribeTopicResult", requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode)
}

func (adapter *EventHandlerAdapter) OnConnectionStateChanged(channelName string, state RTM_CONNECTION_STATE, reason RTM_CONNECTION_CHANGE_REASON) {
	adapter.callUserMethod("OnConnectionStateChanged", channelName, state, reason)
}

func (adapter *EventHandlerAdapter) OnTokenPrivilegeWillExpire(channelName string) {
	adapter.callUserMethod("OnTokenPrivilegeWillExpire", channelName)
}

func (adapter *EventHandlerAdapter) OnSubscribeResult(requestId uint64, channelName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSubscribeResult", requestId, channelName, errorCode)
}

func (adapter *EventHandlerAdapter) OnPublishResult(requestId uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnPublishResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnLoginResult(requestId uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnLoginResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetChannelMetadataResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSetChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnUpdateChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnRemoveChannelMetadataResult", requestId, channelName, channelType, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetChannelMetadataResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, data *IMetadata, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetChannelMetadataResult", requestId, channelName, channelType, data, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetUserMetadataResult(requestId uint64, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSetUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnUpdateUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnRemoveUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetUserMetadataResult(requestId uint64, userId string, data *IMetadata, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetUserMetadataResult", requestId, userId, data, errorCode)
}

func (adapter *EventHandlerAdapter) OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSubscribeUserMetadataResult", requestId, userId, errorCode)
}

func (adapter *EventHandlerAdapter) OnSetLockResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnSetLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnRemoveLockResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnRemoveLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnReleaseLockResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnReleaseLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnAcquireLockResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE, errorDetails string) {
	adapter.callUserMethod("OnAcquireLockResult", requestId, channelName, channelType, lockName, errorCode, errorDetails)
}

func (adapter *EventHandlerAdapter) OnRevokeLockResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnRevokeLockResult", requestId, channelName, channelType, lockName, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetLocksResult(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockDetailList *LockDetail, count uint, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetLocksResult", requestId, channelName, channelType, lockDetailList, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnWhoNowResult(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnWhoNowResult", requestId, userStateList, count, nextPage, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetOnlineUsersResult(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetOnlineUsersResult", requestId, userStateList, count, nextPage, errorCode)
}

func (adapter *EventHandlerAdapter) OnWhereNowResult(requestId uint64, channels *ChannelInfo, count uint, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnWhereNowResult", requestId, channels, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetUserChannelsResult(requestId uint64, channels *ChannelInfo, count uint, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetUserChannelsResult", requestId, channels, count, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceSetStateResult(requestId uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnPresenceSetStateResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceRemoveStateResult(requestId uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnPresenceRemoveStateResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnPresenceGetStateResult(requestId uint64, state *UserState, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnPresenceGetStateResult", requestId, state, errorCode)
}

func (adapter *EventHandlerAdapter) OnLinkStateEvent(event *LinkStateEvent) {
	adapter.callUserMethod("OnLinkStateEvent", event)
}

func (adapter *EventHandlerAdapter) OnLogoutResult(requestId uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnLogoutResult", requestId, errorCode)
}

func (adapter *EventHandlerAdapter) OnRenewTokenResult(requestId uint64, serverType RTM_SERVICE_TYPE, channelName string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnRenewTokenResult", requestId, serverType, channelName, errorCode)
}

func (adapter *EventHandlerAdapter) OnPublishTopicMessageResult(requestId uint64, channelName string, topic string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnPublishTopicMessageResult", requestId, channelName, topic, errorCode)
}

func (adapter *EventHandlerAdapter) OnUnsubscribeTopicResult(requestId uint64, channelName string, topic string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnUnsubscribeTopicResult", requestId, channelName, topic, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetSubscribedUserListResult(requestId uint64, channelName string, topic string, user *UserList, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetSubscribedUserListResult", requestId, channelName, topic, user, errorCode)
}

func (adapter *EventHandlerAdapter) OnGetHistoryMessagesResult(requestId uint64, messageList []HistoryMessage, newStart uint64, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnGetHistoryMessagesResult", requestId, messageList, newStart, errorCode)
}

func (adapter *EventHandlerAdapter) OnUnsubscribeUserMetadataResult(requestId uint64, userId string, errorCode RTM_ERROR_CODE) {
	adapter.callUserMethod("OnUnsubscribeUserMetadataResult", requestId, userId, errorCode)
}
