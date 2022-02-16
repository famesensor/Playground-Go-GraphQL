package models

import (
	"famesensor/go-graphql-jwt/database"
	"famesensor/go-graphql-jwt/graph/model"

	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

func (it User) ToUserGraph() *model.User {
	return &model.User{
		ID:    it.ID,
		Email: it.Email,
		Name:  it.Name,
	}
}

func CreateUser(data *User) *gorm.DB {
	return database.GetDB().Create(data)
}

func FindUser(dest interface{}, condition ...interface{}) *gorm.DB {
	return database.GetDB().Model(&User{}).Take(dest, condition...)
}

func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}

func FindUserById(dest interface{}, id string) *gorm.DB {
	return FindUser(dest, "id = ?", id)
}
