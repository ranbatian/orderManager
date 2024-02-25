package db

import "github.com/simon/orderManager/internal/model"

func CreateUser(user model.User) error {
	return DB.Create(&user).Error
}

func GetByUsername(username string) (*model.User, error) {
	user := model.User{}
	err := DB.Model(&model.User{}).Where("name = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
