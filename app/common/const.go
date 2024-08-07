package common

const (
	AuthRefreshTokenExpire int64 = 7 * 24 * 3600 // 7 days
	AuthRefreshTokenPrefix       = "refresh_token"
	AuthAccessTokenPrefix        = "access_token"
	AuthAccessTokenExpire  int64 = 30 * 60 // 30 minutes
	AuthAccessTokenScopes        = "scopes"

	ScopeManageUser    = "MANAGE_USER"
	ScopeManageMessage = "MANAGE_MESSAGE"
)
