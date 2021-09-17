package service

import (
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

func (s UserService) Index() (*[]model.User, error) {
	users := []model.User{}
	err := s.db.Model(&model.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s UserService) Store() (*model.User, error) {
	return &model.User{}, nil
}

func (s UserService) Update() (*model.User, error) {
	return &model.User{}, nil
}

func (s UserService) Show(id uint) (*model.User, error) {
	user := model.User{}
	err := s.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Destroy(id uint) (*model.User, error) {
	user := model.User{}
	err := s.db.Delete(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
