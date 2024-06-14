package config

import "time"

type TokenExpirationConfig struct {
	Hour   time.Duration
	Minute time.Duration
	Second time.Duration
}

var AccessTokenExpiration = TokenExpirationConfig{
	Hour:   2 * time.Hour,
	Minute: 0 * time.Minute,
	Second: 0 * time.Second,
}

var RefreshTokenExpiration = TokenExpirationConfig{
	Hour:   24 * time.Hour,
	Minute: 0 * time.Minute,
	Second: 0 * time.Second,
}

var ResetPasswordTokenExpiration = TokenExpirationConfig{
	Hour:   0 * time.Hour,
	Minute: 10 * time.Minute,
	Second: 0 * time.Second,
}
