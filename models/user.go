package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
	Name string `json:"name" validate:"min=4"`
	CreatedAt time.Time `json:"createdAt"`
}