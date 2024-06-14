package auth_entities

import "gorm.io/gorm"

type AdminAccount struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	IsRevoked bool   `json:"is_revoked" gorm:"default:false"`
	gorm.Model
}

type AdminRefreshToken struct {
	ID             uint `gorm:"primaryKey"`
	AdminAccountID uint `json:"admin_account_id"`
	AdminAccount   AdminAccount
	Token          string
	Iat            int64
	Exp            int64
	TokenCreatedAt string
	TokenExpAt     string
	IsValid        bool `gorm:"default:true"`
}
