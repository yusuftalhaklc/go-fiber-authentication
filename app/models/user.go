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
	UserRole      Role               `json:"user_role"`
	Avatar        *string            `json:"avatar"`
	Created_at    time.Time          `json:"created_at"`
	Last_login_at time.Time          `json:"last_login_at"`
	Logout_at     time.Time          `json:"logout_at"`
	Deleted_at    time.Time          `json:"deleted_at"`
}

type Role struct {
	RoleDesc string `json:"role_desc"`
	RoleId   int    `json:"role_id"`
}

var roles = []Role{
	{
		RoleDesc: "admin",
		RoleId:   4001,
	},
	{
		RoleDesc: "user",
		RoleId:   2001,
	},
}

func IsRoleValid(role Role) bool {
	for _, value := range roles {
		if value.RoleDesc == role.RoleDesc && value.RoleId == role.RoleId {
			return true
		}
	}
	return false
}
func IsAdminValid(role Role) bool {
	adminDetails := roles[0]
	if role.RoleDesc == adminDetails.RoleDesc && role.RoleId == adminDetails.RoleId {
		return true
	}
	return false
}

type LoginResponse struct {
	Email *string `json:"email"`
	Token string  `json:"token"`
}

type GetResponse struct {
	First_name    *string   `json:"first_name"`
	Last_name     *string   `json:"last_name"`
	Password      *string   `json:"password"`
	Email         *string   `json:"email"`
	Phone         *string   `json:"phone"`
	UserRole      Role      `json:"user_role"`
	Avatar        *string   `json:"avatar"`
	Created_at    time.Time `json:"created_at"`
	Last_login_at time.Time `json:"last_login_at"`
	Logout_at     time.Time `json:"logout_at"`
}
