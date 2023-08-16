package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
	Password   = "" // you can omit it if you didn't set the pwd
)

func initRouter(database *Database) *gin.Engine {
	r := gin.Default()

	// add user
	r.POST("/points", func(ctx *gin.Context) {
		var userJson User
		if err := ctx.ShouldBindJSON(&userJson); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := database.SaveUser(&userJson)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": userJson})
	})

	// query user
	r.GET("/points/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")

		user, err := database.GetUser(username)
		if err != nil {
			if errors.Is(err, ErrNil) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "No record found for " + username})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": user})
	})

	// get leaderboard
	r.GET("/leaderboard", func(ctx *gin.Context) {
		leaderboard, err := database.GetLeaderboard()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
	})

	return r
}

func main() {
	database, err := NewDatabase(RedisAddr, Password)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	router := initRouter(database)
	log.Fatal(router.Run(ListenAddr))
}
