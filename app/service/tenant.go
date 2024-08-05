package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/thanhhaudev/openapi-go/app/common"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/model"
	"github.com/thanhhaudev/openapi-go/app/repository"
)

type (
	TenantService interface {
		GetRefreshToken(ctx context.Context, key string, secret string) (map[string]interface{}, error)
		GetAccessToken(ctx context.Context, refreshToken string) (map[string]interface{}, error)
	}

	tenantService struct {
		TenantRepository repository.TenantRepository
		RedisClient      *redis.Client
		logger           *logrus.Logger
	}
)

// GetAccessToken gets an access token
func (s *tenantService) GetAccessToken(ctx context.Context, refreshToken string) (map[string]interface{}, error) {
	s.logger.WithFields(logrus.Fields{
		"refreshToken": refreshToken,
	}).Info("GetAccessToken called")

	if len(refreshToken) == 0 {
		return nil, &appErr.AuthError{
			Message: "Invalid refresh token",
			Code:    http.StatusBadRequest,
		}
	}

	// Retrieve the API key from Redis
	apiKey, err := s.RedisClient.Get(ctx, fmt.Sprintf("%s.%s", common.AuthRefreshTokenPrefix, refreshToken)).Result()
	if err != nil {
		s.logger.WithError(err).Error("Failed to get API key from Redis")

		return nil, &appErr.AuthError{
			Message: "Internal server error",
			Code:    http.StatusInternalServerError,
		}
	}

	tenant, err := s.TenantRepository.FindByApiKey(apiKey)
	if err != nil {
		s.logger.WithError(err).Error("Failed to find tenant by API key")

		return nil, &appErr.AuthError{
			Message: "Internal server error",
			Code:    http.StatusInternalServerError,
		}
	}

	// verify the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(tenant.ApiSecret), nil
	},
		jwt.WithAudience("tenant"),
		jwt.WithIssuer("localhost"),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		s.logger.WithError(err).Error("Failed to parse refresh token")

		return nil, &appErr.AuthError{
			Message: "Internal server error",
			Code:    http.StatusBadRequest,
		}
	}

	if !token.Valid {
		s.logger.Error("Invalid refresh token")

		return nil, &appErr.AuthError{
			Message: "Invalid refresh token",
			Code:    http.StatusBadRequest,
		}
	}

	expiresIn := common.AuthAccessTokenExpire // 2 days
	accessToken, err := buildToken(tenant, expiresIn)
	if err != nil {
		s.logger.WithError(err).Error("Failed to build access token")

		return nil, err
	}

	// Save the access token to Redis
	s.RedisClient.Set(ctx, fmt.Sprintf("%s.%s", common.AuthAccessTokenPrefix, accessToken), tenant.ApiKey, time.Duration(expiresIn)*time.Second)

	return map[string]interface{}{
		"access_token": accessToken,
		"expires_in":   expiresIn,
		"scope":        tenant.Scope,
	}, nil
}

// GetRefreshToken gets an access token
func (s *tenantService) GetRefreshToken(ctx context.Context, key string, secret string) (map[string]interface{}, error) {
	tenant, err := s.TenantRepository.Find(key, secret)
	if err != nil {
		return nil, &appErr.AuthError{
			Message: "Invalid API key or secret",
			Code:    http.StatusBadRequest,
		}
	}

	expiresIn := common.AuthRefreshTokenExpire
	refreshToken, err := buildToken(tenant, expiresIn)
	if err != nil {
		s.logger.WithError(err).Error("Failed to build access token")

		return nil, err
	}

	// save the refresh token to Redis
	// key: refresh_token.<token string>, value: apiKey
	s.RedisClient.Set(ctx, fmt.Sprintf("%s.%s", common.AuthRefreshTokenPrefix, refreshToken), tenant.ApiKey, time.Duration(expiresIn)*time.Second)

	return map[string]interface{}{
		"access_token": refreshToken,
		"expires_in":   expiresIn,
		"scope":        tenant.Scope,
	}, nil
}

// buildToken builds a token
func buildToken(tenant *model.Tenant, e int64) (string, error) {
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims) // refer to https://datatracker.ietf.org/doc/html/rfc7519#section-4.1 for more details
	claims["iss"] = "localhost"
	claims["aud"] = "tenant"
	claims["sub"] = tenant.ID
	claims["iat"] = now
	claims["nbf"] = now
	claims["exp"] = now + e // 7 days

	refreshToken, err := token.SignedString([]byte(tenant.ApiSecret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

// NewTenantService creates a new TenantService
func NewTenantService(r repository.TenantRepository, s *redis.Client, l *logrus.Logger) TenantService {
	return &tenantService{
		TenantRepository: r,
		RedisClient:      s,
		logger:           l,
	}
}
