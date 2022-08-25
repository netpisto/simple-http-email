package controllers

import (
	"context"
	"firestore/database"
	"firestore/helpers"
	"firestore/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Validator = validator.New()
)

// sign in the user with email and password and give him a token to access with later
func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse json (bad/syntax)",
		})
		return
	}
	user.Name = "1234"
	if err := Validator.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot validate data (bad/data)",
		})
		return
	}
	var foundedUser models.User
	if err := database.UsersColl.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password is not valid (authentication/error)",
		})
		return
	}
	user.Password = helpers.HashString(user.Password)
	if foundedUser.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password is not valid (authentication/error)",
		})
		return
	}
	token, err := helpers.GenerateToken(&foundedUser, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "there is an error contact the team (server/error)",
		})
		return
	}
	c.JSON(200, gin.H{
		"token":       token.Token,
		"destroyedAt": token.DestroyedAt,
		"user_id":     foundedUser.ID.Hex(),
	})
}

// sign up the user with email and password and give him a token to access with later
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot parse json (bad/syntax)",
		})
		return
	}
	if err := Validator.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot validate data (bad/data)",
		})
		return
	}
	var foundedUser models.User
	if err := database.UsersColl.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundedUser); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "there is a user with this email (email/exsist)",
		})
		return
	}
	user.Password = helpers.HashString(user.Password)
	user.CreatedAt = time.Now()
	res, err := database.UsersColl.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "there is an error contact the team (server/error)",
		})
		return
	}
	user.ID =  res.InsertedID.(primitive.ObjectID)
	token, err := helpers.GenerateToken(&user, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "there is an error contact the team (server/error)",
		})
		return
	}
	c.JSON(200, gin.H{
		"token":       token.Token,
		"destroyedAt": token.DestroyedAt,
		"user_id":     res.InsertedID,
	})
}
