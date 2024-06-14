package auth_entities

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaimPayload struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	TokenType string `json:"token"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
	ExpAt     string `json:"expAt"`
	CreatedAt string `json:"createdAt"`
}
