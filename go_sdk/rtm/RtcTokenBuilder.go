package agorartm

import (
	"fmt"
	"os"
	"time"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

// RtcTokenConfig RTC Token配置结构体
type RtcTokenConfig struct {
	AppId                                 string
	AppCertificate                        string
	ChannelName                           string
	Uid                                   uint32
	UidStr                                string
	TokenExpirationInSeconds              uint32
	PrivilegeExpirationInSeconds          uint32
	JoinChannelPrivilegeExpireInSeconds   uint32
	PubAudioPrivilegeExpireInSeconds      uint32
	PubVideoPrivilegeExpireInSeconds      uint32
	PubDataStreamPrivilegeExpireInSeconds uint32
	Role                                  rtctokenbuilder.Role
}

// NewRtcTokenConfig 创建新的RTC Token配置
func NewRtcTokenConfig(appId, appCertificate, channelName, uidStr string) *RtcTokenConfig {
	return &RtcTokenConfig{
		AppId:                                 appId,
		AppCertificate:                        appCertificate,
		ChannelName:                           channelName,
		UidStr:                                uidStr,
		TokenExpirationInSeconds:              3600, // 默认1小时过期
		PrivilegeExpirationInSeconds:          3600,
		JoinChannelPrivilegeExpireInSeconds:   3600,
		PubAudioPrivilegeExpireInSeconds:      3600,
		PubVideoPrivilegeExpireInSeconds:      3600,
		PubDataStreamPrivilegeExpireInSeconds: 3600,
		Role:                                  rtctokenbuilder.RolePublisher,
	}
}

// SetUid 设置数字UID
func (rtc *RtcTokenConfig) SetUid(uid uint32) {
	rtc.Uid = uid
}

// SetExpiration 设置Token过期时间（秒）
func (rtc *RtcTokenConfig) SetExpiration(seconds uint32) {
	rtc.TokenExpirationInSeconds = seconds
}

// SetPrivilegeExpiration 设置权限过期时间（秒）
func (rtc *RtcTokenConfig) SetPrivilegeExpiration(seconds uint32) {
	rtc.PrivilegeExpirationInSeconds = seconds
}

// SetJoinChannelPrivilegeExpiration 设置加入频道权限过期时间（秒）
func (rtc *RtcTokenConfig) SetJoinChannelPrivilegeExpiration(seconds uint32) {
	rtc.JoinChannelPrivilegeExpireInSeconds = seconds
}

// SetPubAudioPrivilegeExpiration 设置发布音频权限过期时间（秒）
func (rtc *RtcTokenConfig) SetPubAudioPrivilegeExpiration(seconds uint32) {
	rtc.PubAudioPrivilegeExpireInSeconds = seconds
}

// SetPubVideoPrivilegeExpiration 设置发布视频权限过期时间（秒）
func (rtc *RtcTokenConfig) SetPubVideoPrivilegeExpiration(seconds uint32) {
	rtc.PubVideoPrivilegeExpireInSeconds = seconds
}

// SetPubDataStreamPrivilegeExpiration 设置发布数据流权限过期时间（秒）
func (rtc *RtcTokenConfig) SetPubDataStreamPrivilegeExpiration(seconds uint32) {
	rtc.PubDataStreamPrivilegeExpireInSeconds = seconds
}

// SetRole 设置用户角色
func (rtc *RtcTokenConfig) SetRole(role rtctokenbuilder.Role) {
	rtc.Role = role
}

// BuildRtcTokenWithUid 使用数字UID生成RTC Token
func BuildRtcTokenWithUid(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithUid(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.Uid,
		config.Role,
		config.TokenExpirationInSeconds,
		config.PrivilegeExpirationInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTC token with UID: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithUserAccount 使用用户账户生成RTC Token
func BuildRtcTokenWithUserAccount(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithUserAccount(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.UidStr,
		config.Role,
		config.TokenExpirationInSeconds,
		config.PrivilegeExpirationInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTC token with user account: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithUidAndPrivilege 使用数字UID和详细权限生成RTC Token
func BuildRtcTokenWithUidAndPrivilege(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithUidAndPrivilege(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.Uid,
		config.TokenExpirationInSeconds,
		config.JoinChannelPrivilegeExpireInSeconds,
		config.PubAudioPrivilegeExpireInSeconds,
		config.PubVideoPrivilegeExpireInSeconds,
		config.PubDataStreamPrivilegeExpireInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTC token with UID and privilege: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithUserAccountAndPrivilege 使用用户账户和详细权限生成RTC Token
func BuildRtcTokenWithUserAccountAndPrivilege(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithUserAccountAndPrivilege(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.UidStr,
		config.TokenExpirationInSeconds,
		config.JoinChannelPrivilegeExpireInSeconds,
		config.PubAudioPrivilegeExpireInSeconds,
		config.PubVideoPrivilegeExpireInSeconds,
		config.PubDataStreamPrivilegeExpireInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTC token with user account and privilege: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithRtm 生成RTM Token
func BuildRtcTokenWithRtm(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithRtm(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.UidStr,
		config.Role,
		config.TokenExpirationInSeconds,
		config.PrivilegeExpirationInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTM token: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithRtm2 生成RTM2 Token
func BuildRtcTokenWithRtm2(config *RtcTokenConfig) (string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return "", err
	}

	token, err := rtctokenbuilder.BuildTokenWithRtm2(
		config.AppId,
		config.AppCertificate,
		config.ChannelName,
		config.UidStr,
		config.Role,
		config.TokenExpirationInSeconds,
		config.PubAudioPrivilegeExpireInSeconds,
		config.PubVideoPrivilegeExpireInSeconds,
		config.PubDataStreamPrivilegeExpireInSeconds,
		config.TokenExpirationInSeconds, // streamExpirationInSeconds
		config.UidStr,                   // streamUid
		config.TokenExpirationInSeconds, // 最后一个streamExpirationInSeconds参数
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTM2 token: %w", err)
	}

	return token, nil
}

// BuildRtcTokenWithDefaults 使用默认配置生成RTC Token（用户账户方式）
func BuildRtcTokenWithDefaults(appId, appCertificate, channelName, uidStr string) (string, error) {
	config := NewRtcTokenConfig(appId, appCertificate, channelName, uidStr)
	return BuildRtcTokenWithUserAccount(config)
}

// BuildRtcTokenFromEnv 从环境变量生成RTC Token
func BuildRtcTokenFromEnv(channelName, uidStr string) (string, error) {
	appId := os.Getenv("AGORA_APP_ID")
	appCertificate := os.Getenv("AGORA_APP_CERTIFICATE")

	if appId == "" {
		return "", fmt.Errorf("environment variable AGORA_APP_ID is not set")
	}

	if appCertificate == "" {
		return "", fmt.Errorf("environment variable AGORA_APP_CERTIFICATE is not set")
	}

	return BuildRtcTokenWithDefaults(appId, appCertificate, channelName, uidStr)
}

// BuildRtcTokenWithParams 使用提供的参数构建RTC Token
func BuildRtcTokenWithParams(appId, appCertificate, channelName, uidStr string) (string, error) {
	if appId == "" {
		return "", fmt.Errorf("appId cannot be empty")
	}

	if appCertificate == "" {
		return "", fmt.Errorf("appCertificate cannot be empty")
	}

	if channelName == "" {
		return "", fmt.Errorf("channelName cannot be empty")
	}

	if uidStr == "" {
		return "", fmt.Errorf("uidStr cannot be empty")
	}

	return BuildRtcTokenWithDefaults(appId, appCertificate, channelName, uidStr)
}

// BuildRtcTokenWithCustomExpiration 生成指定过期时间的RTC Token
func BuildRtcTokenWithCustomExpiration(appId, appCertificate, channelName, uidStr string, expirationInSeconds uint32) (string, error) {
	config := NewRtcTokenConfig(appId, appCertificate, channelName, uidStr)
	config.SetExpiration(expirationInSeconds)
	return BuildRtcTokenWithUserAccount(config)
}

// BuildRtcTokenWithDuration 生成指定持续时间的RTC Token
func BuildRtcTokenWithDuration(appId, appCertificate, channelName, uidStr string, duration time.Duration) (string, error) {
	expirationInSeconds := uint32(duration.Seconds())
	return BuildRtcTokenWithCustomExpiration(appId, appCertificate, channelName, uidStr, expirationInSeconds)
}

// BuildRtcTokenWithRole 生成指定角色的RTC Token
func BuildRtcTokenWithRole(appId, appCertificate, channelName, uidStr string, role rtctokenbuilder.Role) (string, error) {
	config := NewRtcTokenConfig(appId, appCertificate, channelName, uidStr)
	config.SetRole(role)
	return BuildRtcTokenWithUserAccount(config)
}

// validateRtcTokenConfig 验证RTC Token配置
func validateRtcTokenConfig(config *RtcTokenConfig) error {
	if config == nil {
		return fmt.Errorf("RTC token config is nil")
	}

	if config.AppId == "" {
		return fmt.Errorf("appId is required")
	}

	if config.AppCertificate == "" {
		return fmt.Errorf("appCertificate is required")
	}

	if config.ChannelName == "" {
		return fmt.Errorf("channelName is required")
	}

	if config.UidStr == "" {
		return fmt.Errorf("uidStr is required")
	}

	if config.TokenExpirationInSeconds == 0 {
		return fmt.Errorf("token expiration time must be greater than 0")
	}

	return nil
}

// GetRtcTokenInfo 获取RTC Token信息（用于调试）
func GetRtcTokenInfo(config *RtcTokenConfig) string {
	if config == nil {
		return "RTC token config is nil"
	}

	return fmt.Sprintf("RtcTokenConfig{AppId: %s, ChannelName: %s, UidStr: %s, Role: %d, TokenExpiration: %d seconds}",
		maskString(config.AppId), config.ChannelName, config.UidStr, config.Role, config.TokenExpirationInSeconds)
}

// BuildAllRtcTokenTypes 生成所有类型的RTC Token（用于测试）
func BuildAllRtcTokenTypes(config *RtcTokenConfig) (map[string]string, error) {
	if err := validateRtcTokenConfig(config); err != nil {
		return nil, err
	}

	tokens := make(map[string]string)
	var err error

	// 生成各种类型的Token
	if config.Uid > 0 {
		tokens["uid"], err = BuildRtcTokenWithUid(config)
		if err != nil {
			return nil, fmt.Errorf("failed to build token with UID: %w", err)
		}
	}

	tokens["userAccount"], err = BuildRtcTokenWithUserAccount(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build token with user account: %w", err)
	}

	if config.Uid > 0 {
		tokens["uidAndPrivilege"], err = BuildRtcTokenWithUidAndPrivilege(config)
		if err != nil {
			return nil, fmt.Errorf("failed to build token with UID and privilege: %w", err)
		}
	}

	tokens["userAccountAndPrivilege"], err = BuildRtcTokenWithUserAccountAndPrivilege(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build token with user account and privilege: %w", err)
	}

	tokens["rtm"], err = BuildRtcTokenWithRtm(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build RTM token: %w", err)
	}

	tokens["rtm2"], err = BuildRtcTokenWithRtm2(config)
	if err != nil {
		return nil, fmt.Errorf("failed to build RTM2 token: %w", err)
	}

	return tokens, nil
}
