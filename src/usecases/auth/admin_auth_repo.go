package auth_usecases

import (
	auth_entities "fiber-app/src/entities/auth"

	"github.com/golang-jwt/jwt/v5"
)

type AdminAuthRepository interface {
	GetByEmail(email string) (auth_entities.AdminAccount, error)
	IssueRefreshToken(account auth_entities.AdminAccount) (string, error)
	IssueAccessToken(account auth_entities.AdminAccount) (string, error)

	// functions for extracting data from token
	VerifyToken(token string) error
	GetAccountFromToken(token string) (auth_entities.AdminAccount, error)
	GetClaimsFromToken(tokenString string) (jwt.Claims, error)

	// get the current token
	// issue a new token, with the same claims, but different expiration time
	RotateRefreshToken(token string) (string, error)
	InvalidateRefreshToken(token string) error
}
