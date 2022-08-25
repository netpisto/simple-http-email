package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Email struct {
	Content string `json:"content" validate:"required,min=5"`
	To string `json:"to" validate:"required"`
}

type EmailDb struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Content string `json:"content"`
	From string `json:"from"`
	To string `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
	Seen bool `json:"seen"`
}