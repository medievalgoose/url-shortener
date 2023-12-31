package main

import (
	"log"
	"medievalgoose/url-shortener/db"
	"medievalgoose/url-shortener/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.POST("/login/", loginHandler)
	router.POST("/register/", registerHandler)
	router.POST("/create/", middleware.ValidateToken(), createHandler)
	router.GET("/:url/", redirectHandler)

	router.Run("localhost:8080")
}

func loginHandler(ctx *gin.Context) {
	var providedUser db.User

	err := ctx.BindJSON(&providedUser)
	if err != nil {
		log.Fatal(err)
	}

	validUser := db.ValidateUser(providedUser.Username, providedUser.Password)

	if validUser {
		token := middleware.CreateToken(providedUser.Username)
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed login"})
	}
}

func registerHandler(ctx *gin.Context) {
	var newUser db.User

	err := ctx.BindJSON(&newUser)
	if err != nil {
		log.Fatal(err)
	}

	ok := db.CreateUser(newUser.Username, newUser.Password)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user creation failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "your account has been created"})
}

func createHandler(ctx *gin.Context) {
	var UrlRequest db.CustomURL

	err := ctx.BindJSON(&UrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	generatedShortUrl, err := db.AddNewUrl(UrlRequest)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"short url": generatedShortUrl})
}

func redirectHandler(ctx *gin.Context) {
	// TODO: Add redirect logic here.
	requestedUrl := ctx.Param("url")

	redirectUrl, err := db.GetPlainUrl(requestedUrl)
	if err != nil {
		log.Fatal(err)
	}

	ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}
