package service

import (
	"errors"
	"famesensor/go-graphql-jwt/constant"
	"famesensor/go-graphql-jwt/graph/model"
	"famesensor/go-graphql-jwt/models"
	"famesensor/go-graphql-jwt/utils"
	"famesensor/go-graphql-jwt/utils/bcrypt"
	"famesensor/go-graphql-jwt/utils/jwt"
	"strings"
)

func Login(email, password string) (*models.LoginResponse, error) {
	user, err := FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.ComparePassword(user.Password, password); err != nil {
		return nil, errors.New("email/password is incorrect")
	}

	token, err := jwt.JwtGenerate(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{Token: token}, nil
}

func Register(body model.RegisterUser) error {
	user, err := FindUserByEmail(body.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already existing")
	}

	userInfo := models.User{
		ID:       utils.RandomUUID(constant.PREFIX_USER_ID),
		Email:    strings.ToLower(body.Email),
		Name:     body.Name,
		Password: bcrypt.HashPassword(body.Password),
	}

	return CreateUser(&userInfo)
}
