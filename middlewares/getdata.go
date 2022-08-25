package middlewares

import (
	"errors"
	"firestore/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAccessToken(c *gin.Context) (models.AccessToken, error) {
	data := make(map[string]string)

	data["createdAt"] = c.GetString("createdAt")
	data["destroyedAt"] = c.GetString("destroyedAt")
	data["user_id"] = c.GetString("user_id")
	data["token"] = c.GetString("token")
	data["create_ip"] = c.GetString("create_ip")
	data["token_id"] = c.GetString("token_id")
	if data["user_id"] == "" {
		// user is no not logged in
		println("access data is null")
		return models.AccessToken{}, errors.New("user is not logged in")
	}
	var token models.AccessToken
	createdate, _ := time.Parse(data["createdAt"], time.RFC3339)
	destroyeddate, _ := time.Parse(data["destroyedAt"], time.RFC3339)
	if time.Now().Before(destroyeddate) {
		return models.AccessToken{}, errors.New("this is a destroyed token")
	}
	token_id, _ := primitive.ObjectIDFromHex(data["token_id"])
	token = models.AccessToken{
		Token:       data["token"],
		UserId:      data["user_id"],
		CreatedAt:   createdate,
		DestroyedAt: destroyeddate,

		CreateIp: data["create_ip"],
		ID:       token_id,
	}

	return token, nil
}

func A2Str(val any) string {
	return fmt.Sprintf("%v", val)
}
