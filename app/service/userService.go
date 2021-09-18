package service

import (
	"github.com/AchmadAlli/go_auth_api/app"
	"github.com/AchmadAlli/go_auth_api/app/model"
	"github.com/AchmadAlli/go_auth_api/app/request"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func CreateUserService(app *app.App) *UserService {
	return &UserService{app.DB.Model(&model.User{})}
}

func (s *UserService) Index() (*[]model.User, error) {
	users := []model.User{}
	err := s.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *UserService) Store(data *request.StoreUser) (*model.User, error) {
	user := model.User{
		Username: data.UserName,
		Fullname: data.Fullname,
		Password: data.Password,
	}

	err := s.db.Model(&model.User{}).Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Update(data *request.UpdateUser, id uint) (*model.User, error) {

	user, err := s.Show(id)

	if err != nil {
		return nil, err
	}

	userData := handleUpdate(user, data.Password, data.Fullname)
	err = s.db.Model(&user).Update(userData).Error

	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func (s *UserService) UpdateAvatar(path string, user *model.User) (*model.User, error) {
	err := s.db.Model(&user).Update("avatar", path).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Show(id uint) (*model.User, error) {
	user := model.User{}
	err := s.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Destroy(id uint) (*model.User, error) {
	user := model.User{}
	err := s.db.Delete(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func handleUpdate(user *model.User, password, fullname string) model.User {
	if password != "" {
		user.Password = password
	}

	if fullname != "" {
		user.Fullname = fullname
	}

	return model.User{
		Password: user.Password,
		Fullname: user.Fullname,
	}
}
