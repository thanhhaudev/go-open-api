package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/thanhhaudev/openapi-go/app/common"
	appErr "github.com/thanhhaudev/openapi-go/app/error"
	"github.com/thanhhaudev/openapi-go/app/repository"
	"github.com/thanhhaudev/openapi-go/app/util"
)

type AuthMiddleware struct {
	RedisClient      *redis.Client
	TenantRepository repository.TenantRepository
}

func (a AuthMiddleware) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		token := strings.ReplaceAll(auth, "Bearer ", "")
		if len(token) == 0 {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		// Retrieve the API key from Redis
		apiKey, err := a.RedisClient.Get(r.Context(), fmt.Sprintf("%s.%s", common.AuthAccessTokenPrefix, token)).Result()
		if err != nil {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		tenant, err := a.TenantRepository.FindByApiKey(apiKey)
		if err != nil {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		// verify the access token
		accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(tenant.ApiSecret), nil
		},
			jwt.WithAudience("tenant"),
			jwt.WithIssuer("localhost"),
			jwt.WithExpirationRequired(),
		)

		if !accessToken.Valid {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewAuthMiddleware(r repository.TenantRepository, c *redis.Client) *AuthMiddleware {
	return &AuthMiddleware{
		TenantRepository: r,
		RedisClient:      c,
	}
}
