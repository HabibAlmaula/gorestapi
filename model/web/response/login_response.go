package response

import "time"

type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	ExpiresIn    time.Time    `json:"expires_in"`
	User         UserResponse `json:"user"`
}
