package service

import "famesensor/go-graphql-jwt/models"

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := models.FindUserByEmail(&user, email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserById(id string) (*models.User, error) {
	var user models.User
	err := models.FindUserById(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	return models.CreateUser(user).Error
}
