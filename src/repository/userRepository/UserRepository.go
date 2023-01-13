package userrepository

import (
	"backend.go/src/database"
	"backend.go/src/models/user"
)

func GetAll() ([]user.User, error) {
	var users []user.User
	var err error

	if err := database.Connection().Find(&users).Error; err != nil {
		return nil, err
	}

	return users, err
}

func GetOne(id string) (user.User, error) {
	var user user.User
	var err error

	if err := database.Connection().First(&user, id).Error; err != nil {
		return user, err
	}

	return user, err
}

func Create(user user.User) (user.User, error) {
	var err error

	if err := database.Connection().Create(&user).Error; err != nil {
		return user, err
	}

	return user, err
}

func Update(user user.User) (user.User, error) {
	var err error

	if err := database.Connection().Save(&user).Error; err != nil {
		return user, err
	}

	return user, err
}

func Patch(id string, user user.User) (user.User, error) {
	var err error

	if err := database.Connection().Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return user, err
	}

	return user, err
}

func Delete(user user.User) (user.User, error) {
	var err error

	if err := database.Connection().Delete(&user).Error; err != nil {
		return user, err
	}

	return user, err
}

func GetByEmail(email string) (user.User, error) {
	var user user.User
	var err error

	if err := database.Connection().Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, err
}

func GetByUsername(username string) (user.User, error) {
	var user user.User
	var err error

	if err := database.Connection().Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, err
}
