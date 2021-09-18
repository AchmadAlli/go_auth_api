package service

import (
	"errors"
	"time"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/model"
	"github.com/AchmadAlli/go_auth_api/app/request"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
	db *gorm.DB
}

type UserAuth struct {
	User  *model.User `json:"user,omitempty"`
	Token string      `json:"token"`
}

func CreateAuthService(app *app.App) *AuthService {
	return &AuthService{app.DB.Model(&model.User{})}
}

func (s *AuthService) Login(auth *request.AuthUser) (*UserAuth, error) {
	user := &model.User{}
	err := s.db.Where("username = ?", auth.Username).Find(&user).Error
	if err != nil {
		return nil, err
	}

	isCorrect := helper.CheckPasswordHash(auth.Password, user.Password)
	if !isCorrect {
		return nil, errors.New("invalid password")
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	userAuth := UserAuth{
		User:  user,
		Token: token,
	}
	return &userAuth, nil
}

func generateToken(user *model.User) (string, error) {
	expire := time.Now().Add(24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"expire":  expire,
	})
	return token.SignedString([]byte("rahasia"))
}
