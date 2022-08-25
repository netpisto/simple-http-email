package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessToken struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	DestroyedAt time.Time `json:"destroyedAt"`
	UserId string `json:"userid"`
	Token string `json:"token"`
	CreateIp string `json:"createIp"`
}