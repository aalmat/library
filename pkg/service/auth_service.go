package service

import (
	"errors"
	"github.com/aalmat/bookstore/models"
	"github.com/aalmat/bookstore/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenTime = 12 * time.Hour
const signInKey = "Ajbaq#k1ih-vaj"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{r}
}

func (a *AuthService) CreateUser(user models.User) (uint, error) {
	return a.repo.CreateUser(user)
} // id, err

func (a *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := a.repo.GetUser(email, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.UserRole,
	})

	return token.SignedString([]byte(signInKey))
}

func (a *AuthService) ParseToken(tokenString string) (uint, models.Role, error) {
	t, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign in method")
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, models.Client, err
	}

	claims, ok := t.Claims.(*MyClaims)
	if !ok {
		return 0, models.Client, errors.New("invalid token claims")
	}
	return claims.UserId, claims.UserRole, nil

}

type MyClaims struct {
	jwt.StandardClaims
	UserId   uint        `json:"user_id"`
	UserRole models.Role `json:"user_role"`
}
