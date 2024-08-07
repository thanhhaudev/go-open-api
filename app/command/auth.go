package command

type (
	AccessTokenRequest struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	}

	ExchangeTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenRequest struct {
		AccessToken string `json:"access_token"`
	}
)
