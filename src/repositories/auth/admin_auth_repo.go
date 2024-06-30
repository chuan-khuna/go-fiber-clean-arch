package auth_repositories

import (
	"fiber-app/config"
	auth_entities "fiber-app/src/entities/auth"
	auth_usecases "fiber-app/src/usecases/auth"
	jwt_utils "fiber-app/src/utils/jwt"
	password_utils "fiber-app/src/utils/password"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

type GormAdminAuthRepository struct {
	db *gorm.DB
}

func NewGormAdminAuthRepository(db *gorm.DB) auth_usecases.AdminAuthRepository {
	return &GormAdminAuthRepository{db: db}
}

var secretKey = []byte(config.SecretKey)

func (r *GormAdminAuthRepository) GetByEmail(email string) (auth_entities.AdminAccount, error) {
	var account auth_entities.AdminAccount

	if err := r.db.Where("email = ?", email).First(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

func (r *GormAdminAuthRepository) IssueRefreshToken(account auth_entities.AdminAccount) (string, error) {
	tokenString, err := jwt_utils.CreateToken(jwt_utils.JwtPayload{
		ID:          account.ID,
		Email:       account.Email,
		AccountType: "admin",
	}, "refresh")

	if err != nil {
		return "", err
	}

	token, _ := jwt_utils.VerifyToken(tokenString)

	hashedTokenString := password_utils.HashSHA256(tokenString)

	// add data to db
	r.db.Create(&auth_entities.AdminRefreshToken{
		AdminAccountID: account.ID,
		Token:          hashedTokenString,
		Iat:            int64(token.Claims.(jwt.MapClaims)["iat"].(float64)),
		Exp:            int64(token.Claims.(jwt.MapClaims)["exp"].(float64)),
		TokenCreatedAt: token.Claims.(jwt.MapClaims)["createdAt"].(string),
		TokenExpAt:     token.Claims.(jwt.MapClaims)["expAt"].(string),
	})

	// invalidate old refresh token, except the current one
	r.db.Model(&auth_entities.AdminRefreshToken{}).Where("admin_account_id = ? AND token != ?", account.ID, hashedTokenString).Update("is_valid", false)

	// delete tokens to clean up the table
	r.db.Where("exp < ?", time.Now().Unix()).Delete(&auth_entities.AdminRefreshToken{})

	return tokenString, err
}

func (r *GormAdminAuthRepository) IssueAccessToken(account auth_entities.AdminAccount) (string, error) {
	tokenString, err := jwt_utils.CreateToken(jwt_utils.JwtPayload{
		ID:          account.ID,
		Email:       account.Email,
		AccountType: "admin",
	}, "refresh")

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (r *GormAdminAuthRepository) VerifyToken(token string) error {
	_, err := jwt_utils.VerifyToken(token)

	return err
}

func (r *GormAdminAuthRepository) GetAccountFromToken(token string) (auth_entities.AdminAccount, error) {
	claims, err := jwt_utils.VerifyToken(token)

	if err != nil {
		return auth_entities.AdminAccount{}, err
	}

	var account auth_entities.AdminAccount

	if err := r.db.Where("id = ?", claims.Claims.(jwt.MapClaims)["id"]).First(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

func (r *GormAdminAuthRepository) RotateRefreshToken(token string) (string, error) {

	oldToken, err := jwt_utils.VerifyToken(token)
	hashedOldToken := password_utils.HashSHA256(token)

	if err != nil {
		return "", err
	}

	if oldToken.Claims.(jwt.MapClaims)["tokenType"] != "refresh" {
		return "", fmt.Errorf("Invalid token type")
	}

	// check if this token is valid
	var refreshTokenHistory auth_entities.AdminRefreshToken
	r.db.Find(&refreshTokenHistory, "token", hashedOldToken)

	accountID := refreshTokenHistory.AdminAccountID

	if refreshTokenHistory.ID == 0 {
		return "", fmt.Errorf("Invalid refresh token - might be invalidated since token is reused")
	}

	// check if refresh token is reused
	if !refreshTokenHistory.IsValid {
		// delete all refresh token history for this user
		// equivalent to invalidate all refresh token
		// r.db.Model(&auth_entities.AdminRefreshToken{}).Where("admin_account_id", accountID).Update("is_valid", false)
		r.db.Delete(&auth_entities.AdminRefreshToken{}, "admin_account_id", accountID)
		return "", fmt.Errorf("Token is already used")
	}

	newTokenString, err := jwt_utils.RotateToken(token)
	newToken, _ := jwt_utils.VerifyToken(newTokenString)
	hashedNewTokenString := password_utils.HashSHA256(newTokenString)

	if err != nil {
		return "", err
	}

	// add data to db
	r.db.Create(&auth_entities.AdminRefreshToken{
		AdminAccountID: accountID,
		Token:          hashedNewTokenString,
		Iat:            int64(newToken.Claims.(jwt.MapClaims)["iat"].(float64)),
		Exp:            int64(newToken.Claims.(jwt.MapClaims)["exp"].(float64)),
		TokenCreatedAt: newToken.Claims.(jwt.MapClaims)["createdAt"].(string),
		TokenExpAt:     newToken.Claims.(jwt.MapClaims)["expAt"].(string),
	})

	// invalidate old refresh token, except the current one
	r.db.Model(&auth_entities.AdminRefreshToken{}).Where("admin_account_id = ? AND token != ?", accountID, hashedNewTokenString).Update("is_valid", false)

	return newTokenString, nil
}

func (r *GormAdminAuthRepository) InvalidateRefreshToken(token string) error {
	var refreshToken auth_entities.AdminRefreshToken
	r.db.Find(&refreshToken, "token", token).Update("is_valid", false)
	return nil
}

func (r *GormAdminAuthRepository) GetClaimsFromToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt_utils.VerifyToken(tokenString)

	return token.Claims, err
}
