package controllers

import (
	"context"
	"firestore/database"
	"firestore/middlewares"
	"firestore/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// get all emails and return them to the user
func GetAllEmails(c *gin.Context) {
	token, err := middlewares.GetAccessToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "loggin first to see your emails",
		})
		return
	}
	var emails []models.EmailDb = []models.EmailDb{}
	cursor, err := database.EmailColl.Find(context.TODO(), bson.M{"$or": bson.A{
		bson.M{"to": token.UserId}, bson.M{"from": token.UserId},
	}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &emails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "we have an error gherating your emails",
		})
		return
	}

	c.JSON(http.StatusOK, emails)
}

// get exact email with his id
func GetEmail(c *gin.Context) {

}

// func MarkSeenEmail(user_id ,id string ){
// 	database.UsersColl.FindOneAndUpdate()
// }
