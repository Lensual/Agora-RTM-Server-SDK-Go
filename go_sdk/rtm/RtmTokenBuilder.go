package agorartm

import (
	"fmt"
	"os"
	"time"

	rtmtokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
)

// TokenConfig Token配置结构体
type TokenConfig struct {
	AppId               string
	AppCertificate      string
	UserId              string
	ExpirationInSeconds uint32
}

// NewTokenConfig 创建新的Token配置
func NewTokenConfig(appId, appCertificate, userId string) *TokenConfig {
	return &TokenConfig{
		AppId:               appId,
		AppCertificate:      appCertificate,
		UserId:              userId,
		ExpirationInSeconds: 3600, // 默认1小时过期
	}
}

// SetExpiration 设置过期时间（秒）
func (tc *TokenConfig) SetExpiration(seconds uint32) {
	tc.ExpirationInSeconds = seconds
}

// BuildRtmToken 动态生成RTM Token
func BuildRtmToken(config *TokenConfig) (string, error) {
	if config == nil {
		return "", fmt.Errorf("token config is nil")
	}

	if config.AppId == "" || config.AppCertificate == "" || config.UserId == "" {
		return "", fmt.Errorf("appId, appCertificate and userId are required")
	}

	token, err := rtmtokenbuilder.BuildToken(
		config.AppId,
		config.AppCertificate,
		config.UserId,
		config.ExpirationInSeconds,
	)

	if err != nil {
		return "", fmt.Errorf("failed to build RTM token: %w", err)
	}

	return token, nil
}

// BuildRtmTokenWithDefaults 使用默认配置生成RTM Token
func BuildRtmTokenWithDefaults(appId, appCertificate, userId string) (string, error) {
	config := NewTokenConfig(appId, appCertificate, userId)
	return BuildRtmToken(config)
}

// BuildRtmTokenFromEnv 从环境变量生成RTM Token
func BuildRtmTokenFromEnv(userId string) (string, error) {
	appId := os.Getenv("AGORA_APP_ID")
	appCertificate := os.Getenv("AGORA_APP_CERTIFICATE")

	if appId == "" {
		return "", fmt.Errorf("environment variable AGORA_APP_ID is not set")
	}

	if appCertificate == "" {
		return "", fmt.Errorf("environment variable AGORA_APP_CERTIFICATE is not set")
	}

	return BuildRtmTokenWithDefaults(appId, appCertificate, userId)
}

// BuildRtmTokenWithParams 使用提供的参数构建RTM Token
func BuildRtmTokenWithParams(appId, appCertificate, userId string) (string, error) {
	if appId == "" {
		return "", fmt.Errorf("appId cannot be empty")
	}

	if appCertificate == "" {
		return "", fmt.Errorf("appCertificate cannot be empty")
	}

	return BuildRtmTokenWithDefaults(appId, appCertificate, userId)
}

// BuildRtmTokenWithCustomExpiration 生成指定过期时间的RTM Token
func BuildRtmTokenWithCustomExpiration(appId, appCertificate, userId string, expirationInSeconds uint32) (string, error) {
	config := NewTokenConfig(appId, appCertificate, userId)
	config.SetExpiration(expirationInSeconds)
	return BuildRtmToken(config)
}

// BuildRtmTokenWithDuration 生成指定持续时间的RTM Token
func BuildRtmTokenWithDuration(appId, appCertificate, userId string, duration time.Duration) (string, error) {
	expirationInSeconds := uint32(duration.Seconds())
	return BuildRtmTokenWithCustomExpiration(appId, appCertificate, userId, expirationInSeconds)
}

// ValidateTokenConfig 验证Token配置
func ValidateTokenConfig(config *TokenConfig) error {
	if config == nil {
		return fmt.Errorf("token config is nil")
	}

	if config.AppId == "" {
		return fmt.Errorf("appId is required")
	}

	if config.AppCertificate == "" {
		return fmt.Errorf("appCertificate is required")
	}

	if config.UserId == "" {
		return fmt.Errorf("userId is required")
	}

	if config.ExpirationInSeconds == 0 {
		return fmt.Errorf("expiration time must be greater than 0")
	}

	return nil
}

// GetTokenInfo 获取Token信息（用于调试）
func GetTokenInfo(config *TokenConfig) string {
	if config == nil {
		return "Token config is nil"
	}

	return fmt.Sprintf("TokenConfig{AppId: %s, UserId: %s, Expiration: %d seconds}",
		maskString(config.AppId), config.UserId, config.ExpirationInSeconds)
}

// maskString 遮盖敏感信息（显示前4位和后4位）
func maskString(s string) string {
	if len(s) <= 8 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}
