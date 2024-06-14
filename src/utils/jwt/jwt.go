package jwt

import (
	"fiber-app/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// https://permify.co/post/jwt-authentication-go/

type JwtPayload struct {
	ID          uint
	Email       string
	AccountType string
}

func CreateToken(payload JwtPayload, tokenType string) (string, error) {
	secretKey := []byte(config.SecretKey)

	var expirationConf config.TokenExpirationConfig

	if tokenType == "refresh" {
		expirationConf = config.RefreshTokenExpiration
	} else if tokenType == "access" {
		expirationConf = config.AccessTokenExpiration
	} else {
		return "", fmt.Errorf("Invalid token type")
	}

	createdTime := time.Now()

	expTime := createdTime.Add(expirationConf.Hour).Add(expirationConf.Minute).Add(expirationConf.Second)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          payload.ID,
		"email":       payload.Email,
		"accountType": payload.AccountType,
		"iat":         createdTime.Unix(),
		"createdAt":   createdTime.Format("2006-01-02 15:04:05"),
		"exp":         expTime.Unix(),
		"expAt":       expTime.Format("2006-01-02 15:04:05"),
		"tokenType":   tokenType,
	})

	tokenString, err := claims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RotateToken(tokenString string) (string, error) {
	// Regenerate token with old token payload, but new created time

	secretKey := []byte(config.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	createdTime := time.Now()
	expTime := time.Unix(int64(claims["exp"].(float64)), 0)

	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          claims["id"],
		"email":       claims["email"],
		"accountType": claims["accountType"],
		"iat":         createdTime.Unix(),
		"createdAt":   createdTime.Format("2006-01-02 15:04:05"),
		"exp":         expTime.Unix(),
		"expAt":       expTime.Format("2006-01-02 15:04:05"),
		"tokenType":   claims["tokenType"],
	})

	newTokenString, err := newClaims.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// verify token signature and expiration

	secretKey := []byte(config.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	expUnixTime := token.Claims.(jwt.MapClaims)["exp"].(float64)
	if time.Now().Unix() > int64(expUnixTime) {
		return nil, fmt.Errorf("Token is expired")
	}

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func ExtractPayload(token *jwt.Token) JwtPayload {
	// extract payload from token claims

	claims := token.Claims.(jwt.MapClaims)

	return JwtPayload{
		ID:          uint(claims["id"].(float64)),
		Email:       claims["email"].(string),
		AccountType: claims["accountType"].(string),
	}
}
