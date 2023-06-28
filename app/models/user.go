package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	User_id       string             `json:"user_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Avatar        *string            `json:"avatar"`
	Token         *string            `json:"token"`
	Created_at    time.Time          `json:"created_at"`
	Last_login_at time.Time          `json:"last_login_at"`
	Logout_at     time.Time          `json:"logout_at"`
	Deleted_at    time.Time          `json:"deleted_at"`
}

type LoginResponse struct {
	Email *string `json:"email"`
	Token *string `json:"token"`
}

type GetResponse struct {
	First_name    *string   `json:"first_name"`
	Last_name     *string   `json:"last_name"`
	Password      *string   `json:"password"`
	Email         *string   `json:"email"`
	Phone         *string   `json:"phone"`
	Avatar        *string   `json:"avatar"`
	Created_at    time.Time `json:"created_at"`
	Last_login_at time.Time `json:"last_login_at"`
	Logout_at     time.Time `json:"logout_at"`
}
