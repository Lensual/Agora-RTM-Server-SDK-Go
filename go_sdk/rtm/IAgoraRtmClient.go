package agorartm

/*
#include "C_IAgoraRtmClient.h"
#include "C_AgoraRtmBase.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// #region agora

// #region agora::rtm

/**
 *  Configurations for RTM Client.
 */
type RtmConfig struct {
	AppId             string
	UserId            string
	AreaCode          RTM_AREA_CODE
	ProtocolType      uint32
	PresenceTimeout   uint32
	HeartbeatInterval uint32
	Context           unsafe.Pointer
	UseStringUserId   bool
	Multipath         bool
	EventHandler      RtmEventHandler
	LogConfig         *RtmLogConfig
	ProxyConfig       *RtmProxyConfig
	EncryptionConfig  *RtmEncryptionConfig
	PrivateConfig     *RtmPrivateConfig
}

// #region RtmConfig

/**
 * The App ID of your project.
 */
func (this_ *RtmConfig) GetAppId() string {
	return this_.AppId
}

/**
 * The App ID of your project.
 */
func (this_ *RtmConfig) SetAppId(appId string) {
	this_.AppId = appId
}

/**
 * The ID of the user.
 */
func (this_ *RtmConfig) GetUserId() string {
	return this_.UserId
}

/**
 * The ID of the user.
 */
func (this_ *RtmConfig) SetUserId(userId string) {
	this_.UserId = userId
}

/**
 * The region for connection. This advanced feature applies to scenarios that
 * have regional restrictions.
 *
 * For the regions that Agora supports, see #AREA_CODE.
 *
 * After specifying the region, the SDK connects to the Agora servers within
 * that region.
 */
func (this_ *RtmConfig) GetAreaCode() RTM_AREA_CODE {
	return this_.AreaCode
}

/**
 * The region for connection. This advanced feature applies to scenarios that
 * have regional restrictions.
 *
 * For the regions that Agora supports, see #AREA_CODE.
 *
 * After specifying the region, the SDK connects to the Agora servers within
 * that region.
 */
func (this_ *RtmConfig) SetAreaCode(areaCode RTM_AREA_CODE) {
	this_.AreaCode = areaCode
}

/**
 * Presence timeout in seconds, specify the timeout value when you lost connection between sdk
 * and rtm service.
 */
func (this_ *RtmConfig) GetPresenceTimeout() uint32 {
	return this_.PresenceTimeout
}

/**
 * Presence timeout in seconds, specify the timeout value when you lost connection between sdk
 * and rtm service.
 */
func (this_ *RtmConfig) SetPresenceTimeout(presenceTimeout uint32) {
	this_.PresenceTimeout = presenceTimeout
}

/**
 * - For Android, it is the context of Activity or Application.
 * - For Windows, it is the window handle of app. Once set, this parameter enables you to plug
 * or unplug the video devices while they are powered.
 */
func (this_ *RtmConfig) GetContext() unsafe.Pointer {
	return this_.Context
}

/**
 * - For Android, it is the context of Activity or Application.
 * - For Windows, it is the window handle of app. Once set, this parameter enables you to plug
 * or unplug the video devices while they are powered.
 */
func (this_ *RtmConfig) SetContext(context unsafe.Pointer) {
	this_.Context = context
}

/**
 * Whether to use String user IDs, if you are using RTC products with Int user IDs,
 * set this value as 'false'. Otherwise errors might occur.
 */
func (this_ *RtmConfig) GetUseStringUserId() bool {
	return this_.UseStringUserId
}

/**
 * Whether to use String user IDs, if you are using RTC products with Int user IDs,
 * set this value as 'false'. Otherwise errors might occur.
 */
func (this_ *RtmConfig) SetUseStringUserId(useStringUserId bool) {
	this_.UseStringUserId = useStringUserId
}

/**
 * The callbacks handler
 */
func (this_ *RtmConfig) GetEventHandler() RtmEventHandler {
	return this_.EventHandler
}

/**
 * The callbacks handler
 */
func (this_ *RtmConfig) SetEventHandler(eventHandler RtmEventHandler) {
	this_.EventHandler = eventHandler
}

/**
 * The config for customer set log path, log size and log level.
 */
func (this_ *RtmConfig) GetLogConfig() *RtmLogConfig {
	return this_.LogConfig
}

/**
 * The config for customer set log path, log size and log level.
 */
func (this_ *RtmConfig) SetLogConfig(logConfig *RtmLogConfig) {
	this_.LogConfig = logConfig
}

/**
 * The config for proxy setting
 */
func (this_ *RtmConfig) GetProxyConfig() *RtmProxyConfig {
	return this_.ProxyConfig
}

/**
 * The config for proxy setting
 */
func (this_ *RtmConfig) SetProxyConfig(proxyConfig *RtmProxyConfig) {
	this_.ProxyConfig = proxyConfig
}

/**
 * The config for encryption setting
 */
func (this_ *RtmConfig) GetEncryptionConfig() *RtmEncryptionConfig {
	return this_.EncryptionConfig
}

/**
 * The config for encryption setting
 */
func (this_ *RtmConfig) SetEncryptionConfig(encryptionConfig *RtmEncryptionConfig) {
	this_.EncryptionConfig = encryptionConfig
}

func NewRtmConfig() *RtmConfig {
	config := &RtmConfig{
		AppId:             "",
		UserId:            "",
		AreaCode:          RTM_AREA_CODE_GLOB,
		ProtocolType:      0,
		HeartbeatInterval: 0,
		Context:           nil,
		UseStringUserId:   true,
		Multipath:         false,
		EventHandler:      nil,
		LogConfig:         nil,
		ProxyConfig:       nil,
		EncryptionConfig:  nil,
		PrivateConfig:     nil,
		PresenceTimeout:   30,
	}

	return config
}

// #endregion RtmConfig

/**
 * The IRtmEventHandler class.
 *
 * The SDK uses this class to send callback event notifications to the app, and the app inherits
 * the methods in this class to retrieve these event notifications.
 *
 * All methods in this class have their default (empty)  implementations, and the app can inherit
 * only some of the required events instead of all. In the callback methods, the app should avoid
 * time-consuming tasks or calling blocking APIs, otherwise the SDK may not work properly.
 */
// old IRtmEventHandler is deleted, please use new RtmEventHandler interface

// new user friendly event handler interface design

// RtmEventHandler define the event handler interface that user can implement
// user only need to implement the needed methods, the methods that are not implemented will be ignored by SDK
//
// commonly used callback methods:
//   - OnLoginResult: login result callback
//   - OnLogoutResult: logout result callback
//   - OnMessageEvent: message event callback
//   - OnPresenceEvent: online status event callback
//   - OnSubscribeResult: subscribe result callback
//   - OnPublishResult: publish result callback
//
// usage one (object oriented):
//
//	type MyEventHandler struct{}
//	func (h *MyEventHandler) OnLoginResult(requestId uint64, errorCode RTM_ERROR_CODE) {
//	    // handle login result
//	}
//	rtmConfig.SetEventHandler(&MyEventHandler{})
//
// usage two (functional):
//
//	handler := &RtmEventHandlerConfig{
//	    OnLoginResult: func(requestId uint64, errorCode RTM_ERROR_CODE) {
//	        // handle login result
//	    },
//	}
//	rtmConfig.SetEventHandler(handler)
type RtmEventHandler interface {
	// mark interface, for type constraint
	// user must implement this method to identify itself as an event handler, empty implementation is enough
	// this can ensure type safety, prevent passing wrong type
	IsRtmEventHandler()
}

// RtmEventHandlerConfig provide function style event handler config
// user only need to set the needed callback functions, others keep nil
type RtmEventHandlerConfig struct {
	OnLoginResult     func(requestId uint64, errorCode RTM_ERROR_CODE)
	OnLogoutResult    func(requestId uint64, errorCode RTM_ERROR_CODE)
	OnMessageEvent    func(event *MessageEvent)
	OnPresenceEvent   func(event *PresenceEvent)
	OnSubscribeResult func(requestId uint64, channelName string, errorCode RTM_ERROR_CODE)
	OnPublishResult   func(requestId uint64, errorCode RTM_ERROR_CODE)

	OnTopicEvent               func(event *TopicEvent)
	OnLockEvent                func(event *LockEvent)
	OnStorageEvent             func(event *StorageEvent)
	OnConnectionStateChanged   func(channelName string, state RTM_CONNECTION_STATE, reason RTM_CONNECTION_CHANGE_REASON)
	OnTokenPrivilegeWillExpire func(channelName string)
	OnJoinResult               func(requestId uint64, channelName string, userId string, errorCode RTM_ERROR_CODE)
	OnLeaveResult              func(requestId uint64, channelName string, userId string, errorCode RTM_ERROR_CODE)
	OnJoinTopicResult          func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode RTM_ERROR_CODE)
	OnLeaveTopicResult         func(requestId uint64, channelName string, userId string, topic string, meta string, errorCode RTM_ERROR_CODE)
	OnSubscribeTopicResult     func(requestId uint64, channelName string, userId string, topic string, succeedUsers UserList, failedUsers UserList, errorCode RTM_ERROR_CODE)

	OnSetChannelMetadataResult      func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE)
	OnUpdateChannelMetadataResult   func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE)
	OnRemoveChannelMetadataResult   func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, errorCode RTM_ERROR_CODE)
	OnGetChannelMetadataResult      func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, data *IMetadata, errorCode RTM_ERROR_CODE)
	OnSetUserMetadataResult         func(requestId uint64, userId string, errorCode RTM_ERROR_CODE)
	OnUpdateUserMetadataResult      func(requestId uint64, userId string, errorCode RTM_ERROR_CODE)
	OnRemoveUserMetadataResult      func(requestId uint64, userId string, errorCode RTM_ERROR_CODE)
	OnGetUserMetadataResult         func(requestId uint64, userId string, data *IMetadata, errorCode RTM_ERROR_CODE)
	OnSubscribeUserMetadataResult   func(requestId uint64, userId string, errorCode RTM_ERROR_CODE)
	OnUnsubscribeUserMetadataResult func(requestId uint64, userId string, errorCode RTM_ERROR_CODE)

	OnSetLockResult     func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE)
	OnRemoveLockResult  func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE)
	OnReleaseLockResult func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE)
	OnAcquireLockResult func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE, errorDetails string)
	OnRevokeLockResult  func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockName string, errorCode RTM_ERROR_CODE)
	OnGetLocksResult    func(requestId uint64, channelName string, channelType RTM_CHANNEL_TYPE, lockDetailList *LockDetail, count uint, errorCode RTM_ERROR_CODE)

	OnWhoNowResult              func(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode RTM_ERROR_CODE)
	OnGetOnlineUsersResult      func(requestId uint64, userStateList *UserState, count uint, nextPage string, errorCode RTM_ERROR_CODE)
	OnWhereNowResult            func(requestId uint64, channels *ChannelInfo, count uint, errorCode RTM_ERROR_CODE)
	OnGetUserChannelsResult     func(requestId uint64, channels *ChannelInfo, count uint, errorCode RTM_ERROR_CODE)
	OnPresenceSetStateResult    func(requestId uint64, errorCode RTM_ERROR_CODE)
	OnPresenceRemoveStateResult func(requestId uint64, errorCode RTM_ERROR_CODE)
	OnPresenceGetStateResult    func(requestId uint64, state *UserState, errorCode RTM_ERROR_CODE)

	OnLinkStateEvent              func(event *LinkStateEvent)
	OnPublishTopicMessageResult   func(requestId uint64, channelName string, topic string, errorCode RTM_ERROR_CODE)
	OnRenewTokenResult            func(requestId uint64, serverType RTM_SERVICE_TYPE, channelName string, errorCode RTM_ERROR_CODE)
	OnUnsubscribeTopicResult      func(requestId uint64, channelName string, topic string, errorCode RTM_ERROR_CODE)
	OnGetSubscribedUserListResult func(requestId uint64, channelName string, topic string, user *UserList, errorCode RTM_ERROR_CODE)
	OnGetHistoryMessagesResult    func(requestId uint64, messageList []HistoryMessage, newStart uint64, errorCode RTM_ERROR_CODE)
}

// implement RtmEventHandler interface
func (config *RtmEventHandlerConfig) IsRtmEventHandler() {}

// #region MessageEvent
type MessageEvent struct {
	ChannelType   RTM_CHANNEL_TYPE
	MessageType   RTM_MESSAGE_TYPE
	ChannelName   string
	ChannelTopic  string
	Message       []byte
	MessageLength int32
	Publisher     string
	CustomType    string
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *MessageEvent) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *MessageEvent) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

/**
 * Message type
 */
func (this_ *MessageEvent) GetMessageType() RTM_MESSAGE_TYPE {
	return this_.MessageType
}

/**
 * Message type
 */
func (this_ *MessageEvent) SetMessageType(messageType RTM_MESSAGE_TYPE) {
	this_.MessageType = messageType
}

/**
 * The channel which the message was published
 */
func (this_ *MessageEvent) GetChannelName() string {
	return this_.ChannelName
}

/**
 * The channel which the message was published
 */
func (this_ *MessageEvent) SetChannelName(channelName string) {
	this_.ChannelName = channelName
}

/**
 * If the channelType is RTM_CHANNEL_TYPE_STREAM, which topic the message came from. only for RTM_CHANNEL_TYPE_STREAM
 */
func (this_ *MessageEvent) GetChannelTopic() string {
	return this_.ChannelTopic
}

/**
 * If the channelType is RTM_CHANNEL_TYPE_STREAM, which topic the message came from. only for RTM_CHANNEL_TYPE_STREAM
 */
func (this_ *MessageEvent) SetChannelTopic(channelTopic string) {
	this_.ChannelTopic = channelTopic
}

/**
 * The payload
 */
func (this_ *MessageEvent) GetMessage() []byte {
	return this_.Message
}

/**
 * The payload
 */
func (this_ *MessageEvent) SetMessage(message []byte) {
	this_.Message = message
	this_.MessageLength = int32(len(message))
}

/**
 * The payload length
 */
func (this_ *MessageEvent) GetMessageLength() uint {
	return uint(this_.MessageLength)
}

/**
 * The payload length
 */
func (this_ *MessageEvent) SetMessageLength(messageLength uint) {
	this_.MessageLength = int32(messageLength)
}

/**
 * The publisher
 */
func (this_ *MessageEvent) GetPublisher() string {
	return this_.Publisher
}

/**
 * The publisher
 */
func (this_ *MessageEvent) SetPublisher(publisher string) {
	this_.Publisher = publisher
}

/**
 * The custom type of the message
 */
func (this_ *MessageEvent) GetCustomType() string {
	return this_.CustomType
}

/**
 * The publisher
 */
func (this_ *MessageEvent) SetCustomType(customType string) {
	this_.CustomType = customType
}

func NewMessageEvent() *MessageEvent {
	event := &MessageEvent{
		ChannelType:   RTM_CHANNEL_TYPE_NONE,
		MessageType:   RTM_MESSAGE_TYPE_STRING,
		ChannelName:   "",
		ChannelTopic:  "",
		Message:       make([]byte, 0),
		MessageLength: 0,
		Publisher:     "",
		CustomType:    "",
	}

	return event
}

func (this_ *MessageEvent) fromC(cEvent *C.struct_C_MessageEvent) {
	if cEvent == nil {
		return
	}

	if !IsValidMemory(unsafe.Pointer(cEvent)) {
		return
	}

	this_.ChannelType = RTM_CHANNEL_TYPE(cEvent.channelType)
	this_.MessageType = RTM_MESSAGE_TYPE(cEvent.messageType)

	if cEvent.channelName != nil {
		this_.ChannelName = C.GoString(cEvent.channelName)
	}
	if cEvent.channelTopic != nil {
		this_.ChannelTopic = C.GoString(cEvent.channelTopic)
	}
	if cEvent.publisher != nil {
		this_.Publisher = C.GoString(cEvent.publisher)
	}
	if cEvent.customType != nil {
		this_.CustomType = C.GoString(cEvent.customType)
	}

	if cEvent.message != nil && cEvent.messageLength > 0 {
		if IsValidMemory(unsafe.Pointer(cEvent.message)) {
			messageSize := int(cEvent.messageLength)
			if messageSize > 0 {
				this_.Message = make([]byte, messageSize)
				for i := 0; i < messageSize; i++ {
					this_.Message[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.message)) + uintptr(i)))
				}
				this_.MessageLength = int32(messageSize)
			} else {
				this_.Message = make([]byte, 0)
				this_.MessageLength = 0
			}
		} else {
			this_.Message = make([]byte, 0)
			this_.MessageLength = 0
		}
	} else {
		this_.Message = make([]byte, 0)
		this_.MessageLength = 0
	}
}

// #endregion MessageEvent

type IntervalInfo struct {
	JoinUserList    *UserList
	LeaveUserList   *UserList
	TimeoutUserList *UserList
	UserStateList   []*UserState
	UserStateCount  uint
}

// #region IntervalInfo

/**
 * Joined users during this interval
 */
func (this_ *IntervalInfo) GetJoinUserList() *UserList {
	return this_.JoinUserList
}

/**
 * Joined users during this interval
 */
func (this_ *IntervalInfo) SetJoinUserList(joinUserList *UserList) {
	this_.JoinUserList = joinUserList
}

/**
 * Left users during this interval
 */
func (this_ *IntervalInfo) GetLeaveUserList() *UserList {
	return this_.LeaveUserList
}

/**
 * Left users during this interval
 */
func (this_ *IntervalInfo) SetLeaveUserList(leaveUserList *UserList) {
	this_.LeaveUserList = leaveUserList
}

/**
 * Timeout users during this interval
 */
func (this_ *IntervalInfo) GetTimeoutUserList() *UserList {
	return this_.TimeoutUserList
}

/**
 * Timeout users during this interval
 */
func (this_ *IntervalInfo) SetTimeoutUserList(timeoutUserList *UserList) {
	this_.TimeoutUserList = timeoutUserList
}

/**
 * The user state changed during this interval
 */
func (this_ *IntervalInfo) GetUserStateList() []*UserState {
	return this_.UserStateList
}

/**
 * The user state changed during this interval
 */
func (this_ *IntervalInfo) SetUserStateList(userStateList []*UserState) {
	this_.UserStateList = userStateList
}

/**
 * The user count
 */
func (this_ *IntervalInfo) GetUserStateCount() uint {
	return this_.UserStateCount
}

/**
 * The user count
 */
func (this_ *IntervalInfo) SetUserStateCount(userStateCount uint) {
	this_.UserStateCount = userStateCount
}

func NewIntervalInfo() *IntervalInfo {
	info := &IntervalInfo{
		JoinUserList:    NewUserList(),
		LeaveUserList:   NewUserList(),
		TimeoutUserList: NewUserList(),
		UserStateList:   make([]*UserState, 0),
		UserStateCount:  0,
	}

	return info
}

// #endregion IntervalInfo

type SnapshotInfo struct {
	UserStateList []*UserState
	UserCount     uint
}

// #region SnapshotInfo

/**
 * The user state in this snapshot event
 */
func (this_ *SnapshotInfo) GetUserStateList() []*UserState {
	return this_.UserStateList
}

/**
 * The user state in this snapshot event
 */
func (this_ *SnapshotInfo) SetUserStateList(userStateList []*UserState) {
	this_.UserStateList = userStateList
}

/**
 * The user count
 */
func (this_ *SnapshotInfo) GetUserCount() uint {
	return this_.UserCount
}

/**
 * The user count
 */
func (this_ *SnapshotInfo) SetUserCount(userCount uint) {
	this_.UserCount = userCount
}

func NewSnapshotInfo() *SnapshotInfo {
	info := &SnapshotInfo{
		UserStateList: make([]*UserState, 0),
		UserCount:     0,
	}

	return info
}

// #endregion SnapshotInfo

type PresenceEvent struct {
	Type           RTM_PRESENCE_EVENT_TYPE
	ChannelType    RTM_CHANNEL_TYPE
	ChannelName    string
	Publisher      string
	StateItems     []*StateItem
	StateItemCount uint
	Interval       *IntervalInfo
	Snapshot       *SnapshotInfo
}

// #region PresenceEvent

/**
 * Indicate presence event type
 */
func (this_ *PresenceEvent) GetType() RTM_PRESENCE_EVENT_TYPE {
	return this_.Type
}

/**
 * Indicate presence event type
 */
func (this_ *PresenceEvent) SetType(_type RTM_PRESENCE_EVENT_TYPE) {
	this_.Type = _type
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *PresenceEvent) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *PresenceEvent) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

/**
 * The channel which the presence event was triggered
 */
func (this_ *PresenceEvent) GetChannelName() string {
	return this_.ChannelName
}

/**
 * The channel which the presence event was triggered
 */
func (this_ *PresenceEvent) SetChannelName(channelName string) {
	this_.ChannelName = channelName
}

/**
 * The user who triggered this event.
 */
func (this_ *PresenceEvent) GetPublisher() string {
	return this_.Publisher
}

/**
 * The user who triggered this event.
 */
func (this_ *PresenceEvent) SetPublisher(publisher string) {
	this_.Publisher = publisher
}

/**
 * The user states
 */
func (this_ *PresenceEvent) GetStateItems() []*StateItem {
	return this_.StateItems
}

/**
 * The user states
 */
func (this_ *PresenceEvent) SetStateItems(stateItems []*StateItem) {
	this_.StateItems = stateItems
}

/**
 * The states count
 */
func (this_ *PresenceEvent) GetStateItemCount() uint {
	return this_.StateItemCount
}

/**
 * The states count
 */
func (this_ *PresenceEvent) SetStateItemCount(stateItemCount uint) {
	this_.StateItemCount = stateItemCount
}

/**
 * Only valid when in interval mode
 */
func (this_ *PresenceEvent) GetInterval() *IntervalInfo {
	return this_.Interval
}

/**
 * Only valid when in interval mode
 */
func (this_ *PresenceEvent) SetInterval(interval *IntervalInfo) {
	this_.Interval = interval
}

/**
 * Only valid when receive snapshot event
 */
func (this_ *PresenceEvent) GetSnapshot() *SnapshotInfo {
	return this_.Snapshot
}

/**
 * Only valid when in interval mode
 */
func (this_ *PresenceEvent) SetSnapshot(snapshot *SnapshotInfo) {
	this_.Snapshot = snapshot
}

func NewPresenceEvent() *PresenceEvent {
	event := &PresenceEvent{
		Type:           RTM_PRESENCE_EVENT_TYPE_NONE,
		ChannelType:    RTM_CHANNEL_TYPE_NONE,
		ChannelName:    "",
		Publisher:      "",
		StateItems:     make([]*StateItem, 0),
		StateItemCount: 0,
		Interval:       nil,
		Snapshot:       nil,
	}

	return event
}

func (this_ *PresenceEvent) fromC(cEvent *C.struct_C_PresenceEvent) {
	if cEvent == nil {
		return
	}

	if !IsValidMemory(unsafe.Pointer(cEvent)) {
		return
	}

	this_.Type = RTM_PRESENCE_EVENT_TYPE(cEvent._type)
	this_.ChannelType = RTM_CHANNEL_TYPE(cEvent.channelType)

	if cEvent.channelName != nil {
		this_.ChannelName = C.GoString(cEvent.channelName)
	}
	if cEvent.publisher != nil {
		this_.Publisher = C.GoString(cEvent.publisher)
	}

	if cEvent.stateItems != nil && cEvent.stateItemCount > 0 {
		if IsValidMemory(unsafe.Pointer(cEvent.stateItems)) {
			itemCount := int(cEvent.stateItemCount)
			if itemCount > 0 {
				this_.StateItems = make([]*StateItem, itemCount)
				this_.StateItemCount = uint(itemCount)

				for i := 0; i < itemCount; i++ {
					cItem := (*C.struct_C_StateItem)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.stateItems)) + uintptr(i)*unsafe.Sizeof(C.struct_C_StateItem{})))
					if cItem != nil && IsValidMemory(unsafe.Pointer(cItem)) {
						stateItem := NewStateItem()
						if stateItem != nil {
							if cItem.key != nil {
								stateItem.Key = FastSafeCGoString(cItem.key)
							}
							if cItem.value != nil {
								stateItem.Value = FastSafeCGoString(cItem.value)
							}
							this_.StateItems[i] = stateItem
						}
					}
				}
			} else {
				this_.StateItems = make([]*StateItem, 0)
				this_.StateItemCount = 0
			}
		} else {
			this_.StateItems = make([]*StateItem, 0)
			this_.StateItemCount = 0
		}
	} else {
		this_.StateItems = make([]*StateItem, 0)
		this_.StateItemCount = 0
	}

	this_.Interval = NewIntervalInfo()
	if this_.Interval != nil {
		if cEvent.interval.userStateList != nil && cEvent.interval.userStateCount > 0 {
			if IsValidMemory(unsafe.Pointer(cEvent.interval.userStateList)) {
				userCount := int(cEvent.interval.userStateCount)
				if userCount > 0 {
					this_.Interval.UserStateList = make([]*UserState, userCount)
					this_.Interval.UserStateCount = uint(userCount)

					for i := 0; i < userCount; i++ {
						cUserState := (*C.struct_C_UserState)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.interval.userStateList)) + uintptr(i)*unsafe.Sizeof(C.struct_C_UserState{})))
						if cUserState != nil && IsValidMemory(unsafe.Pointer(cUserState)) {
							userState := NewUserState()
							if userState != nil {
								if cUserState.userId != nil {
									userState.UserId = FastSafeCGoString(cUserState.userId)
								}
								if cUserState.states != nil && cUserState.statesCount > 0 {
									if IsValidMemory(unsafe.Pointer(cUserState.states)) {
										stateCount := int(cUserState.statesCount)
										if stateCount > 0 {
											userState.States = make([]StateItem, stateCount)
											userState.StatesCount = uint(stateCount)

											for j := 0; j < stateCount; j++ {
												cState := (*C.struct_C_StateItem)(unsafe.Pointer(uintptr(unsafe.Pointer(cUserState.states)) + uintptr(j)*unsafe.Sizeof(C.struct_C_StateItem{})))
												if cState != nil && IsValidMemory(unsafe.Pointer(cState)) {
													stateItem := StateItem{}
													if cState.key != nil {
														stateItem.Key = FastSafeCGoString(cState.key)
													}
													if cState.value != nil {
														stateItem.Value = FastSafeCGoString(cState.value)
													}
													userState.States[j] = stateItem
												}
											}
										}
									}
								}
								this_.Interval.UserStateList[i] = userState
							}
						}
					}
				}
			}
		}
	}

	this_.Snapshot = NewSnapshotInfo()
	if this_.Snapshot != nil {
		if cEvent.snapshot.userStateList != nil && cEvent.snapshot.userCount > 0 {
			if IsValidMemory(unsafe.Pointer(cEvent.snapshot.userStateList)) {
				userCount := int(cEvent.snapshot.userCount)
				if userCount > 0 {
					this_.Snapshot.UserStateList = make([]*UserState, userCount)
					this_.Snapshot.UserCount = uint(userCount)

					for i := 0; i < userCount; i++ {
						cUserState := (*C.struct_C_UserState)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.snapshot.userStateList)) + uintptr(i)*unsafe.Sizeof(C.struct_C_UserState{})))
						if cUserState != nil && IsValidMemory(unsafe.Pointer(cUserState)) {
							userState := NewUserState()
							if userState != nil {
								if cUserState.userId != nil {
									userState.UserId = FastSafeCGoString(cUserState.userId)
								}
								if cUserState.states != nil && cUserState.statesCount > 0 {
									if IsValidMemory(unsafe.Pointer(cUserState.states)) {
										stateCount := int(cUserState.statesCount)
										if stateCount > 0 {
											userState.States = make([]StateItem, stateCount)
											userState.StatesCount = uint(stateCount)

											for j := 0; j < stateCount; j++ {
												cState := (*C.struct_C_StateItem)(unsafe.Pointer(uintptr(unsafe.Pointer(cUserState.states)) + uintptr(j)*unsafe.Sizeof(C.struct_C_StateItem{})))
												if cState != nil && IsValidMemory(unsafe.Pointer(cState)) {
													stateItem := StateItem{}
													if cState.key != nil {
														stateItem.Key = FastSafeCGoString(cState.key)
													}
													if cState.value != nil {
														stateItem.Value = FastSafeCGoString(cState.value)
													}
													userState.States[j] = stateItem
												}
											}
										}
									}
								}
								this_.Snapshot.UserStateList[i] = userState
							}
						}
					}
				}
			}
		}
	}
}

// #endregion PresenceEvent

type TopicEvent struct {
	Type           RTM_TOPIC_EVENT_TYPE
	ChannelName    string
	Publisher      string
	TopicInfos     []*TopicInfo
	TopicInfoCount uint
}

// #region TopicEvent

/**
 * Indicate topic event type
 */
func (this_ *TopicEvent) GetType() RTM_TOPIC_EVENT_TYPE {
	return this_.Type
}

/**
 * Indicate topic event type
 */
func (this_ *TopicEvent) SetType(_type RTM_TOPIC_EVENT_TYPE) {
	this_.Type = _type
}

/**
 * The channel which the topic event was triggered
 */
func (this_ *TopicEvent) GetChannelName() string {
	return this_.ChannelName
}

/**
 * The channel which the topic event was triggered
 */
func (this_ *TopicEvent) SetChannelName(channelName string) {
	this_.ChannelName = channelName
}

/**
 * The user who triggered this event.
 */
func (this_ *TopicEvent) GetPublisher() string {
	return this_.Publisher
}

/**
 * The user who triggered this event.
 */
func (this_ *TopicEvent) SetPublisher(publisher string) {
	this_.Publisher = publisher
}

/**
 * Topic information array.
 */
func (this_ *TopicEvent) GetTopicInfos() []*TopicInfo {
	return this_.TopicInfos
}

/**
 * Topic information array.
 */
func (this_ *TopicEvent) SetTopicInfos(topicInfos []*TopicInfo) {
	this_.TopicInfos = topicInfos
}

/**
 * The count of topicInfos.
 */
func (this_ *TopicEvent) GetTopicInfoCount() uint {
	return this_.TopicInfoCount
}

/**
 * The count of topicInfos.
 */
func (this_ *TopicEvent) SetTopicInfoCount(topicInfoCount uint) {
	this_.TopicInfoCount = topicInfoCount
}

func NewTopicEvent() *TopicEvent {
	event := &TopicEvent{
		Type:           RTM_TOPIC_EVENT_TYPE_NONE,
		ChannelName:    "",
		Publisher:      "",
		TopicInfos:     make([]*TopicInfo, 0),
		TopicInfoCount: 0,
	}

	return event
}

func (this_ *TopicEvent) fromC(cEvent *C.struct_C_TopicEvent) {
	if cEvent == nil {
		return
	}

	if !IsValidMemory(unsafe.Pointer(cEvent)) {
		return
	}

	this_.Type = RTM_TOPIC_EVENT_TYPE(cEvent._type)

	if cEvent.channelName != nil {
		this_.ChannelName = C.GoString(cEvent.channelName)
	}
	if cEvent.publisher != nil {
		this_.Publisher = C.GoString(cEvent.publisher)
	}

	if cEvent.topicInfos != nil && cEvent.topicInfoCount > 0 {
		if IsValidMemory(unsafe.Pointer(cEvent.topicInfos)) {
			infoCount := int(cEvent.topicInfoCount)
			if infoCount > 0 {
				this_.TopicInfos = make([]*TopicInfo, infoCount)
				this_.TopicInfoCount = uint(infoCount)

				for i := 0; i < infoCount; i++ {
					cTopicInfo := (*C.struct_C_TopicInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.topicInfos)) + uintptr(i)*unsafe.Sizeof(C.struct_C_TopicInfo{})))
					if cTopicInfo != nil && IsValidMemory(unsafe.Pointer(cTopicInfo)) {
						topicInfo := NewTopicInfo()
						if topicInfo != nil {
							if cTopicInfo.topic != nil {
								topicInfo.Topic = FastSafeCGoString(cTopicInfo.topic)
							}
							if cTopicInfo.publishers != nil && cTopicInfo.publisherCount > 0 {
								if IsValidMemory(unsafe.Pointer(cTopicInfo.publishers)) {
									pubCount := int(cTopicInfo.publisherCount)
									if pubCount > 0 {
										topicInfo.Publishers = make([]PublisherInfo, pubCount)
										topicInfo.PublisherCount = uint(pubCount)

										for j := 0; j < pubCount; j++ {
											cPublisher := (*C.struct_C_PublisherInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(cTopicInfo.publishers)) + uintptr(j)*unsafe.Sizeof(C.struct_C_PublisherInfo{})))
											if cPublisher != nil && IsValidMemory(unsafe.Pointer(cPublisher)) {
												publisherInfo := PublisherInfo{}
												if cPublisher.publisherUserId != nil {
													publisherInfo.PublisherUserId = FastSafeCGoString(cPublisher.publisherUserId)
												}
												if cPublisher.publisherMeta != nil {
													publisherInfo.PublisherMeta = FastSafeCGoString(cPublisher.publisherMeta)
												}
												topicInfo.Publishers[j] = publisherInfo
											}
										}
									}
								} else {
									topicInfo.Publishers = make([]PublisherInfo, 0)
									topicInfo.PublisherCount = 0
								}
							} else {
								topicInfo.Publishers = make([]PublisherInfo, 0)
								topicInfo.PublisherCount = 0
							}
							this_.TopicInfos[i] = topicInfo
						}
					}
				}
			}
		}
	} else {
		this_.TopicInfos = make([]*TopicInfo, 0)
		this_.TopicInfoCount = 0
	}
}

// #endregion TopicEvent

type LockEvent struct {
	ChannelType    RTM_CHANNEL_TYPE
	EventType      RTM_LOCK_EVENT_TYPE
	ChannelName    string
	LockDetailList []*LockDetail
	Count          uint
}

// #region LockEvent

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *LockEvent) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *LockEvent) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

/**
 * Lock event type, indicate lock states
 */
func (this_ *LockEvent) GetEventType() RTM_LOCK_EVENT_TYPE {
	return this_.EventType
}

/**
 * Lock event type, indicate lock states
 */
func (this_ *LockEvent) SetEventType(eventType RTM_LOCK_EVENT_TYPE) {
	this_.EventType = eventType
}

/**
 * The channel which the lock event was triggered
 */
func (this_ *LockEvent) GetChannelName() string {
	return this_.ChannelName
}

/**
 * The channel which the lock event was triggered
 */
func (this_ *LockEvent) SetChannelName(channelName string) {
	this_.ChannelName = channelName
}

/**
 * The detail information of locks
 */
func (this_ *LockEvent) GetLockDetailList() []*LockDetail {
	return this_.LockDetailList
}

/**
 * The detail information of locks
 */
func (this_ *LockEvent) SetLockDetailList(lockDetailList []*LockDetail) {
	this_.LockDetailList = lockDetailList
}

/**
 * The count of locks
 */
func (this_ *LockEvent) GetCount() uint {
	return this_.Count
}

/**
 * The count of locks
 */
func (this_ *LockEvent) SetCount(count uint) {
	this_.Count = count
}

func NewLockEvent() *LockEvent {
	event := &LockEvent{
		ChannelType:    RTM_CHANNEL_TYPE_NONE,
		EventType:      RTM_LOCK_EVENT_TYPE_NONE,
		ChannelName:    "",
		LockDetailList: make([]*LockDetail, 0),
		Count:          0,
	}

	return event
}

func (this_ *LockEvent) fromC(cEvent *C.struct_C_LockEvent) {
	if cEvent == nil {
		return
	}

	if !IsValidMemory(unsafe.Pointer(cEvent)) {
		return
	}

	this_.ChannelType = RTM_CHANNEL_TYPE(cEvent.channelType)
	this_.EventType = RTM_LOCK_EVENT_TYPE(cEvent.eventType)

	if cEvent.channelName != nil {
		this_.ChannelName = C.GoString(cEvent.channelName)
	}

	if cEvent.lockDetailList != nil && cEvent.count > 0 {
		if IsValidMemory(unsafe.Pointer(cEvent.lockDetailList)) {
			detailCount := int(cEvent.count)
			if detailCount > 0 {
				this_.LockDetailList = make([]*LockDetail, detailCount)
				this_.Count = uint(detailCount)

				for i := 0; i < detailCount; i++ {
					cLockDetail := (*C.struct_C_LockDetail)(unsafe.Pointer(uintptr(unsafe.Pointer(cEvent.lockDetailList)) + uintptr(i)*unsafe.Sizeof(C.struct_C_LockDetail{})))
					if cLockDetail != nil && IsValidMemory(unsafe.Pointer(cLockDetail)) {
						lockDetail := NewLockDetail()
						if lockDetail != nil {
							if cLockDetail.lockName != nil {
								lockDetail.SetLockName(FastSafeCGoString(cLockDetail.lockName))
							}
							if cLockDetail.owner != nil {
								lockDetail.SetOwner(FastSafeCGoString(cLockDetail.owner))
							}
							lockDetail.SetTtl(uint32(cLockDetail.ttl))
							this_.LockDetailList[i] = lockDetail
						}
					}
				}
			}
		}
	} else {
		this_.LockDetailList = make([]*LockDetail, 0)
		this_.Count = 0
	}
}

// #endregion LockEvent

type StorageEvent struct {
	ChannelType RTM_CHANNEL_TYPE
	StorageType RTM_STORAGE_TYPE
	EventType   RTM_STORAGE_EVENT_TYPE
	Target      string
	Data        *IMetadata
}

// #region StorageEvent

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *StorageEvent) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *StorageEvent) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

/**
 * Storage type, RTM_STORAGE_TYPE_USER or RTM_STORAGE_TYPE_CHANNEL
 */
func (this_ *StorageEvent) GetStorageType() RTM_STORAGE_TYPE {
	return this_.StorageType
}

/**
 * Storage type, RTM_STORAGE_TYPE_USER or RTM_STORAGE_TYPE_CHANNEL
 */
func (this_ *StorageEvent) SetStorageType(storageType RTM_STORAGE_TYPE) {
	this_.StorageType = storageType
}

/**
 * Indicate storage event type
 */
func (this_ *StorageEvent) GetEventType() RTM_STORAGE_EVENT_TYPE {
	return this_.EventType
}

/**
 * Indicate storage event type
 */
func (this_ *StorageEvent) SetEventType(eventType RTM_STORAGE_EVENT_TYPE) {
	this_.EventType = eventType
}

/**
 * The target name of user or channel, depends on the RTM_STORAGE_TYPE
 */
func (this_ *StorageEvent) GetTarget() string {
	return this_.Target
}

/**
 * The target name of user or channel, depends on the RTM_STORAGE_TYPE
 */
func (this_ *StorageEvent) SetTarget(target string) {
	this_.Target = target
}

/**
 * The metadata information
 */
func (this_ *StorageEvent) GetData() *IMetadata {
	return this_.Data
}

/**
 * The metadata information
 */
func (this_ *StorageEvent) SetData(data *IMetadata) {
	this_.Data = data
}

func NewStorageEvent() *StorageEvent {
	event := &StorageEvent{
		ChannelType: RTM_CHANNEL_TYPE_NONE,
		StorageType: RTM_STORAGE_TYPE_NONE,
		EventType:   RTM_STORAGE_EVENT_TYPE_NONE,
		Target:      "",
		Data:        nil,
	}

	return event
}

func (this_ *StorageEvent) fromC(cEvent *C.struct_C_StorageEvent) {
	if cEvent == nil {
		return
	}

	if !IsValidMemory(unsafe.Pointer(cEvent)) {
		return
	}

	this_.ChannelType = RTM_CHANNEL_TYPE(cEvent.channelType)
	this_.StorageType = RTM_STORAGE_TYPE(cEvent.storageType)
	this_.EventType = RTM_STORAGE_EVENT_TYPE(cEvent.eventType)

	if cEvent.target != nil {
		this_.Target = C.GoString(cEvent.target)
	}

	if cEvent.data != nil {
		this_.Data = CMetadataToIMetadata(cEvent.data)
	}
}

// #endregion StorageEvent

/**
 * The IRtmClient class.
 *
 * This class provides the main methods that can be invoked by your app.
 *
 * IRtmClient is the basic interface class of the Agora RTM SDK.
 * Creating an IRtmClient object and then calling the methods of
 * this object enables you to use Agora RTM SDK's functionality.
 */

type IRtmClient struct {
	rtmClient unsafe.Pointer
	adapter   *EventHandlerAdapter
	bridge    *RtmEventHandlerBridge
}

// #region IRtmClient

/**
 * Initializes the rtm client instance.
 *
 * @param [in] config The configurations for RTM Client.
 * @param [in] eventHandler .
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
/**
 * Creates the rtm client object and returns the pointer.
 *
 * @return Pointer of the rtm client object.
 */
func CreateAgoraRtmClient(config *RtmConfig) *IRtmClient {
	if config == nil {
		return nil
	}

	cConfig := C.C_RtmConfig_New()
	if cConfig == nil {
		return nil
	}
	defer C.C_RtmConfig_Delete(cConfig)

	cConfig.appId = C.CString(config.AppId)
	defer C.free(unsafe.Pointer(cConfig.appId))
	cConfig.userId = C.CString(config.UserId)
	defer C.free(unsafe.Pointer(cConfig.userId))
	cConfig.areaCode = C.enum_C_RTM_AREA_CODE(config.AreaCode)
	cConfig.protocolType = config.ProtocolType
	cConfig.presenceTimeout = C.uint32_t(config.PresenceTimeout)
	cConfig.heartbeatInterval = C.uint32_t(config.HeartbeatInterval)
	cConfig.multipath = C.bool(config.Multipath)
	cConfig.context = config.Context
	cConfig.useStringUserId = C.bool(config.UseStringUserId)

	var adapter *EventHandlerAdapter
	var bridge *RtmEventHandlerBridge
	if config.EventHandler != nil {
		adapter = NewEventHandlerAdapter(config.EventHandler)
		bridge = NewRtmEventHandlerBridge(adapter)
		if bridge != nil {
			cConfig.eventHandler = unsafe.Pointer(bridge.cBridge)
		} else {
			cConfig.eventHandler = nil
		}
	} else {
		cConfig.eventHandler = nil
	}

	var errorCode C.int
	rtmClient := C.agora_rtm_client_create(cConfig, &errorCode)

	if rtmClient == nil {
		if bridge != nil {
			bridge.Delete()
		}
		return nil
	}

	return &IRtmClient{
		rtmClient: rtmClient,
		adapter:   adapter,
		bridge:    bridge,
	}
}

/**
 * Release the rtm client instance.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Release() int {
	ret := int(C.agora_rtm_client_release(this_.rtmClient))
	
	if this_.bridge != nil {
		this_.bridge.Delete()
		this_.bridge = nil
	}
	this_.adapter = nil
	this_.rtmClient = nil
	
	return ret
}

/**
 * Login the Agora RTM service. The operation result will be notified by \ref agora::rtm::IRtmEventHandler::onLoginResult callback.
 *
 * @param [in] token Token used to login RTM service.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Login(token string) int {
	var requestId uint64
	cToken := C.CString(token)
	defer C.free(unsafe.Pointer(cToken))
	ret := int(C.agora_rtm_client_login(this_.rtmClient,
		cToken,
		(*C.uint64_t)(unsafe.Pointer(&requestId)),
	))
	return int(ret)
}

/**
 * Logout the Agora RTM service. Be noticed that this method will break the rtm service including storage/lock/presence.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Logout() int {
	var requestId uint64
	return int(C.agora_rtm_client_logout(this_.rtmClient,
		(*C.uint64_t)(unsafe.Pointer(&requestId)),
	))
}

/**
 * Get the storage instance.
 *
 * @return
 * - return NULL if error occurred
 */
func (this_ *IRtmClient) GetStorage() *IRtmStorage {
	cStorage := C.agora_rtm_client_get_storage(this_.rtmClient)
	if cStorage == nil {
		return nil
	}
	return &IRtmStorage{rtmStorage: unsafe.Pointer(cStorage)}
}

/**
 * Get the lock instance.
 *
 * @return
 * - return NULL if error occurred
 */
func (this_ *IRtmClient) GetLock() *IRtmLock {
	cLock := C.agora_rtm_client_get_lock(this_.rtmClient)
	if cLock == nil {
		return nil
	}
	return &IRtmLock{rtmLock: unsafe.Pointer(cLock)}
}

/**
 * Get the presence instance.
 *
 * @return
 * - return NULL if error occurred
 */
func (this_ *IRtmClient) GetPresence() *IRtmPresence {
	cPresence := C.agora_rtm_client_get_presence(this_.rtmClient)
	if cPresence == nil {
		return nil
	}
	return &IRtmPresence{rtmPresence: unsafe.Pointer(cPresence)}
}

/**
 * Get the history instance.
 *
 * @return
 * - return NULL if error occurred
 */
func (this_ *IRtmClient) GetHistory() *IRtmHistory {
	cHistory := C.agora_rtm_client_get_history(this_.rtmClient)
	if cHistory == nil {
		return nil
	}
	return &IRtmHistory{rtmHistory: unsafe.Pointer(cHistory)}
}

/**
 * Renews the token. Once a token is enabled and used, it expires after a certain period of time.
 * You should generate a new token on your server, call this method to renew it.
 *
 * @param [in] token Token used renew.
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) RenewToken(token string) int {
	cToken := C.CString(token)
	defer C.free(unsafe.Pointer(cToken))
	var requestId uint64
	ret := int(C.agora_rtm_client_renew_token(this_.rtmClient,
		cToken,
		(*C.uint64_t)(unsafe.Pointer(&requestId)),
	))
	return int(ret)
}

/**
 * Publish a message in the channel.
 *
 * @param [in] channelName The name of the channel.
 * @param [in] message The content of the message.
 * @param [in] length The length of the message.
 * @param [in] option The option of the message.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Publish(channelName string, message []byte, length uint, option *PublishOptions, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	defer C.free(unsafe.Pointer(cChannelName))
	cMessage := C.CBytes(message)
	defer C.free(unsafe.Pointer(cMessage))
	var cOption unsafe.Pointer
	if option != nil {
		cOption = option.toC()
		defer freePublishOptions(cOption)
	} else {
		cOption = NewPublishOptions().toC()
		defer freePublishOptions(cOption)
	}
	ret := int(C.agora_rtm_client_publish(this_.rtmClient,
		cChannelName,
		(*C.char)(cMessage),
		C.size_t(length),
		(*C.struct_C_PublishOptions)(cOption),
		(*C.uint64_t)(requestId),
	))
	return ret
}

/**
 * Send a message to a channel.
 *
 * @param [in] channelName The name of the channel.
 * @param [in] message The content of the message.
 * @param [in] length The length of the message.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
*/
func (this_ *IRtmClient) SendChannelMessage(channelName string, message []byte, length uint, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	defer C.free(unsafe.Pointer(cChannelName))
	cMessage := C.CBytes(message)
	defer C.free(unsafe.Pointer(cMessage))
	opt := NewPublishOptions()
	opt.SetChannelType(RTM_CHANNEL_TYPE_MESSAGE)
	opt.SetMessageType(RTM_MESSAGE_TYPE_BINARY)

	cOption := opt.toC()
	defer freePublishOptions(cOption)

	ret := int(C.agora_rtm_client_publish(this_.rtmClient,
		cChannelName,
		(*C.char)(cMessage),
		C.size_t(length),
		(*C.struct_C_PublishOptions)(cOption),
		(*C.uint64_t)(requestId),
	))
	return ret
}

/**
 * Send a message to a user.
 *
 * @param [in] userId The id of the user.
 * @param [in] message The content of the message.
 * @param [in] length The length of the message.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
*/
func (this_ *IRtmClient) SendUserMessage(userId string, message []byte, length uint, requestId *uint64) int {
	cUserId := C.CString(userId)
	defer C.free(unsafe.Pointer(cUserId))
	cMessage := C.CBytes(message)
	defer C.free(unsafe.Pointer(cMessage))
	opt := NewPublishOptions()
	opt.SetChannelType(RTM_CHANNEL_TYPE_USER)
	opt.SetMessageType(RTM_MESSAGE_TYPE_BINARY)

	cOption := opt.toC()
	defer freePublishOptions(cOption)

	ret := int(C.agora_rtm_client_publish(this_.rtmClient,
		cUserId,
		(*C.char)(cMessage),
		C.size_t(length),
		(*C.struct_C_PublishOptions)(cOption),
		(*C.uint64_t)(requestId),
	))
	return ret
}

/**
 * Subscribe a channel.
 *
 * @param [in] channelName The name of the channel.
 * @param [in] options The options of subscribe the channel.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Subscribe(channelName string, option *SubscribeOptions, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	defer C.free(unsafe.Pointer(cChannelName))
	cOption := option.toC()
	defer freeSubscribeOptions(cOption)
	ret := int(C.agora_rtm_client_subscribe(this_.rtmClient,
		cChannelName,
		(*C.struct_C_SubscribeOptions)(cOption),
		(*C.uint64_t)(requestId),
	))
	return ret
}

/**
 * Unsubscribe a channel.
 *
 * @param [in] channelName The name of the channel.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) Unsubscribe(channelName string) int {
	cChannelName := C.CString(channelName)
	defer C.free(unsafe.Pointer(cChannelName))
	var requestId uint64
	ret := int(C.agora_rtm_client_unsubscribe(this_.rtmClient,
		cChannelName,
		(*C.uint64_t)(unsafe.Pointer(&requestId)),
	))
	return int(ret)
}

/**
 * Create a stream channel instance.
 *
 * @param [in] channelName The Name of the channel.
 * @return
 * - return NULL if error occurred
 */
func (this_ *IRtmClient) CreateStreamChannel(channelName string) *IStreamChannel {
	cChannelName := C.CString(channelName)
	defer C.free(unsafe.Pointer(cChannelName))

	var errorCode C.int
	ret := C.agora_rtm_client_create_stream_channel(this_.rtmClient,
		cChannelName,
		&errorCode,
	)

	if ret == nil || errorCode != 0 {
		return nil
	}

	streamChannel := &IStreamChannel{streamChannel: unsafe.Pointer(ret)}
	return streamChannel
}

/**
 * Set parameters of the sdk or engine
 *
 * @param [in] parameters The parameters in json format
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmClient) SetParameters(parameters string) int {
	cParameters := C.CString(parameters)
	defer C.free(unsafe.Pointer(cParameters))

	ret := int(C.agora_rtm_client_set_parameters(this_.rtmClient,
		cParameters,
	))
	return ret
}

// #endregion IRtmClient

/**
 * Convert error code to error string
 *
 * @param [in] errorCode Received error code
 * @return The error reason
 */
func GetErrorReason(errorCode int) string {
	return C.GoString(C.agora_rtm_client_get_error_reason(C.int(errorCode)))
}

/**
 * Get the version info of the Agora RTM SDK.
 *
 * @return The version info of the Agora RTM SDK.
 */
func GetVersion() string {
	return C.GoString(C.agora_rtm_client_get_version())
}

// #endregion agora::rtm

// #endregion agora
