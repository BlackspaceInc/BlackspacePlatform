package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

type JwtCustomClaims struct {
	Id   string         `json:"id"`
	User models.UserORM `json:"user"`
	jwt.StandardClaims
}

type TokenValidationResponse struct {
	Id        string         `json:"id"`
	ExpiresAt time.Time      `json:"expires_at"`
	User      models.UserORM `json:"user"`
}
