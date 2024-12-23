package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,"`
	FirstName string `json:"first_name" validate:"required, min=2, max=100"`
	LastName string `json:"last_name " validate:"required, min=2, max=100"`
	Password string `json:"password" validate:"required, min=6, max=100"`
	Email string `json:"email" validate:"email, required"`
	Phone string `json:"phone" validate:"required"`
	Token string `json:"token"`
	Refresh_token string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User_id int `json:"user_id"`
}