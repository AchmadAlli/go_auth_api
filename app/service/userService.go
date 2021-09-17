package service

import (
	"errors"

	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/model"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func CreateUserService(app *app.App) *UserService {
	return &UserService{app.DB}
}

func (u UserService) Index() (*[]model.User, error) {
	return &[]model.User{}, errors.New("this func is worked")
}

func (u UserService) Store() (*model.User, error) {
	return &model.User{}, nil
}

func (u UserService) Update() (*model.User, error) {
	return &model.User{}, nil
}

func (u UserService) Show() (*model.User, error) {
	return &model.User{}, nil
}

func (u UserService) Destroy() error {
	return nil
}