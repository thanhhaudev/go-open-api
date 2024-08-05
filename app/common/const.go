package common

const (
	AuthRefreshTokenExpire int64 = 7 * 24 * 3600
	AuthRefreshTokenPrefix       = "refresh_token"
	AuthAccessTokenPrefix        = "access_token"
	AuthAccessTokenExpire  int64 = 30 * 60
)
