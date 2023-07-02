package facebook

/**
    @date: 2023/7/2
**/

// AccessTokenResp
type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// DebugTokenResp
type DebugTokenResp struct {
	Data struct {
		AppID     string `json:"app_id"`
		ExpiresAt int64  `json:"expires_at"`
		UserID    string `json:"user_id"`
	} `json:"data"`
}
