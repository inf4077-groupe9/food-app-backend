package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// type User struct {
// 	UserName string `bson:"username,omitempty" json:"username" validate:"required,min=6,max=20"`
// 	Email    string `bson:"email,omitempty" json:"email" validate:"required,email"`
// 	Password string `bson:"password,omitempty" json:"password" validate:"required,min=8"`
// }

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	UserName string             `json:"username" validate:"required,min=6,max=20"`
	Email    string             `json:"email" validate:"required,email"`
}

type UserRegister struct {
	UserName             string `json:"username" validate:"required,min=6,max=20"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password-confirmation" validate:"required,min=8"`
}

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	UserName string             `json:"username" validate:"required,min=6,max=20"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required,min=8"`
	Token    interface{}        `json:"token" validate:"jwt"`
	IsActive bool               `json:"is-active" bson:"is-active" validate:"boolean"`
}

func ParseToUserResponse(u User, ur *UserResponse) {
	ur.Id = u.Id
	ur.UserName = u.UserName
	ur.Email = u.Email
}

func ParseUserRegisterToUser(u UserRegister, ur *User) {
	ur.UserName = u.UserName
	ur.Email = u.Email
	ur.Password = u.Password
}
