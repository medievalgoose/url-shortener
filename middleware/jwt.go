package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	_ "github.com/joho/godotenv"
)

func CreateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["aud"] = username

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "include your token in the authorization header using bearer schema"})
			ctx.Abort()
			return
		}

		authHeaderSplit := strings.Split(authHeader, " ")
		providedToken := authHeaderSplit[1]

		token, err := jwt.Parse(providedToken, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return "", fmt.Errorf("error when parsing")
			}

			secretKeyByte := []byte(os.Getenv("SECRET_KEY"))
			return secretKeyByte, nil
		})

		if err != nil {
			log.Fatal(err)
		}

		if token.Valid {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}
