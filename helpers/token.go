package helpers

import (
	"context"
	"errors"
	"firestore/database"
	"firestore/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(user *models.User, c *gin.Context) (models.AccessToken, error) {
	tokentext := HashString(time.Now().String() + user.Email + user.ID.Hex())
	token := models.AccessToken{
		CreatedAt:   time.Now(),
		UserId:      user.ID.Hex(),
		CreateIp:    c.ClientIP(),
		DestroyedAt: time.Now().Add(time.Hour * 24),
		Token:       tokentext,
	}
	_, err := database.AuthsColl.InsertOne(context.TODO(), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "there is an error contact the team (server/error)",
		})
		return models.AccessToken{}, errors.New("we cannot generate a token")
	}
	return token, nil
}
