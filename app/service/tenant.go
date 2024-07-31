package service

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/repository"
)

type (
	TenantService interface {
		GetRefreshToken(key string, secret string) (map[string]interface{}, error)
	}

	tenantService struct {
		TenantRepository repository.TenantRepository
	}
)

// GetRefreshToken gets an access token
func (s tenantService) GetRefreshToken(key string, secret string) (map[string]interface{}, error) {
	tenant, err := s.TenantRepository.FindByKeys(key, secret)
	if err != nil {
		return nil, &appErr.AuthError{
			Message: "Invalid API key or secret",
			Code:    http.StatusBadRequest,
		}
	}

	now := time.Now().Unix()
	expiresIn := int64(7 * 24 * 3600)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims) // refer to https://datatracker.ietf.org/doc/html/rfc7519#section-4.1 for more details
	claims["iss"] = "localhost"
	claims["aud"] = "tenant"
	claims["sub"] = tenant.ID
	claims["iat"] = now
	claims["nbf"] = now
	claims["exp"] = now + expiresIn // 7 days
	accessToken, err := token.SignedString([]byte(tenant.AppSecret))

	return map[string]interface{}{
		"access_token": accessToken,
		"expires_in":   expiresIn,
		"scope":        tenant.Scope,
	}, nil
}

// NewTenantService creates a new TenantService
func NewTenantService(r repository.TenantRepository) TenantService {
	return &tenantService{
		TenantRepository: r,
	}
}
