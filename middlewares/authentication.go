package middlewares

import (
	"context"
	"firestore/database"
	"firestore/models"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
)

// this is an authentication middleware
func Authenticate(c *gin.Context) {
	token := c.GetHeader("token")
	var tokenmodel models.AccessToken
	err := database.AuthsColl.FindOne(context.TODO(), gin.H{"token": token}).Decode(&tokenmodel)
	if err != nil {
		return
	}
	
	c.Set("createdAt",tokenmodel.CreatedAt.Format(time.RFC3339))
	c.Set("destroyedAt",tokenmodel.DestroyedAt.Format(time.RFC3339))  
	c.Set("user_id",tokenmodel.UserId)
	c.Set("token",token)
	c.Set("create_ip",tokenmodel.CreateIp)
	c.Set("token_id",tokenmodel.ID.Hex())

}
