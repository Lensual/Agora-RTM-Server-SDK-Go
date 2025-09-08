package agorartm

/*
//引入Agora C封装
#cgo linux CFLAGS: -I${SRCDIR}/../../agora_sdk/agora_rtm_sdk_c/include -I${SRCDIR}/../../agora_sdk/agora_rtm_sdk_c/agora_rtm_sdk/high_level_api/include
#cgo darwin CFLAGS: -I${SRCDIR}/../../agora_sdk_mac/agora_rtm_sdk_c/include -I${SRCDIR}/../../agora_sdk_mac/agora_rtm_sdk_c/agora_rtm_sdk/high_level_api/include

//链接AgoraRTM SDK
#cgo linux LDFLAGS: -L${SRCDIR}/../../agora_sdk -lagora_rtm_sdk -laosl -lagora_rtm_sdk_c
#cgo darwin LDFLAGS: -L${SRCDIR}/../../agora_sdk_mac -lAgoraRtmKit -laosl -lagora_rtm_sdk_c
#include <stdlib.h>
#include "C_AgoraRtmBase.h"
*/
import "C"
import "unsafe"

// #region agora

// #region agora::rtm

const DEFAULT_LOG_SIZE_IN_KB = C.C_DEFAULT_LOG_SIZE_IN_KB

/**
 * IP areas.
 */
type RTM_AREA_CODE C.enum_C_RTM_AREA_CODE

const (
	/**
	 * Mainland China.
	 */
	RTM_AREA_CODE_CN RTM_AREA_CODE = C.RTM_AREA_CODE_CN
	/**
	 * North America.
	 */
	RTM_AREA_CODE_NA RTM_AREA_CODE = C.RTM_AREA_CODE_NA
	/**
	 * Europe.
	 */
	RTM_AREA_CODE_EU RTM_AREA_CODE = C.RTM_AREA_CODE_EU
	/**
	 * Asia, excluding Mainland China.
	 */
	RTM_AREA_CODE_AS RTM_AREA_CODE = C.RTM_AREA_CODE_AS
	/**
	 * Japan.
	 */
	RTM_AREA_CODE_JP RTM_AREA_CODE = C.RTM_AREA_CODE_JP
	/**
	 * India.
	 */
	RTM_AREA_CODE_IN RTM_AREA_CODE = C.RTM_AREA_CODE_IN
	/**
	 * (Default) Global.
	 */
	RTM_AREA_CODE_GLOB RTM_AREA_CODE = C.RTM_AREA_CODE_GLOB
)

/**
 * The log level for rtm sdk.
 */
type RTM_LOG_LEVEL C.enum_C_RTM_LOG_LEVEL

const (
	/**
	 * 0x0000: No logging.
	 */
	RTM_LOG_LEVEL_NONE RTM_LOG_LEVEL = C.RTM_LOG_LEVEL_NONE
	/**
	 * 0x0001: Informational messages.
	 */
	RTM_LOG_LEVEL_INFO RTM_LOG_LEVEL = C.RTM_LOG_LEVEL_INFO
	/**
	 * 0x0002: Warnings.
	 */
	RTM_LOG_LEVEL_WARN RTM_LOG_LEVEL = C.RTM_LOG_LEVEL_WARN
	/**
	 * 0x0004: Errors.
	 */
	RTM_LOG_LEVEL_ERROR RTM_LOG_LEVEL = C.RTM_LOG_LEVEL_ERROR
	/**
	 * 0x0008: Critical errors that may lead to program termination.
	 */
	RTM_LOG_LEVEL_FATAL RTM_LOG_LEVEL = C.RTM_LOG_LEVEL_FATAL
)

/**
 * The encryption mode.
 */
type RTM_ENCRYPTION_MODE C.enum_C_RTM_ENCRYPTION_MODE

const (
	/**
	 * Disable message encryption.
	 */
	RTM_ENCRYPTION_MODE_NONE RTM_ENCRYPTION_MODE = C.RTM_ENCRYPTION_MODE_NONE
	/**
	 * 128-bit AES encryption, GCM mode.
	 */
	RTM_ENCRYPTION_MODE_AES_128_GCM RTM_ENCRYPTION_MODE = C.RTM_ENCRYPTION_MODE_AES_128_GCM
	/**
	 * 256-bit AES encryption, GCM mode.
	 */
	RTM_ENCRYPTION_MODE_AES_256_GCM RTM_ENCRYPTION_MODE = C.RTM_ENCRYPTION_MODE_AES_256_GCM
)

/**
 * The error codes of rtm client.
 */
type RTM_ERROR_CODE C.enum_C_RTM_ERROR_CODE

const (
	/**
	 * 0: No error occurs.
	 */
	RTM_ERROR_OK RTM_ERROR_CODE = C.RTM_ERROR_OK

	/**
	 * -10001 ~ -11000 : reserved for generic error.
	 * -10001: The SDK is not initialized.
	 */
	RTM_ERROR_NOT_INITIALIZED RTM_ERROR_CODE = C.RTM_ERROR_NOT_INITIALIZED
	/**
	 * -10002: The user didn't login the RTM system.
	 */
	RTM_ERROR_NOT_LOGIN RTM_ERROR_CODE = C.RTM_ERROR_NOT_LOGIN
	/**
	 * -10003: The app ID is invalid.
	 */
	RTM_ERROR_INVALID_APP_ID RTM_ERROR_CODE = C.RTM_ERROR_INVALID_APP_ID
	/**
	 * -10004: The event handler is invalid.
	 */
	RTM_ERROR_INVALID_EVENT_HANDLER RTM_ERROR_CODE = C.RTM_ERROR_INVALID_EVENT_HANDLER
	/**
	 * -10005: The token is invalid.
	 */
	RTM_ERROR_INVALID_TOKEN RTM_ERROR_CODE = C.RTM_ERROR_INVALID_TOKEN
	/**
	 * -10006: The user ID is invalid.
	 */
	RTM_ERROR_INVALID_USER_ID RTM_ERROR_CODE = C.RTM_ERROR_INVALID_USER_ID
	/**
	 * -10007: The service is not initialized.
	 */
	RTM_ERROR_INIT_SERVICE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_INIT_SERVICE_FAILED
	/**
	 * -10008: The channel name is invalid.
	 */
	RTM_ERROR_INVALID_CHANNEL_NAME RTM_ERROR_CODE = C.RTM_ERROR_INVALID_CHANNEL_NAME
	/**
	 * -10009: The token has expired.
	 */
	RTM_ERROR_TOKEN_EXPIRED RTM_ERROR_CODE = C.RTM_ERROR_TOKEN_EXPIRED
	/**
	 * -10010: There is no server resources now.
	 */
	RTM_ERROR_LOGIN_NO_SERVER_RESOURCES RTM_ERROR_CODE = C.RTM_ERROR_LOGIN_NO_SERVER_RESOURCES
	/**
	 * -10011: The login timeout.
	 */
	RTM_ERROR_LOGIN_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_LOGIN_TIMEOUT
	/**
	 * -10012: The login is rejected by server.
	 */
	RTM_ERROR_LOGIN_REJECTED RTM_ERROR_CODE = C.RTM_ERROR_LOGIN_REJECTED
	/**
	 * -10013: The login is aborted due to unrecoverable error.
	 */
	RTM_ERROR_LOGIN_ABORTED RTM_ERROR_CODE = C.RTM_ERROR_LOGIN_ABORTED
	/**
	 * -10014: The parameter is invalid.
	 */
	RTM_ERROR_INVALID_PARAMETER RTM_ERROR_CODE = C.RTM_ERROR_INVALID_PARAMETER
	/**
	 * -10015: The login is not authorized. Happens user login the RTM system without granted from console.
	 */
	RTM_ERROR_LOGIN_NOT_AUTHORIZED RTM_ERROR_CODE = C.RTM_ERROR_LOGIN_NOT_AUTHORIZED
	/**
	 * -10016: Try to login or join with inconsistent app ID.
	 */
	RTM_ERROR_INCONSISTENT_APPID RTM_ERROR_CODE = C.RTM_ERROR_INCONSISTENT_APPID
	/**
	 * -10017: Already call same request.
	 */
	RTM_ERROR_DUPLICATE_OPERATION RTM_ERROR_CODE = C.RTM_ERROR_DUPLICATE_OPERATION
	/**
	 * -10018: Already call destroy or release, this_instance is forbidden to call any api, please create new instance.
	 */
	RTM_ERROR_INSTANCE_ALREADY_RELEASED RTM_ERROR_CODE = C.RTM_ERROR_INSTANCE_ALREADY_RELEASED
	/**
	 * -10019: The channel type is invalid.
	 */
	RTM_ERROR_INVALID_CHANNEL_TYPE RTM_ERROR_CODE = C.RTM_ERROR_INVALID_CHANNEL_TYPE
	/**
	 * -10020: The encryption parameter is invalid.
	 */
	RTM_ERROR_INVALID_ENCRYPTION_PARAMETER RTM_ERROR_CODE = C.RTM_ERROR_INVALID_ENCRYPTION_PARAMETER
	/**
	 * -10021: The operation is too frequent.
	 */
	RTM_ERROR_OPERATION_RATE_EXCEED_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_OPERATION_RATE_EXCEED_LIMITATION

	/**
	 * -11001 ~ -12000 : reserved for channel error.
	 * -11001: The user has not joined the channel.
	 */
	RTM_ERROR_CHANNEL_NOT_JOINED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_NOT_JOINED
	/**
	 * -11002: The user has not subscribed the channel.
	 */
	RTM_ERROR_CHANNEL_NOT_SUBSCRIBED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_NOT_SUBSCRIBED
	/**
	 * -11003: The topic member count exceeds the limit.
	 */
	RTM_ERROR_CHANNEL_EXCEED_TOPIC_USER_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_EXCEED_TOPIC_USER_LIMITATION
	/**
	 * -11004: The channel is reused in RTC.
	 */
	RTM_ERROR_CHANNEL_IN_REUSE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_IN_REUSE
	/**
	 * -11005: The channel instance count exceeds the limit.
	 */
	RTM_ERROR_CHANNEL_INSTANCE_EXCEED_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INSTANCE_EXCEED_LIMITATION
	/**
	 * -11006: The channel is in error state.
	 */
	RTM_ERROR_CHANNEL_IN_ERROR_STATE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_IN_ERROR_STATE
	/**
	 * -11007: The channel join failed.
	 */
	RTM_ERROR_CHANNEL_JOIN_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_JOIN_FAILED
	/**
	 * -11008: The topic name is invalid.
	 */
	RTM_ERROR_CHANNEL_INVALID_TOPIC_NAME RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INVALID_TOPIC_NAME
	/**
	 * -11009: The message is invalid.
	 */
	RTM_ERROR_CHANNEL_INVALID_MESSAGE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INVALID_MESSAGE
	/**
	 * -11010: The message length exceeds the limit.
	 */
	RTM_ERROR_CHANNEL_MESSAGE_LENGTH_EXCEED_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_MESSAGE_LENGTH_EXCEED_LIMITATION
	/**
	 * -11011: The user list is invalid.
	 */
	RTM_ERROR_CHANNEL_INVALID_USER_LIST RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INVALID_USER_LIST
	/**
	 * -11012: The stream channel is not available.
	 */
	RTM_ERROR_CHANNEL_NOT_AVAILABLE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_NOT_AVAILABLE
	/**
	 * -11013: The topic is not subscribed.
	 */
	RTM_ERROR_CHANNEL_TOPIC_NOT_SUBSCRIBED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_TOPIC_NOT_SUBSCRIBED
	/**
	 * -11014: The topic count exceeds the limit.
	 */
	RTM_ERROR_CHANNEL_EXCEED_TOPIC_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_EXCEED_TOPIC_LIMITATION
	/**
	 * -11015: Join topic failed.
	 */
	RTM_ERROR_CHANNEL_JOIN_TOPIC_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_JOIN_TOPIC_FAILED
	/**
	 * -11016: The topic is not joined.
	 */
	RTM_ERROR_CHANNEL_TOPIC_NOT_JOINED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_TOPIC_NOT_JOINED
	/**
	 * -11017: The topic does not exist.
	 */
	RTM_ERROR_CHANNEL_TOPIC_NOT_EXIST RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_TOPIC_NOT_EXIST
	/**
	 * -11018: The topic meta is invalid.
	 */
	RTM_ERROR_CHANNEL_INVALID_TOPIC_META RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INVALID_TOPIC_META
	/**
	 * -11019: Subscribe channel timeout.
	 */
	RTM_ERROR_CHANNEL_SUBSCRIBE_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_SUBSCRIBE_TIMEOUT
	/**
	 * -11020: Subscribe channel too frequent.
	 */
	RTM_ERROR_CHANNEL_SUBSCRIBE_TOO_FREQUENT RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_SUBSCRIBE_TOO_FREQUENT
	/**
	 * -11021: Subscribe channel failed.
	 */
	RTM_ERROR_CHANNEL_SUBSCRIBE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_SUBSCRIBE_FAILED
	/**
	 * -11022: Unsubscribe channel failed.
	 */
	RTM_ERROR_CHANNEL_UNSUBSCRIBE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_UNSUBSCRIBE_FAILED
	/**
	 * -11023: Encrypt message failed.
	 */
	RTM_ERROR_CHANNEL_ENCRYPT_MESSAGE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_ENCRYPT_MESSAGE_FAILED
	/**
	 * -11024: Publish message failed.
	 */
	RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_FAILED
	/**
	 * -11025: Publish message too frequent.
	 */
	RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_TOO_FREQUENT RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_TOO_FREQUENT
	/**
	 * -11026: Publish message timeout.
	 */
	RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_PUBLISH_MESSAGE_TIMEOUT
	/**
	 * -11027: The connection state is invalid.
	 */
	RTM_ERROR_CHANNEL_NOT_CONNECTED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_NOT_CONNECTED
	/**
	 * -11028: Leave channel failed.
	 */
	RTM_ERROR_CHANNEL_LEAVE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_LEAVE_FAILED
	/**
	 * -11029: The custom type length exceeds the limit.
	 */
	RTM_ERROR_CHANNEL_CUSTOM_TYPE_LENGTH_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_CUSTOM_TYPE_LENGTH_OVERFLOW
	/**
	 * -11030: The custom type is invalid.
	 */
	RTM_ERROR_CHANNEL_INVALID_CUSTOM_TYPE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_INVALID_CUSTOM_TYPE
	/**
	 * -11031: unsupported message type (in MacOS/iOS platform，message only support NSString and NSData)
	 */
	RTM_ERROR_CHANNEL_UNSUPPORTED_MESSAGE_TYPE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_UNSUPPORTED_MESSAGE_TYPE
	/**
	 * -11032: The channel presence is not ready.
	 */
	RTM_ERROR_CHANNEL_PRESENCE_NOT_READY RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_PRESENCE_NOT_READY
	/**
	 * -11033: The destination user of publish message is offline.
	 */
	RTM_ERROR_CHANNEL_RECEIVER_OFFLINE RTM_ERROR_CODE = C.RTM_ERROR_CHANNEL_RECEIVER_OFFLINE

	/**
	 * -12001 ~ -13000 : reserved for storage error.
	 * -12001: The storage operation failed.
	 */
	RTM_ERROR_STORAGE_OPERATION_FAILED RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_OPERATION_FAILED
	/**
	 * -12002: The metadata item count exceeds the limit.
	 */
	RTM_ERROR_STORAGE_METADATA_ITEM_EXCEED_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_METADATA_ITEM_EXCEED_LIMITATION
	/**
	 * -12003: The metadata item is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_METADATA_ITEM RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_METADATA_ITEM
	/**
	 * -12004: The argument in storage operation is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_ARGUMENT RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_ARGUMENT
	/**
	 * -12005: The revision in storage operation is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_REVISION RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_REVISION
	/**
	 * -12006: The metadata length exceeds the limit.
	 */
	RTM_ERROR_STORAGE_METADATA_LENGTH_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_METADATA_LENGTH_OVERFLOW
	/**
	 * -12007: The lock name in storage operation is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_LOCK_NAME RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_LOCK_NAME
	/**
	 * -12008: The lock in storage operation is not acquired.
	 */
	RTM_ERROR_STORAGE_LOCK_NOT_ACQUIRED RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_LOCK_NOT_ACQUIRED
	/**
	 * -12009: The metadata key is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_KEY RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_KEY
	/**
	 * -12010: The metadata value is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_VALUE RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_VALUE
	/**
	 * -12011: The metadata key length exceeds the limit.
	 */
	RTM_ERROR_STORAGE_KEY_LENGTH_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_KEY_LENGTH_OVERFLOW
	/**
	 * -12012: The metadata value length exceeds the limit.
	 */
	RTM_ERROR_STORAGE_VALUE_LENGTH_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_VALUE_LENGTH_OVERFLOW
	/**
	 * -12013: The metadata key already exists.
	 */
	RTM_ERROR_STORAGE_DUPLICATE_KEY RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_DUPLICATE_KEY
	/**
	 * -12014: The revision in storage operation is outdated.
	 */
	RTM_ERROR_STORAGE_OUTDATED_REVISION RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_OUTDATED_REVISION
	/**
	 * -12015: The storage operation performed without subscribing.
	 */
	RTM_ERROR_STORAGE_NOT_SUBSCRIBE RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_NOT_SUBSCRIBE
	/**
	 * -12016: The metadata item is invalid.
	 */
	RTM_ERROR_STORAGE_INVALID_METADATA_INSTANCE RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_INVALID_METADATA_INSTANCE
	/**
	 * -12017: The user count exceeds the limit when try to subscribe.
	 */
	RTM_ERROR_STORAGE_SUBSCRIBE_USER_EXCEED_LIMITATION RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_SUBSCRIBE_USER_EXCEED_LIMITATION
	/**
	 * -12018: The storage operation timeout.
	 */
	RTM_ERROR_STORAGE_OPERATION_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_OPERATION_TIMEOUT
	/**
	 * -12019: The storage service not available.
	 */
	RTM_ERROR_STORAGE_NOT_AVAILABLE RTM_ERROR_CODE = C.RTM_ERROR_STORAGE_NOT_AVAILABLE

	/**
	 * -13001 ~ -14000 : reserved for presence error.
	 * -13001: The user is not connected.
	 */
	RTM_ERROR_PRESENCE_NOT_CONNECTED RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_NOT_CONNECTED
	/**
	 * -13002: The presence is not writable.
	 */
	RTM_ERROR_PRESENCE_NOT_WRITABLE RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_NOT_WRITABLE
	/**
	 * -13003: The argument in presence operation is invalid.
	 */
	RTM_ERROR_PRESENCE_INVALID_ARGUMENT RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_INVALID_ARGUMENT
	/**
	 * -13004: The cached presence state count exceeds the limit.
	 */
	RTM_ERROR_PRESENCE_CACHED_TOO_MANY_STATES RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_CACHED_TOO_MANY_STATES
	/**
	 * -13005: The state count exceeds the limit.
	 */
	RTM_ERROR_PRESENCE_STATE_COUNT_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_STATE_COUNT_OVERFLOW
	/**
	 * -13006: The state key is invalid.
	 */
	RTM_ERROR_PRESENCE_INVALID_STATE_KEY RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_INVALID_STATE_KEY
	/**
	 * -13007: The state value is invalid.
	 */
	RTM_ERROR_PRESENCE_INVALID_STATE_VALUE RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_INVALID_STATE_VALUE
	/**
	 * -13008: The state key length exceeds the limit.
	 */
	RTM_ERROR_PRESENCE_STATE_KEY_SIZE_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_STATE_KEY_SIZE_OVERFLOW
	/**
	 * -13009: The state value length exceeds the limit.
	 */
	RTM_ERROR_PRESENCE_STATE_VALUE_SIZE_OVERFLOW RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_STATE_VALUE_SIZE_OVERFLOW
	/**
	 * -13010: The state key already exists.
	 */
	RTM_ERROR_PRESENCE_STATE_DUPLICATE_KEY RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_STATE_DUPLICATE_KEY
	/**
	 * -13011: The user is not exist.
	 */
	RTM_ERROR_PRESENCE_USER_NOT_EXIST RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_USER_NOT_EXIST
	/**
	 * -13012: The presence operation timeout.
	 */
	RTM_ERROR_PRESENCE_OPERATION_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_OPERATION_TIMEOUT
	/**
	 * -13013: The presence operation failed.
	 */
	RTM_ERROR_PRESENCE_OPERATION_FAILED RTM_ERROR_CODE = C.RTM_ERROR_PRESENCE_OPERATION_FAILED

	/**
	 * -14001 ~ -15000 : reserved for lock error.
	 * -14001: The lock operation failed.
	 */
	RTM_ERROR_LOCK_OPERATION_FAILED RTM_ERROR_CODE = C.RTM_ERROR_LOCK_OPERATION_FAILED
	/**
	 * -14002: The lock operation timeout.
	 */
	RTM_ERROR_LOCK_OPERATION_TIMEOUT RTM_ERROR_CODE = C.RTM_ERROR_LOCK_OPERATION_TIMEOUT
	/**
	 * -14003: The lock operation is performing.
	 */
	RTM_ERROR_LOCK_OPERATION_PERFORMING RTM_ERROR_CODE = C.RTM_ERROR_LOCK_OPERATION_PERFORMING
	/**
	 * -14004: The lock already exists.
	 */
	RTM_ERROR_LOCK_ALREADY_EXIST RTM_ERROR_CODE = C.RTM_ERROR_LOCK_ALREADY_EXIST
	/**
	 * -14005: The lock name is invalid.
	 */
	RTM_ERROR_LOCK_INVALID_NAME RTM_ERROR_CODE = C.RTM_ERROR_LOCK_INVALID_NAME
	/**
	 * -14006: The lock is not acquired.
	 */
	RTM_ERROR_LOCK_NOT_ACQUIRED RTM_ERROR_CODE = C.RTM_ERROR_LOCK_NOT_ACQUIRED
	/**
	 * -14007: Acquire lock failed.
	 */
	RTM_ERROR_LOCK_ACQUIRE_FAILED RTM_ERROR_CODE = C.RTM_ERROR_LOCK_ACQUIRE_FAILED
	/**
	 * -14008: The lock is not exist.
	 */
	RTM_ERROR_LOCK_NOT_EXIST RTM_ERROR_CODE = C.RTM_ERROR_LOCK_NOT_EXIST
	/**
	 * -14009: The lock service is not available.
	 */
	RTM_ERROR_LOCK_NOT_AVAILABLE RTM_ERROR_CODE = C.RTM_ERROR_LOCK_NOT_AVAILABLE
)

/**
 * Connection states between rtm sdk and agora server.
 */
type RTM_CONNECTION_STATE C.enum_C_RTM_CONNECTION_STATE

const (
	/**
	 * 1: The SDK is disconnected with server.
	 */
	RTM_CONNECTION_STATE_DISCONNECTED RTM_CONNECTION_STATE = C.RTM_CONNECTION_STATE_DISCONNECTED
	/**
	 * 2: The SDK is connecting to the server.
	 */
	RTM_CONNECTION_STATE_CONNECTING RTM_CONNECTION_STATE = C.RTM_CONNECTION_STATE_CONNECTING
	/**
	 * 3: The SDK is connected to the server and has joined a channel. You can now publish or subscribe to
	 * a track in the channel.
	 */
	RTM_CONNECTION_STATE_CONNECTED RTM_CONNECTION_STATE = C.RTM_CONNECTION_STATE_CONNECTED
	/**
	 * 4: The SDK keeps rejoining the channel after being disconnected from the channel, probably because of
	 * network issues.
	 */
	RTM_CONNECTION_STATE_RECONNECTING RTM_CONNECTION_STATE = C.RTM_CONNECTION_STATE_RECONNECTING
	/**
	 * 5: The SDK fails to connect to the server or join the channel.
	 */
	RTM_CONNECTION_STATE_FAILED RTM_CONNECTION_STATE = C.RTM_CONNECTION_STATE_FAILED
)

/**
 * Reasons for connection state change.
 */
type RTM_CONNECTION_CHANGE_REASON C.enum_C_RTM_CONNECTION_CHANGE_REASON

const (
	/**
	 * 0: The SDK is connecting to the server.
	 */
	RTM_CONNECTION_CHANGED_CONNECTING RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_CONNECTING
	/**
	 * 1: The SDK has joined the channel successfully.
	 */
	RTM_CONNECTION_CHANGED_JOIN_SUCCESS RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_JOIN_SUCCESS
	/**
	 * 2: The connection between the SDK and the server is interrupted.
	 */
	RTM_CONNECTION_CHANGED_INTERRUPTED RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_INTERRUPTED
	/**
	 * 3: The connection between the SDK and the server is banned by the server.
	 */
	RTM_CONNECTION_CHANGED_BANNED_BY_SERVER RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_BANNED_BY_SERVER
	/**
	 * 4: The SDK fails to join the channel for more than 20 minutes and stops reconnecting to the channel.
	 */
	RTM_CONNECTION_CHANGED_JOIN_FAILED RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_JOIN_FAILED
	/**
	 * 5: The SDK has left the channel.
	 */
	RTM_CONNECTION_CHANGED_LEAVE_CHANNEL RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_LEAVE_CHANNEL
	/**
	 * 6: The connection fails because the App ID is not valid.
	 */
	RTM_CONNECTION_CHANGED_INVALID_APP_ID RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_INVALID_APP_ID
	/**
	 * 7: The connection fails because the channel name is not valid.
	 */
	RTM_CONNECTION_CHANGED_INVALID_CHANNEL_NAME RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_INVALID_CHANNEL_NAME
	/**
	 * 8: The connection fails because the token is not valid.
	 */
	RTM_CONNECTION_CHANGED_INVALID_TOKEN RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_INVALID_TOKEN
	/**
	 * 9: The connection fails because the token has expired.
	 */
	RTM_CONNECTION_CHANGED_TOKEN_EXPIRED RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_TOKEN_EXPIRED
	/**
	 * 10: The connection is rejected by the server.
	 */
	RTM_CONNECTION_CHANGED_REJECTED_BY_SERVER RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_REJECTED_BY_SERVER
	/**
	 * 11: The connection changes to reconnecting because the SDK has set a proxy server.
	 */
	RTM_CONNECTION_CHANGED_SETTING_PROXY_SERVER RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_SETTING_PROXY_SERVER
	/**
	 * 12: When the connection state changes because the app has renewed the token.
	 */
	RTM_CONNECTION_CHANGED_RENEW_TOKEN RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_RENEW_TOKEN
	/**
	 * 13: The IP Address of the app has changed. A change in the network type or IP/Port changes the IP
	 * address of the app.
	 */
	RTM_CONNECTION_CHANGED_CLIENT_IP_ADDRESS_CHANGED RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_CLIENT_IP_ADDRESS_CHANGED
	/**
	 * 14: A timeout occurs for the keep-alive of the connection between the SDK and the server.
	 */
	RTM_CONNECTION_CHANGED_KEEP_ALIVE_TIMEOUT RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_KEEP_ALIVE_TIMEOUT
	/**
	 * 15: The SDK has rejoined the channel successfully.
	 */
	RTM_CONNECTION_CHANGED_REJOIN_SUCCESS RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_REJOIN_SUCCESS
	/**
	 * 16: The connection between the SDK and the server is lost.
	 */
	RTM_CONNECTION_CHANGED_LOST RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_LOST
	/**
	 * 17: The change of connection state is caused by echo test.
	 */
	RTM_CONNECTION_CHANGED_ECHO_TEST RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_ECHO_TEST
	/**
	 * 18: The local IP Address is changed by user.
	 */
	RTM_CONNECTION_CHANGED_CLIENT_IP_ADDRESS_CHANGED_BY_USER RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_CLIENT_IP_ADDRESS_CHANGED_BY_USER
	/**
	 * 19: The connection is failed due to join the same channel on another device with the same uid.
	 */
	RTM_CONNECTION_CHANGED_SAME_UID_LOGIN RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_SAME_UID_LOGIN
	/**
	 * 20: The connection is failed due to too many broadcasters in the channel.
	 */
	RTM_CONNECTION_CHANGED_TOO_MANY_BROADCASTERS RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_TOO_MANY_BROADCASTERS
	/**
	 * 21: The connection is failed due to license validation failure.
	 */
	RTM_CONNECTION_CHANGED_LICENSE_VALIDATION_FAILURE RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_LICENSE_VALIDATION_FAILURE
	/**
	 * 22: The connection is failed due to user vid not support stream channel.
	 */
	RTM_CONNECTION_CHANGED_STREAM_CHANNEL_NOT_AVAILABLE RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_STREAM_CHANNEL_NOT_AVAILABLE
	/**
	 * 23: The connection is failed due to token and appid inconsistent.
	 */
	RTM_CONNECTION_CHANGED_INCONSISTENT_APPID RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_INCONSISTENT_APPID
	/**
	 * 10001: The connection of rtm edge service has been successfully established.
	 */
	RTM_CONNECTION_CHANGED_LOGIN_SUCCESS RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_LOGIN_SUCCESS
	/**
	 * 10002: User log out Agora RTM system.
	 */
	RTM_CONNECTION_CHANGED_LOGOUT RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_LOGOUT
	/**
	 * 10003: User log out Agora RTM system.
	 */
	RTM_CONNECTION_CHANGED_PRESENCE_NOT_READY RTM_CONNECTION_CHANGE_REASON = C.RTM_CONNECTION_CHANGED_PRESENCE_NOT_READY
)

/**
 * RTM channel type.
 */
type RTM_CHANNEL_TYPE C.enum_C_RTM_CHANNEL_TYPE

const (
	/**
	 * 0: Unknown channel type.
	 */
	RTM_CHANNEL_TYPE_NONE RTM_CHANNEL_TYPE = C.RTM_CHANNEL_TYPE_NONE
	/**
	 * 1: Message channel.
	 */
	RTM_CHANNEL_TYPE_MESSAGE RTM_CHANNEL_TYPE = C.RTM_CHANNEL_TYPE_MESSAGE
	/**
	 * 2: Stream channel.
	 */
	RTM_CHANNEL_TYPE_STREAM RTM_CHANNEL_TYPE = C.RTM_CHANNEL_TYPE_STREAM
	/**
	 * 3: User.
	 */
	RTM_CHANNEL_TYPE_USER RTM_CHANNEL_TYPE = C.RTM_CHANNEL_TYPE_USER
)

/*
*
@brief Message type when user publish message to channel or topic
*/
type RTM_MESSAGE_TYPE C.enum_C_RTM_MESSAGE_TYPE

const (
	/**
	  0: The binary message.
	*/
	RTM_MESSAGE_TYPE_BINARY RTM_MESSAGE_TYPE = C.RTM_MESSAGE_TYPE_BINARY
	/**
	  1: The ascii message.
	*/
	RTM_MESSAGE_TYPE_STRING RTM_MESSAGE_TYPE = C.RTM_MESSAGE_TYPE_STRING
)

/*
*
@brief Storage type indicate the storage event was triggered by user or channel
*/
type RTM_STORAGE_TYPE C.enum_C_RTM_STORAGE_TYPE

const (
	/**
	  0: Unknown type.
	*/
	RTM_STORAGE_TYPE_NONE RTM_STORAGE_TYPE = C.RTM_STORAGE_TYPE_NONE
	/**
	  1: The user storage event.
	*/
	RTM_STORAGE_TYPE_USER RTM_STORAGE_TYPE = C.RTM_STORAGE_TYPE_USER
	/**
	  2: The channel storage event.
	*/
	RTM_STORAGE_TYPE_CHANNEL RTM_STORAGE_TYPE = C.RTM_STORAGE_TYPE_CHANNEL
)

/**
 * The storage event type, indicate storage operation
 */
type RTM_STORAGE_EVENT_TYPE C.enum_C_RTM_STORAGE_EVENT_TYPE

const (
	/**
	  0: Unknown event type.
	*/
	RTM_STORAGE_EVENT_TYPE_NONE RTM_STORAGE_EVENT_TYPE = C.RTM_STORAGE_EVENT_TYPE_NONE
	/**
	  1: Triggered when user subscribe user metadata state or join channel with options.withMetadata = true
	*/
	RTM_STORAGE_EVENT_TYPE_SNAPSHOT RTM_STORAGE_EVENT_TYPE = C.RTM_STORAGE_EVENT_TYPE_SNAPSHOT
	/**
	  2: Triggered when a remote user set metadata
	*/
	RTM_STORAGE_EVENT_TYPE_SET RTM_STORAGE_EVENT_TYPE = C.RTM_STORAGE_EVENT_TYPE_SET
	/**
	  3: Triggered when a remote user update metadata
	*/
	RTM_STORAGE_EVENT_TYPE_UPDATE RTM_STORAGE_EVENT_TYPE = C.RTM_STORAGE_EVENT_TYPE_UPDATE
	/**
	  4: Triggered when a remote user remove metadata
	*/
	RTM_STORAGE_EVENT_TYPE_REMOVE RTM_STORAGE_EVENT_TYPE = C.RTM_STORAGE_EVENT_TYPE_REMOVE
)

/**
 * The lock event type, indicate lock operation
 */
type RTM_LOCK_EVENT_TYPE C.enum_C_RTM_LOCK_EVENT_TYPE

const (
	/**
	 * 0: Unknown event type
	 */
	RTM_LOCK_EVENT_TYPE_NONE RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_NONE
	/**
	 * 1: Triggered when user subscribe lock state
	 */
	RTM_LOCK_EVENT_TYPE_SNAPSHOT RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_SNAPSHOT
	/**
	 * 2: Triggered when a remote user set lock
	 */
	RTM_LOCK_EVENT_TYPE_LOCK_SET RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_LOCK_SET
	/**
	 * 3: Triggered when a remote user remove lock
	 */
	RTM_LOCK_EVENT_TYPE_LOCK_REMOVED RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_LOCK_REMOVED
	/**
	 * 4: Triggered when a remote user acquired lock
	 */
	RTM_LOCK_EVENT_TYPE_LOCK_ACQUIRED RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_LOCK_ACQUIRED
	/**
	 * 5: Triggered when a remote user released lock
	 */
	RTM_LOCK_EVENT_TYPE_LOCK_RELEASED RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_LOCK_RELEASED
	/**
	 * 6: Triggered when user reconnect to rtm service,
	 * detect the lock has been acquired and released by others.
	 */
	RTM_LOCK_EVENT_TYPE_LOCK_EXPIRED RTM_LOCK_EVENT_TYPE = C.RTM_LOCK_EVENT_TYPE_LOCK_EXPIRED
)

/**
 * The proxy type
 */
type RTM_PROXY_TYPE C.enum_C_RTM_PROXY_TYPE

const (
	/**
	 * 0: Link without proxy
	 */
	RTM_PROXY_TYPE_NONE RTM_PROXY_TYPE = C.RTM_PROXY_TYPE_NONE
	/**
	 * 1: Link with http proxy
	 */
	RTM_PROXY_TYPE_HTTP RTM_PROXY_TYPE = C.RTM_PROXY_TYPE_HTTP
	/**
	 * 2: Link with tcp cloud proxy
	 */
	RTM_PROXY_TYPE_CLOUD_TCP RTM_PROXY_TYPE = C.RTM_PROXY_TYPE_CLOUD_TCP
)

/*
*
@brief Topic event type
*/
type RTM_TOPIC_EVENT_TYPE C.enum_C_RTM_TOPIC_EVENT_TYPE

const (
	/**
	 * 0: Unknown event type
	 */
	RTM_TOPIC_EVENT_TYPE_NONE RTM_TOPIC_EVENT_TYPE = C.RTM_TOPIC_EVENT_TYPE_NONE
	/**
	 * 1: The topic snapshot of this_channel
	 */
	RTM_TOPIC_EVENT_TYPE_SNAPSHOT RTM_TOPIC_EVENT_TYPE = C.RTM_TOPIC_EVENT_TYPE_SNAPSHOT
	/**
	 * 2: Triggered when remote user join a topic
	 */
	RTM_TOPIC_EVENT_TYPE_REMOTE_JOIN_TOPIC RTM_TOPIC_EVENT_TYPE = C.RTM_TOPIC_EVENT_TYPE_REMOTE_JOIN_TOPIC
	/**
	 * 3: Triggered when remote user leave a topic
	 */
	RTM_TOPIC_EVENT_TYPE_REMOTE_LEAVE_TOPIC RTM_TOPIC_EVENT_TYPE = C.RTM_TOPIC_EVENT_TYPE_REMOTE_LEAVE_TOPIC
)

/*
*
@brief Presence event type
*/
type RTM_PRESENCE_EVENT_TYPE C.enum_C_RTM_PRESENCE_EVENT_TYPE

const (
	/**
	 * 0: Unknown event type
	 */
	RTM_PRESENCE_EVENT_TYPE_NONE RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_NONE
	/**
	 * 1: The presence snapshot of this_channel
	 */
	RTM_PRESENCE_EVENT_TYPE_SNAPSHOT RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_SNAPSHOT
	/**
	 * 2: The presence event triggered in interval mode
	 */
	RTM_PRESENCE_EVENT_TYPE_INTERVAL RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_INTERVAL
	/**
	 * 3: Triggered when remote user join channel
	 */
	RTM_PRESENCE_EVENT_TYPE_REMOTE_JOIN_CHANNEL RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_REMOTE_JOIN_CHANNEL
	/**
	 * 4: Triggered when remote user leave channel
	 */
	RTM_PRESENCE_EVENT_TYPE_REMOTE_LEAVE_CHANNEL RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_REMOTE_LEAVE_CHANNEL
	/**
	 * 5: Triggered when remote user's connection timeout
	 */
	RTM_PRESENCE_EVENT_TYPE_REMOTE_TIMEOUT RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_REMOTE_TIMEOUT
	/**
	 * 6: Triggered when user changed state
	 */
	RTM_PRESENCE_EVENT_TYPE_REMOTE_STATE_CHANGED RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_REMOTE_STATE_CHANGED
	/**
	 * 7: Triggered when user joined channel without presence service
	 */
	RTM_PRESENCE_EVENT_TYPE_ERROR_OUT_OF_SERVICE RTM_PRESENCE_EVENT_TYPE = C.RTM_PRESENCE_EVENT_TYPE_ERROR_OUT_OF_SERVICE
)

/*
*
@brief Rtm link operation
*/
type RTM_LINK_OPERATION C.enum_C_RTM_LINK_OPERATION

const (
	/**
	 * 0: Login.
	 */
	RTM_LINK_OPERATION_LOGIN RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_LOGIN
	/**
	 * 1: Logout.
	 */
	RTM_LINK_OPERATION_LOGOUT RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_LOGOUT
	/**
	 * 2: Join.
	 */
	RTM_LINK_OPERATION_JOIN RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_JOIN
	/**
	 * 3: Server reject.
	 */
	RTM_LINK_OPERATION_SERVER_REJECT RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_SERVER_REJECT
	/**
	 * 4: Auto reconnect.
	 */
	RTM_LINK_OPERATION_AUTO_RECONNECT RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_AUTO_RECONNECT
	/**
	 * 5: Reconnected.
	 */
	RTM_LINK_OPERATION_RECONNECTED RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_RECONNECTED
	/**
	 * 6: Heartbeat lost.
	 */
	RTM_LINK_OPERATION_HEARTBEAT_LOST RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_HEARTBEAT_LOST
	/**
	 * 7: Server timeout.
	 */
	RTM_LINK_OPERATION_SERVER_TIMEOUT RTM_LINK_OPERATION = C.RTM_LINK_OPERATION_SERVER_TIMEOUT
)

type RTM_LINK_STATE C.enum_C_RTM_LINK_STATE

const (
	/**
	 * The initial state.
	 */
	RTM_LINK_STATE_IDLE RTM_LINK_STATE = C.RTM_LINK_STATE_IDLE
	/**
	 * The SDK is connecting to the server.
	 */
	RTM_LINK_STATE_CONNECTING RTM_LINK_STATE = C.RTM_LINK_STATE_CONNECTING
	/**
	 * The SDK has connected to the server.
	 */
	RTM_LINK_STATE_CONNECTED RTM_LINK_STATE = C.RTM_LINK_STATE_CONNECTED
	/**
	 * The SDK is disconnected from the server.
	 */
	RTM_LINK_STATE_DISCONNECTED RTM_LINK_STATE = C.RTM_LINK_STATE_DISCONNECTED
	/**
	 * The SDK link is suspended.
	 */
	RTM_LINK_STATE_SUSPENDED RTM_LINK_STATE = C.RTM_LINK_STATE_SUSPENDED
	/**
	 * The SDK is failed to connect to the server.
	 */
	RTM_LINK_STATE_FAILED RTM_LINK_STATE = C.RTM_LINK_STATE_FAILED
)

type RTM_LINK_STATE_CHANGE_REASON C.enum_C_RTM_LINK_STATE_CHANGE_REASON

const (
	/**
	 * Unknown reason.
	 */
	RTM_LINK_STATE_CHANGE_REASON_UNKNOWN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_UNKNOWN
	/**
	 * Login.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN
	/**
	 * Login success.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN_SUCCESS RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN_SUCCESS
	/**
	 * Login timeout.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN_TIMEOUT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN_TIMEOUT
	/**
	 * Login not authorized.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN_NOT_AUTHORIZED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN_NOT_AUTHORIZED
	/**
	 * Login rejected.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN_REJECTED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN_REJECTED
	/**
	 * Re-login.
	 */
	RTM_LINK_STATE_CHANGE_REASON_RELOGIN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_RELOGIN
	/**
	 * Logout.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGOUT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGOUT
	/**
	 * Auto reconnect.
	 */
	RTM_LINK_STATE_CHANGE_REASON_AUTO_RECONNECT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_AUTO_RECONNECT
	/**
	 * Reconnect timeout.
	 */
	RTM_LINK_STATE_CHANGE_REASON_RECONNECT_TIMEOUT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_RECONNECT_TIMEOUT
	/**
	 * Reconnect success.
	 */
	RTM_LINK_STATE_CHANGE_REASON_RECONNECT_SUCCESS RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_RECONNECT_SUCCESS
	/**
	 * Join.
	 */
	RTM_LINK_STATE_CHANGE_REASON_JOIN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_JOIN
	/**
	 * Join success.
	 */
	RTM_LINK_STATE_CHANGE_REASON_JOIN_SUCCESS RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_JOIN_SUCCESS
	/**
	 * Join failed.
	 */
	RTM_LINK_STATE_CHANGE_REASON_JOIN_FAILED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_JOIN_FAILED
	/**
	 * Rejoin.
	 */
	RTM_LINK_STATE_CHANGE_REASON_REJOIN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_REJOIN
	/**
	 * Leave.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LEAVE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LEAVE
	/**
	 * Invalid token.
	 */
	RTM_LINK_STATE_CHANGE_REASON_INVALID_TOKEN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_INVALID_TOKEN
	/**
	 * Token expired.
	 */
	RTM_LINK_STATE_CHANGE_REASON_TOKEN_EXPIRED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_TOKEN_EXPIRED
	/**
	 * Inconsistent app ID.
	 */
	RTM_LINK_STATE_CHANGE_REASON_INCONSISTENT_APP_ID RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_INCONSISTENT_APP_ID
	/**
	 * Invalid channel name.
	 */
	RTM_LINK_STATE_CHANGE_REASON_INVALID_CHANNEL_NAME RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_INVALID_CHANNEL_NAME
	/**
	 * Invalid user ID.
	 */
	RTM_LINK_STATE_CHANGE_REASON_INVALID_USER_ID RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_INVALID_USER_ID
	/**
	 * Not initialized.
	 */
	RTM_LINK_STATE_CHANGE_REASON_NOT_INITIALIZED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_NOT_INITIALIZED
	/**
	 * Rtm service not connected.
	 */
	RTM_LINK_STATE_CHANGE_REASON_RTM_SERVICE_NOT_CONNECTED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_RTM_SERVICE_NOT_CONNECTED
	/**
	 * Channel instance exceed limitation.
	 */
	RTM_LINK_STATE_CHANGE_REASON_CHANNEL_INSTANCE_EXCEED_LIMITATION RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_CHANNEL_INSTANCE_EXCEED_LIMITATION
	/**
	 * Operation rate exceed limitation.
	 */
	RTM_LINK_STATE_CHANGE_REASON_OPERATION_RATE_EXCEED_LIMITATION RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_OPERATION_RATE_EXCEED_LIMITATION
	/**
	 * Channel in error state.
	 */
	RTM_LINK_STATE_CHANGE_REASON_CHANNEL_IN_ERROR_STATE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_CHANNEL_IN_ERROR_STATE
	/**
	 * Presence not connected.
	 */
	RTM_LINK_STATE_CHANGE_REASON_PRESENCE_NOT_CONNECTED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_PRESENCE_NOT_CONNECTED
	/**
	 * Same UID login.
	 */
	RTM_LINK_STATE_CHANGE_REASON_SAME_UID_LOGIN RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_SAME_UID_LOGIN
	/**
	 * Kicked out by server.
	 */
	RTM_LINK_STATE_CHANGE_REASON_KICKED_OUT_BY_SERVER RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_KICKED_OUT_BY_SERVER
	/**
	 * Keep alive timeout.
	 */
	RTM_LINK_STATE_CHANGE_REASON_KEEP_ALIVE_TIMEOUT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_KEEP_ALIVE_TIMEOUT
	/**
	 * Connection error.
	 */
	RTM_LINK_STATE_CHANGE_REASON_CONNECTION_ERROR RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_CONNECTION_ERROR
	/**
	 * Presence not ready.
	 */
	RTM_LINK_STATE_CHANGE_REASON_PRESENCE_NOT_READY RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_PRESENCE_NOT_READY
	/**
	 * Network change.
	 */
	RTM_LINK_STATE_CHANGE_REASON_NETWORK_CHANGE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_NETWORK_CHANGE
	/**
	 * Service not supported.
	 */
	RTM_LINK_STATE_CHANGE_REASON_SERVICE_NOT_SUPPORTED RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_SERVICE_NOT_SUPPORTED
	/**
	 * Stream channel not available.
	 */
	RTM_LINK_STATE_CHANGE_REASON_STREAM_CHANNEL_NOT_AVAILABLE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_STREAM_CHANNEL_NOT_AVAILABLE
	/**
	 * storage not available.
	 */
	RTM_LINK_STATE_CHANGE_REASON_STORAGE_NOT_AVAILABLE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_STORAGE_NOT_AVAILABLE
	/**
	 * Lock not available.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOCK_NOT_AVAILABLE RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOCK_NOT_AVAILABLE
	/**
	 * Login too frequent.
	 */
	RTM_LINK_STATE_CHANGE_REASON_LOGIN_TOO_FREQUENT RTM_LINK_STATE_CHANGE_REASON = C.RTM_LINK_STATE_CHANGE_REASON_LOGIN_TOO_FREQUENT
)

type RTM_PROTOCOL_TYPE C.enum_C_RTM_PROTOCOL_TYPE

const (
	/**
	 * TCP and UDP (default).
	 */
	RTM_PROTOCOL_TYPE_TCP_UDP RTM_PROTOCOL_TYPE = C.RTM_PROTOCOL_TYPE_TCP_UDP
	/**
	 * Use TCP only.
	 */
	RTM_PROTOCOL_TYPE_TCP_ONLY RTM_PROTOCOL_TYPE = C.RTM_PROTOCOL_TYPE_TCP_ONLY
)

/**
 * Definition of LogConfiguration
 */
type RtmLogConfig struct {
	FilePath     string
	FileSizeInKB uint32
	Level        RTM_LOG_LEVEL
}

// #region RtmLogConfig

/**
 * The log file path, default is NULL for default log path
 */
func (this_ *RtmLogConfig) GetFilePath() string {
	return this_.FilePath
}

/**
 * The log file path, default is NULL for default log path
 */
func (this_ *RtmLogConfig) SetFilePath(filePath string) {
	this_.FilePath = filePath
}

/**
 * The log file size, KB , set 1024KB to use default log size
 */
func (this_ *RtmLogConfig) GetFileSizeInKB() uint32 {
	return this_.FileSizeInKB
}

/**
 * The log file size, KB , set 1024KB to use default log size
 */
func (this_ *RtmLogConfig) SetFileSizeInKB(fileSizeInKB uint32) {
	this_.FileSizeInKB = fileSizeInKB
}

/**
 *  The log level, set LOG_LEVEL_INFO to use default log level
 */
func (this_ *RtmLogConfig) GetLevel() RTM_LOG_LEVEL {
	return this_.Level
}

/**
 *  The log level, set LOG_LEVEL_INFO to use default log level
 */
func (this_ *RtmLogConfig) SetLevel(level RTM_LOG_LEVEL) {
	this_.Level = level
}

func NewRtmLogConfig() *RtmLogConfig {
	config := &RtmLogConfig{
		FileSizeInKB: DEFAULT_LOG_SIZE_IN_KB,
		Level:        RTM_LOG_LEVEL_INFO,
	}

	return config
}

// #endregion RtmLogConfig

/**
 * User list.
 */
type UserList struct {
	Users     []string
	UserCount uint
}

// #region UserList

/**
 * The list of users.
 */
func (this_ *UserList) GetUsers() []string {
	return this_.Users
}

/**
 * The list of users.
 */
func (this_ *UserList) SetUsers(users []string) {
	this_.Users = users
	this_.UserCount = uint(len(users))
}

/**
 * The number of users.
 */
func (this_ *UserList) GetUserCount() uint {
	return this_.UserCount
}

/**
 * The number of users.
 */
func (this_ *UserList) SetUserCount(userCount uint) {
	this_.UserCount = userCount
}

func NewUserList() *UserList {
	userList := &UserList{
		Users:     make([]string, 0),
		UserCount: 0,
	}

	return userList
}

// #endregion UserList

/*
*
@brief Topic publisher information
*/
type PublisherInfo struct {
	PublisherUserId string
	PublisherMeta   string
}

// #region PublisherInfo

/**
 * The publisher user ID
 */
func (this_ *PublisherInfo) GetPublisherUserId() string {
	return this_.PublisherUserId
}

/**
 * The publisher user ID
 */
func (this_ *PublisherInfo) SetPublisherUserId(publisherUserId string) {
	this_.PublisherUserId = publisherUserId
}

/**
 * The metadata of the publisher
 */
func (this_ *PublisherInfo) GetPublisherMeta() string {
	return this_.PublisherMeta
}

/**
 * The metadata of the publisher
 */
func (this_ *PublisherInfo) SetPublisherMeta(publisherMeta string) {
	this_.PublisherMeta = publisherMeta
}

func NewPublisherInfo() *PublisherInfo {
	publisher := &PublisherInfo{
		PublisherUserId: "",
		PublisherMeta:   "",
	}

	return publisher
}

// #endregion PublisherInfo

/*
*
@brief Topic information
*/
type TopicInfo struct {
	Topic          string
	Publishers     []PublisherInfo
	PublisherCount uint
}

// #region TopicInfo

/**
 * The name of the topic
 */
func (this_ *TopicInfo) GetTopic() string {
	return this_.Topic
}

/**
 * The name of the topic
 */
func (this_ *TopicInfo) SetTopic(topic string) {
	this_.Topic = topic
}

/**
 * The publisher array
 */
func (this_ *TopicInfo) GetPublishers() []PublisherInfo {
	return this_.Publishers
}

/**
 * The publisher array
 */
func (this_ *TopicInfo) SetPublishers(publishers []PublisherInfo) {
	this_.Publishers = publishers
	this_.PublisherCount = uint(len(publishers))
}

/**
 * The count of publisher in current topic
 */
func (this_ *TopicInfo) GetPublisherCount() uint {
	return this_.PublisherCount
}

/**
 * The count of publisher in current topic
 */
func (this_ *TopicInfo) SetPublisherCount(publisherCount uint) {
	this_.PublisherCount = publisherCount
}

func NewTopicInfo() *TopicInfo {
	topicInfo := &TopicInfo{
		Topic:          "",
		Publishers:     make([]PublisherInfo, 0),
		PublisherCount: 0,
	}

	return topicInfo
}

// #endregion TopicInfo

/*
*
@brief User state property
*/
type StateItem struct {
	Key   string
	Value string
}

// #region StateItem

/**
 * The key of the state item.
 */
func (this_ *StateItem) GetKey() string {
	return this_.Key
}

/**
 * The key of the state item.
 */
func (this_ *StateItem) SetKey(key string) {
	this_.Key = key
}

/**
 * The value of the state item.
 */
func (this_ *StateItem) GetValue() string {
	return this_.Value
}

/**
 * The value of the state item.
 */
func (this_ *StateItem) SetValue(value string) {
	this_.Value = value
}

func NewStateItem() *StateItem {
	stateItem := &StateItem{
		Key:   "",
		Value: "",
	}

	return stateItem
}

// #endregion StateItem

/**
*  The information of a Lock.
 */
type LockDetail struct {
	lockName string
	owner    string
	ttl      uint32
}

// #region LockDetail

/**
 * The name of the lock.
 */
func (this_ *LockDetail) GetLockName() string {
	return this_.lockName
}

/**
 * The name of the lock.
 */
func (this_ *LockDetail) SetLockName(lockName string) {
	this_.lockName = lockName
}

/**
 * The owner of the lock. Only valid when user getLocks or receive LockEvent with RTM_LOCK_EVENT_TYPE_SNAPSHOT
 */
func (this_ *LockDetail) GetOwner() string {
	return this_.owner
}

/**
 * The owner of the lock. Only valid when user getLocks or receive LockEvent with RTM_LOCK_EVENT_TYPE_SNAPSHOT
 */
func (this_ *LockDetail) SetOwner(owner string) {
	this_.owner = owner
}

/**
 * The ttl of the lock.
 */
func (this_ *LockDetail) GetTtl() uint32 {
	return this_.ttl
}

/**
 * The ttl of the lock.
 */
func (this_ *LockDetail) SetTtl(ttl uint32) {
	this_.ttl = ttl
}

func NewLockDetail() *LockDetail {
	return &LockDetail{
		lockName: "",
		owner:    "",
		ttl:      0,
	}
}

// 内部方法：转换为C对象

// #endregion LockDetail

/**
*  The states of user.
 */
type UserState struct {
	UserId      string
	States      []StateItem
	StatesCount uint
}

// #region UserState

/**
 * The user id.
 */
func (this_ *UserState) GetUserId() string {
	return this_.UserId
}

/**
 * The user id.
 */
func (this_ *UserState) SetUserId(userId string) {
	this_.UserId = userId
}

/**
 * The user states.
 */
func (this_ *UserState) GetStates() []StateItem {
	return this_.States
}

/**
 * The user states.
 */
func (this_ *UserState) SetStates(states []StateItem) {
	this_.States = states
	this_.StatesCount = uint(len(states))
}

/**
 * The count of user states.
 */
func (this_ *UserState) GetStatesCount() uint {
	return this_.StatesCount
}

/**
 * The count of user states.
 */
func (this_ *UserState) SetStatesCount(statesCount uint) {
	this_.StatesCount = statesCount
}

func NewUserState() *UserState {
	userState := &UserState{
		UserId:      "",
		States:      make([]StateItem, 0),
		StatesCount: 0,
	}

	return userState
}

// #endregion UserState

/**
 *  The subscribe option.
 */
type SubscribeOptions struct {
	WithMessage  bool
	WithMetadata bool
	WithPresence bool
	WithLock     bool
	BeQuiet      bool
}

// #region SubscribeOptions

/**
 * Whether to subscribe channel with message
 */
func (this_ *SubscribeOptions) GetWithMessage() bool {
	return this_.WithMessage
}

/**
 * Whether to subscribe channel with message
 */
func (this_ *SubscribeOptions) SetWithMessage(withMessage bool) {
	this_.WithMessage = withMessage
}

/**
 * Whether to subscribe channel with metadata
 */
func (this_ *SubscribeOptions) GetWithMetadata() bool {
	return this_.WithMetadata
}

/**
 * Whether to subscribe channel with metadata
 */
func (this_ *SubscribeOptions) SetWithMetadata(withMetadata bool) {
	this_.WithMetadata = withMetadata
}

/**
 * Whether to subscribe channel with user presence
 */
func (this_ *SubscribeOptions) GetWithPresence() bool {
	return this_.WithPresence
}

/**
 * Whether to subscribe channel with user presence
 */
func (this_ *SubscribeOptions) SetWithPresence(withPresence bool) {
	this_.WithPresence = withPresence
}

/**
 * Whether to subscribe channel with quiet
 */
func (this_ *SubscribeOptions) SetWithQuiet(withQuiet bool) {
	this_.BeQuiet = withQuiet
}

/**
 * Whether to subscribe channel with quiet
 */
func (this_ *SubscribeOptions) GetWithQuiet() bool {
	return this_.BeQuiet
}

/**
 * Whether to subscribe channel with lock
 */
func (this_ *SubscribeOptions) GetWithLock() bool {
	return this_.WithLock
}

/**
 * Whether to subscribe channel with lock
 */
func (this_ *SubscribeOptions) SetWithLock(withLock bool) {
	this_.WithLock = withLock
}

func NewSubscribeOptions() *SubscribeOptions {
	return &SubscribeOptions{
		WithMessage:  true,
		WithMetadata: false,
		WithPresence: true,
		WithLock:     false,
		BeQuiet:      false,
	}
}

func (this_ *SubscribeOptions) toC() unsafe.Pointer {
	if this_ == nil {
		return nil
	}
	cOpt := C.C_SubscribeOptions_New()
	if cOpt == nil {
		return nil
	}
	cOpt.withMessage = C.bool(this_.WithMessage)
	cOpt.withMetadata = C.bool(this_.WithMetadata)
	cOpt.withPresence = C.bool(this_.WithPresence)
	cOpt.withLock = C.bool(this_.WithLock)
	cOpt.beQuiet = C.bool(this_.BeQuiet)
	return unsafe.Pointer(cOpt)
}

func freeSubscribeOptions(cOpt unsafe.Pointer) {
	if cOpt != nil {
		C.C_SubscribeOptions_Delete((*C.struct_C_SubscribeOptions)(cOpt))
	}
}

// #endregion SubscribeOptions

/**
 *  The channel information.
 */
type ChannelInfo struct {
	ChannelName string
	ChannelType RTM_CHANNEL_TYPE
}

// #region ChannelInfo

/**
 * The channel which the message was published
 */
func (this_ *ChannelInfo) GetChannelName() string {
	return this_.ChannelName
}

/**
 * The channel which the message was published
 */
func (this_ *ChannelInfo) SetChannelName(channelName string) {
	this_.ChannelName = channelName
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *ChannelInfo) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/**
 * Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE
 */
func (this_ *ChannelInfo) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

func NewChannelInfo() *ChannelInfo {
	channelInfo := &ChannelInfo{
		ChannelName: "",
		ChannelType: 0,
	}

	return channelInfo
}

// #endregion ChannelInfo

/**
 *  The option to query user presence.
 */
type PresenceOptions struct {
	IncludeUserId bool
	IncludeState  bool
	Page          string
}

// #region PresenceOptions

/**
 * Whether to display user id in query result
 */
func (this_ *PresenceOptions) GetIncludeUserId() bool {
	return this_.IncludeUserId
}

/**
 * Whether to display user id in query result
 */
func (this_ *PresenceOptions) SetIncludeUserId(includeUserId bool) {
	this_.IncludeUserId = includeUserId
}

/**
 * Whether to display user state in query result
 */
func (this_ *PresenceOptions) GetIncludeState() bool {
	return this_.IncludeState
}

/**
 * Whether to display user state in query result
 */
func (this_ *PresenceOptions) SetIncludeState(includeState bool) {
	this_.IncludeState = includeState
}

/**
 * The paging object used for pagination.
 */
func (this_ *PresenceOptions) GetPage() string {
	return this_.Page
}

/**
 * The paging object used for pagination.
 */
func (this_ *PresenceOptions) SetPage(page string) {
	this_.Page = page
}

func NewPresenceOptions() *PresenceOptions {
	presenceOptions := &PresenceOptions{
		IncludeUserId: true,
		IncludeState:  false,
		Page:          "",
	}

	return presenceOptions
}

// #endregion PresenceOptions

/**
*  The option to query user presence.
 */
type GetOnlineUsersOptions struct {
	IncludeUserId bool
	IncludeState  bool
	Page          string
}

// #region GetOnlineUsersOptions

/**
 * Whether to display user id in query result
 */
func (this_ *GetOnlineUsersOptions) GetIncludeUserId() bool {
	return this_.IncludeUserId
}

/**
 * Whether to display user id in query result
 */
func (this_ *GetOnlineUsersOptions) SetIncludeUserId(includeUserId bool) {
	this_.IncludeUserId = includeUserId
}

/**
 * Whether to display user state in query result
 */
func (this_ *GetOnlineUsersOptions) GetIncludeState() bool {
	return this_.IncludeState
}

/**
 * Whether to display user state in query result
 */
func (this_ *GetOnlineUsersOptions) SetIncludeState(includeState bool) {
	this_.IncludeState = includeState
}

/**
 * The paging object used for pagination.
 */
func (this_ *GetOnlineUsersOptions) GetPage() string {
	return this_.Page
}

/**
 * The paging object used for pagination.
 */
func (this_ *GetOnlineUsersOptions) SetPage(page string) {
	this_.Page = page
}

func NewGetOnlineUsersOptions() *GetOnlineUsersOptions {
	getOnlineUsersOptions := &GetOnlineUsersOptions{
		IncludeUserId: false,
		IncludeState:  false,
		Page:          "",
	}

	return getOnlineUsersOptions
}

// #endregion GetOnlineUsersOptions

/*
*

	@brief Publish message option
*/
type PublishOptions struct {
	ChannelType    RTM_CHANNEL_TYPE
	MessageType    RTM_MESSAGE_TYPE
	CustomType     string
	StoreInHistory bool
}

// #region PublishOptions

/*
The channel type.
*/
func (this_ *PublishOptions) GetChannelType() RTM_CHANNEL_TYPE {
	return this_.ChannelType
}

/*
The channel type.
*/
func (this_ *PublishOptions) SetChannelType(channelType RTM_CHANNEL_TYPE) {
	this_.ChannelType = channelType
}

/*
*

	The message type.
*/
func (this_ *PublishOptions) GetMessageType() RTM_MESSAGE_TYPE {
	return this_.MessageType
}

/*
*

	The message type.
*/
func (this_ *PublishOptions) SetMessageType(messageType RTM_MESSAGE_TYPE) {
	this_.MessageType = messageType
}

/*
*

	The custom type of the message, up to 32 bytes for customize
*/
func (this_ *PublishOptions) GetCustomType() string {
	return this_.CustomType
}

/*
*

	The custom type of the message, up to 32 bytes for customize
*/
func (this_ *PublishOptions) SetCustomType(customType string) {
	this_.CustomType = customType
}

/*
*

	Whether to store in history, true to enable
*/
func (this_ *PublishOptions) GetStoreInHistory() bool {
	return this_.StoreInHistory
}

/*
*

	Whether to store in history, true to enable
*/
func (this_ *PublishOptions) SetStoreInHistory(storeInHistory bool) {
	this_.StoreInHistory = storeInHistory
}

func NewPublishOptions() *PublishOptions {
	publishOptions := &PublishOptions{
		ChannelType:    0,
		MessageType:    0,
		CustomType:     "",
		StoreInHistory: false,
	}

	return publishOptions
}

func (this_ *PublishOptions) toC() unsafe.Pointer {
	if this_ == nil {
		return nil
	}
	
	cOpt := C.C_PublishOptions_New()
	if cOpt == nil {
		return nil
	}
	
	cOpt.channelType = C.enum_C_RTM_CHANNEL_TYPE(this_.ChannelType)
	cOpt.messageType = C.enum_C_RTM_MESSAGE_TYPE(this_.MessageType)
	
	if this_.CustomType != "" {
		cOpt.customType = C.CString(this_.CustomType)
	} else {
		cOpt.customType = nil
	}
	
	cOpt.storeInHistory = C.bool(this_.StoreInHistory)
	
	return unsafe.Pointer(cOpt)
}

func freePublishOptions(cOpt unsafe.Pointer) {
	if cOpt != nil {
		cPublishOpt := (*C.struct_C_PublishOptions)(cOpt)
		if cPublishOpt.customType != nil {
			C.free(unsafe.Pointer(cPublishOpt.customType))
		}
		
		C.C_PublishOptions_Delete((*C.struct_C_PublishOptions)(cOpt))
	}
}

// #endregion PublishOptions

/*
@brief topic message option
*/
type TopicMessageOptions struct {
	MessageType RTM_MESSAGE_TYPE
	SendTs      uint64
	CustomType  string
}

// #region TopicMessageOptions

/*
The message type.
*/
func (this_ *TopicMessageOptions) GetMessageType() RTM_MESSAGE_TYPE {
	return this_.MessageType
}

/*
The message type.
*/
func (this_ *TopicMessageOptions) SetMessageType(messageType RTM_MESSAGE_TYPE) {
	this_.MessageType = messageType
}

/*
The time to calibrate data with media,
only valid when user join topic with syncWithMedia in stream channel
*/
func (this_ *TopicMessageOptions) GetSendTs() uint64 {
	return this_.SendTs
}

/*
The time to calibrate data with media,
only valid when user join topic with syncWithMedia in stream channel
*/
func (this_ *TopicMessageOptions) SetSendTs(sendTs uint64) {
	this_.SendTs = sendTs
}

/*
The custom type of the message, up to 32 bytes for customize
*/
func (this_ *TopicMessageOptions) GetCustomType() string {
	return this_.CustomType
}

/*
*

	The custom type of the message, up to 32 bytes for customize
*/
func (this_ *TopicMessageOptions) SetCustomType(customType string) {
	this_.CustomType = customType
}

func NewTopicMessageOptions() *TopicMessageOptions {
	topicMessageOptions := &TopicMessageOptions{
		MessageType: 0,
		SendTs:      0,
		CustomType:  "",
	}

	return topicMessageOptions
}

// #endregion TopicMessageOptions

/*
*
@brief Proxy configuration
*/
type RtmProxyConfig struct {
	ProxyType RTM_PROXY_TYPE
	Server    string
	Port      uint16
	Account   string
	Password  string
}

// #region RtmProxyConfig

/*
*

	The Proxy type.
*/
func (this_ *RtmProxyConfig) GetProxyType() RTM_PROXY_TYPE {
	return this_.ProxyType
}

/*
*

	The Proxy type.
*/
func (this_ *RtmProxyConfig) SetProxyType(proxyType RTM_PROXY_TYPE) {
	this_.ProxyType = proxyType
}

/*
*

	The Proxy server address.
*/
func (this_ *RtmProxyConfig) GetServer() string {
	return this_.Server
}

/*
*

	The Proxy server address.
*/
func (this_ *RtmProxyConfig) SetServer(server string) {
	this_.Server = server
}

/*
*

	The Proxy server port.
*/
func (this_ *RtmProxyConfig) GetPort() uint16 {
	return this_.Port
}

/*
*

	The Proxy server port.
*/
func (this_ *RtmProxyConfig) SetPort(port uint16) {
	this_.Port = port
}

/*
*

	The Proxy user account.
*/
func (this_ *RtmProxyConfig) GetAccount() string {
	return this_.Account
}

/*
*

	The Proxy user account.
*/
func (this_ *RtmProxyConfig) SetAccount(account string) {
	this_.Account = account
}

/*
*

	The Proxy password.
*/
func (this_ *RtmProxyConfig) GetPassword() string {
	return this_.Password
}

/*
*

	The Proxy password.
*/
func (this_ *RtmProxyConfig) SetPassword(password string) {
	this_.Password = password
}

func NewRtmProxyConfig() *RtmProxyConfig {
	rtmProxyConfig := &RtmProxyConfig{
		ProxyType: 0,
		Server:    "",
		Port:      0,
		Account:   "",
		Password:  "",
	}

	return rtmProxyConfig
}

// #endregion RtmProxyConfig

/*
*
@brief encryption configuration
*/
type RtmEncryptionConfig struct {
	EncryptionMode RTM_ENCRYPTION_MODE
	EncryptionKey  string
	EncryptionSalt [32]uint8
}

// #region RtmEncryptionConfig

/**
 * The encryption mode.
 */
func (this_ *RtmEncryptionConfig) GetEncryptionMode() RTM_ENCRYPTION_MODE {
	return this_.EncryptionMode
}

/**
 * The encryption mode.
 */
func (this_ *RtmEncryptionConfig) SetEncryptionMode(encryptionMode RTM_ENCRYPTION_MODE) {
	this_.EncryptionMode = encryptionMode
}

/**
 * The encryption key in the string format.
 */
func (this_ *RtmEncryptionConfig) GetEncryptionKey() string {
	return this_.EncryptionKey
}

/**
 * The encryption key in the string format.
 */
func (this_ *RtmEncryptionConfig) SetEncryptionKey(encryptionKey string) {
	this_.EncryptionKey = encryptionKey
}

/**
 * The encryption salt.
 */
func (this_ *RtmEncryptionConfig) GetEncryptionSalt() [32]uint8 {
	return this_.EncryptionSalt
}

/**
 * The encryption salt.
 */
func (this_ *RtmEncryptionConfig) SetEncryptionSalt(encryptionSalt [32]uint8) {
	this_.EncryptionSalt = encryptionSalt
}

func NewRtmEncryptionConfig() *RtmEncryptionConfig {
	rtmEncryptionConfig := &RtmEncryptionConfig{
		EncryptionMode: 0,
		EncryptionKey:  "",
		EncryptionSalt: [32]uint8{},
	}

	return rtmEncryptionConfig
}

// #endregion RtmEncryptionConfig

// #region RtmPrivateConfig
type RtmPrivateConfig struct {
	ServiceType           uint32
	AccessPointHosts      []string
	AccessPointHostsCount uint
}

func (this_ *RtmPrivateConfig) GetServiceType() uint32 {
	return this_.ServiceType
}

func (this_ *RtmPrivateConfig) SetServiceType(serviceType uint32) {
	this_.ServiceType = serviceType
}

func (this_ *RtmPrivateConfig) GetAccessPointHosts() []string {
	return this_.AccessPointHosts
}

func (this_ *RtmPrivateConfig) SetAccessPointHosts(accessPointHosts []string) {
	this_.AccessPointHosts = accessPointHosts
	this_.AccessPointHostsCount = uint(len(accessPointHosts))
}

func (this_ *RtmPrivateConfig) GetAccessPointHostsCount() uint {
	return this_.AccessPointHostsCount
}

func (this_ *RtmPrivateConfig) SetAccessPointHostsCount(accessPointHostsCount uint) {
	this_.AccessPointHostsCount = accessPointHostsCount
}

func NewRtmPrivateConfig() *RtmPrivateConfig {
	rtmPrivateConfig := &RtmPrivateConfig{
		ServiceType:           0,
		AccessPointHosts:      make([]string, 0),
		AccessPointHostsCount: 0,
	}
	return rtmPrivateConfig
}

// #endregion RtmPrivateConfig

// link state event

type RTM_SERVICE_TYPE C.enum_C_RTM_SERVICE_TYPE

const (
	RTM_SERVICE_TYPE_RTM     RTM_SERVICE_TYPE = 0
	RTM_SERVICE_TYPE_MESSAGE RTM_SERVICE_TYPE = 1
	RTM_SERVICE_TYPE_STREAM  RTM_SERVICE_TYPE = 2
)

type LinkStateEvent struct {
	CurrentState           RTM_LINK_STATE
	PreviousState          RTM_LINK_STATE
	ServiceType            RTM_SERVICE_TYPE
	Operation              RTM_LINK_OPERATION
	ReasonCode             RTM_LINK_STATE_CHANGE_REASON
	Reason                 string
	AffectedChannels       []string
	AffectedChannelCount   uint
	UnrestoredChannels     []string
	UnrestoredChannelCount uint
	IsResumed              bool
	Timestamp              uint64
}

// #region LinkStateEvent

/**
 * Current state
 */
func (this_ *LinkStateEvent) GetCurrentState() RTM_LINK_STATE {
	return this_.CurrentState
}

/**
 * Current state
 */
func (this_ *LinkStateEvent) SetCurrentState(currentState RTM_LINK_STATE) {
	this_.CurrentState = currentState
}

/**
 * Previous state
 */
func (this_ *LinkStateEvent) GetPreviousState() RTM_LINK_STATE {
	return this_.PreviousState
}

/**
 * Previous state
 */
func (this_ *LinkStateEvent) SetPreviousState(previousState RTM_LINK_STATE) {
	this_.PreviousState = previousState
}

/**
 * Service type
 */
func (this_ *LinkStateEvent) GetServiceType() RTM_SERVICE_TYPE {
	return this_.ServiceType
}

/**
 * Service type
 */
func (this_ *LinkStateEvent) SetServiceType(serviceType RTM_SERVICE_TYPE) {
	this_.ServiceType = serviceType
}

/**
 * Operation
 */
func (this_ *LinkStateEvent) GetOperation() RTM_LINK_OPERATION {
	return this_.Operation
}

/**
 * Operation
 */
func (this_ *LinkStateEvent) SetOperation(operation RTM_LINK_OPERATION) {
	this_.Operation = operation
}

/**
 * Reason code
 */
func (this_ *LinkStateEvent) GetReasonCode() RTM_LINK_STATE_CHANGE_REASON {
	return this_.ReasonCode
}

/**
 * Reason code
 */
func (this_ *LinkStateEvent) SetReasonCode(reasonCode RTM_LINK_STATE_CHANGE_REASON) {
	this_.ReasonCode = reasonCode
}

/**
 * Reason
 */
func (this_ *LinkStateEvent) GetReason() string {
	return this_.Reason
}

/**
 * Reason
 */
func (this_ *LinkStateEvent) SetReason(reason string) {
	this_.Reason = reason
}

/**
 * Affected channels
 */
func (this_ *LinkStateEvent) GetAffectedChannels() []string {
	return this_.AffectedChannels
}

/**
 * Affected channels
 */
func (this_ *LinkStateEvent) SetAffectedChannels(affectedChannels []string) {
	this_.AffectedChannels = affectedChannels
	this_.AffectedChannelCount = uint(len(affectedChannels))
}

/**
 * Affected channel count
 */
func (this_ *LinkStateEvent) GetAffectedChannelCount() uint {
	return this_.AffectedChannelCount
}

/**
 * Affected channel count
 */
func (this_ *LinkStateEvent) SetAffectedChannelCount(affectedChannelCount uint) {
	this_.AffectedChannelCount = affectedChannelCount
}

/**
 * Unrestored channels
 */
func (this_ *LinkStateEvent) GetUnrestoredChannels() []string {
	return this_.UnrestoredChannels
}

/**
 * Unrestored channels
 */
func (this_ *LinkStateEvent) SetUnrestoredChannels(unrestoredChannels []string) {
	this_.UnrestoredChannels = unrestoredChannels
	this_.UnrestoredChannelCount = uint(len(unrestoredChannels))
}

/**
 * Unrestored channel count
 */
func (this_ *LinkStateEvent) GetUnrestoredChannelCount() uint {
	return this_.UnrestoredChannelCount
}

/**
 * Unrestored channel count
 */
func (this_ *LinkStateEvent) SetUnrestoredChannelCount(unrestoredChannelCount uint) {
	this_.UnrestoredChannelCount = unrestoredChannelCount
}

/**
 * Is resumed
 */
func (this_ *LinkStateEvent) GetIsResumed() bool {
	return this_.IsResumed
}

/**
 * Is resumed
 */
func (this_ *LinkStateEvent) SetIsResumed(isResumed bool) {
	this_.IsResumed = isResumed
}

/**
 * Timestamp
 */
func (this_ *LinkStateEvent) GetTimestamp() uint64 {
	return this_.Timestamp
}

/**
 * Timestamp
 */
func (this_ *LinkStateEvent) SetTimestamp(timestamp uint64) {
	this_.Timestamp = timestamp
}

func NewLinkStateEvent() *LinkStateEvent {
	linkStateEvent := &LinkStateEvent{
		CurrentState:           RTM_LINK_STATE_IDLE,
		PreviousState:          RTM_LINK_STATE_IDLE,
		ServiceType:            RTM_SERVICE_TYPE_RTM,
		Operation:              RTM_LINK_OPERATION_LOGIN,
		ReasonCode:             RTM_LINK_STATE_CHANGE_REASON_UNKNOWN,
		Reason:                 "",
		AffectedChannels:       make([]string, 0),
		AffectedChannelCount:   0,
		UnrestoredChannels:     make([]string, 0),
		UnrestoredChannelCount: 0,
		IsResumed:              false,
		Timestamp:              0,
	}

	return linkStateEvent
}

// #region LinkStateEvent

type HistoryMessage struct {
	MessageType   RTM_MESSAGE_TYPE
	Publisher     string
	Message       string
	MessageLength uint
	CustomType    string
	Timestamp     uint64
}

// #region HistoryMessage

/**
 * Message type
 */
func (this_ *HistoryMessage) GetMessageType() RTM_MESSAGE_TYPE {
	return this_.MessageType
}

/**
 * Message type
 */
func (this_ *HistoryMessage) SetMessageType(messageType RTM_MESSAGE_TYPE) {
	this_.MessageType = messageType
}

/**
 * The publisher
 */
func (this_ *HistoryMessage) GetPublisher() string {
	return this_.Publisher
}

/**
 * The publisher
 */
func (this_ *HistoryMessage) SetPublisher(publisher string) {
	this_.Publisher = publisher
}

/**
 * The payload
 */
func (this_ *HistoryMessage) GetMessage() string {
	return this_.Message
}

/**
 * The payload
 */
func (this_ *HistoryMessage) SetMessage(message string) {
	this_.Message = message
	this_.MessageLength = uint(len(message))
}

/**
 * The payload length
 */
func (this_ *HistoryMessage) GetMessageLength() uint {
	return this_.MessageLength
}

/**
 * The payload length
 */
func (this_ *HistoryMessage) SetMessageLength(messageLength uint) {
	this_.MessageLength = messageLength
}

/**
 * The custom type of the message
 */
func (this_ *HistoryMessage) GetCustomType() string {
	return this_.CustomType
}

/**
 * The custom type of the message
 */
func (this_ *HistoryMessage) SetCustomType(customType string) {
	this_.CustomType = customType
}

/**
 * Timestamp of the message received by rtm server
 */
func (this_ *HistoryMessage) GetTimestamp() uint64 {
	return this_.Timestamp
}

/**
 * Timestamp of the message received by rtm server
 */
func (this_ *HistoryMessage) SetTimestamp(timestamp uint64) {
	this_.Timestamp = timestamp
}

func NewHistoryMessage() *HistoryMessage {
	historyMessage := &HistoryMessage{
		MessageType:   0,
		Publisher:     "",
		Message:       "",
		MessageLength: 0,
		CustomType:    "",
		Timestamp:     0,
	}

	return historyMessage
}

// #endregion LinkStateEvent

// #endregion RtmEncryptionConfig

// #endregion agora::rtm

// #endregion agora
