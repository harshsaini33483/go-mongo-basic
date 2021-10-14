package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mongo-basic/lib"
	"github.com/go-mongo-basic/model"
	"gopkg.in/mgo.v2/bson"
)

var db, collection = "test", "user"

func main() {

	var hostEnv = "localhost"
	if len(os.Getenv("HOST")) > 0 {
		hostEnv = os.Getenv("HOST")
	}

	lib.Init(hostEnv, db, collection)
	StartServer()
	fmt.Printf("Server Started")
}

func StartServer() {
	router := gin.Default()

	// Get Welcome Message
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Dummy App")

	})
	// Create user record
	router.POST("/users", func(c *gin.Context) {
		user := model.User{}
		name := c.Query("name")
		if len(name) <= 0 {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"status":  "failed",
					"message": "Invalid request body",
				})
			return
		}
		user.ID = bson.NewObjectId()
		user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()

		err := lib.DefaultCollection.Insert(user)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"status":  "failed",
					"message": "Error in the user insertion",
				})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{
				"status": "success",
				"user":   &user,
			})
	})

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8000")
}
