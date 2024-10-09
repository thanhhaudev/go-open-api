package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/thanhhaudev/go-open-api/app/common"
	appErr "github.com/thanhhaudev/go-open-api/app/error"
	"github.com/thanhhaudev/go-open-api/app/util"
)

// verifyScope middleware
func verifyScope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the scope from the request
		scopes, ok := r.Context().Value(common.AuthAccessTokenScopes).([]interface{})
		if !ok {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		if len(scopes) == 0 {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		route := mux.CurrentRoute(r)
		pathTemplate, _ := route.GetPathTemplate()
		requiredScope := detectScope(pathTemplate, r.Method)
		if requiredScope == nil {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		logger.Debugf("required scope: %s for %s: %s", *requiredScope, r.Method, r.URL.Path)

		for _, scope := range scopes {
			if scope == *requiredScope {
				next.ServeHTTP(w, r)
				return
			}
		}

		util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)
	})
}

// detectScope detects the scope of the path and method
func detectScope(p, m string) *string {
	for scope, route := range routeMap {
		for _, r := range route {
			if r.Path == p && r.Method == m {
				return &scope
			}
		}
	}

	return nil
}

// verifyToken middleware
func verifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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
		apiKey, err := redisStore.Client.Get(ctx, fmt.Sprintf("%s.%s", common.AuthAccessTokenPrefix, token)).Result()
		if err != nil {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		tenant, err := tenantRepo.FindByApiKey(apiKey)
		if err != nil {
			util.Response(w, appErr.NewUnauthorizedError(), http.StatusUnauthorized)

			return
		}

		// verify the access token
		accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(tenant.ApiSecret), nil // verify the token with the tenant's secret
		},
			jwt.WithAudience("tenant"),
			jwt.WithIssuer("localhost"),
			jwt.WithExpirationRequired(),
		)

		// valid when audience, issuer, and expiration are valid
		if !accessToken.Valid {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if !ok {
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		scopes := claims["scopes"]
		if len(scopes.([]interface{})) == 0 {
			// Why is this assertion to []interface{}?
			// When GetScopes() returns []string, but when it is assigned to jwt.MapClaims it becomes []interface{} because jwt.MapClaims is a map[string]interface{}
			util.Response(w, appErr.NewForbiddenError(), http.StatusForbidden)

			return
		}

		// add scopes to the context
		r = r.WithContext(context.WithValue(ctx, common.AuthAccessTokenScopes, scopes))

		next.ServeHTTP(w, r)
	})
}
