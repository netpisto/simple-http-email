package controllers

import (
	"context"
	"firestore/database"
	"firestore/middlewares"
	"firestore/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SendMessageToUser(c *gin.Context) {

	var email models.Email
	err := c.BindJSON(&email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse this json"})
		return
	}
	token, err := middlewares.GetAccessToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Validator.Struct(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "we cannot validate your email"})
		return
	}

	toUser, err := database.GetUserFromEmail(email.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "we cannot validate your email"})
		return
	}
	user, err := database.GetUserFromId(token.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "we cannot validate your email"})
		return
	}
	if user.Email == toUser.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot send an email to your self"})
		return
	}
	emaildb := models.EmailDb{
		To:      toUser.ID.Hex(),
		From:    token.UserId,
		Content: email.Content,
		CreatedAt: time.Now(),
		
	}
	res, err := database.EmailColl.InsertOne(context.TODO(), emaildb)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "we cannot register your email (contact the team)"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"email_id": res.InsertedID, "from": user.Email, "to": toUser.Email})
}
