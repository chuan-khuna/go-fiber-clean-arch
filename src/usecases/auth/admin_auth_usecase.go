package auth_usecases

import (
	auth_entities "fiber-app/src/entities/auth"
	"fiber-app/src/utils/password"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type AdminAuthUseCase interface {
	Login(payload auth_entities.LoginPayload) (auth_entities.AdminAccount, auth_entities.TokenPair, error)
	RefreshAccessToken(refreshToken string) (auth_entities.TokenPair, error)
	VerifyToken(token string) error
	Logout(tokenPair auth_entities.TokenPair) error
}

type AdminAuthService struct {
	repo AdminAuthRepository
}

// ------
// implement the methods of the interface below
// ------

func NewAdminAuthService(repo AdminAuthRepository) AdminAuthUseCase {
	return &AdminAuthService{repo: repo}
}

func (s *AdminAuthService) Login(payload auth_entities.LoginPayload) (auth_entities.AdminAccount, auth_entities.TokenPair, error) {
	account, err := s.repo.GetByEmail(payload.Email)

	if err != nil {
		return auth_entities.AdminAccount{}, auth_entities.TokenPair{}, fmt.Errorf("Invalid credentials")
	}

	isAuthenticated, err := password.VerifyPassword(payload.Password, account.Password)
	if err != nil {
		return account, auth_entities.TokenPair{}, err
	}

	if !isAuthenticated {
		return account, auth_entities.TokenPair{}, fmt.Errorf("Invalid email or password")
	}

	if account.IsRevoked {
		return account, auth_entities.TokenPair{}, fmt.Errorf("Account is revoked")
	}

	// issue token
	accessToken, _ := s.repo.IssueAccessToken(account)
	refreshToken, _ := s.repo.IssueRefreshToken(account)

	tokenPair := auth_entities.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return account, tokenPair, nil
}

func (s *AdminAuthService) RefreshAccessToken(refreshToken string) (auth_entities.TokenPair, error) {
	claims, err := s.repo.GetClaimsFromToken(refreshToken)
	if err != nil {
		return auth_entities.TokenPair{}, err
	}

	if claims.(jwt.MapClaims)["tokenType"] != "refresh" {
		return auth_entities.TokenPair{}, fmt.Errorf("Invalid token type")
	}

	newToken, err := s.repo.RotateRefreshToken(refreshToken)

	if err != nil {
		return auth_entities.TokenPair{}, err
	}

	account, err := s.repo.GetAccountFromToken(newToken)
	if err != nil {
		return auth_entities.TokenPair{}, err
	}

	accessToken, err := s.repo.IssueAccessToken(account)
	if err != nil {
		return auth_entities.TokenPair{}, err
	}

	return auth_entities.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newToken,
	}, nil
}

func (s *AdminAuthService) VerifyToken(token string) error {
	err := s.repo.VerifyToken(token)
	return err
}

func (s *AdminAuthService) Logout(tokenPair auth_entities.TokenPair) error {
	s.repo.InvalidateRefreshToken(tokenPair.RefreshToken)
	return nil
}
