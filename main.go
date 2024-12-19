package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
var googleOauthConfig = &oauth2.Config{}
func init(){
	err:=godotenv.Load()
	if err!=nil{
		panic("Error loading .env file")
	}
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth/google/callback", googleCallbackHandler)
	r.Run()
}

func googleCallbackHandler(c *gin.Context){
	code:=c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err!=nil{
		fmt.Println("Error exchanging code: "+ err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}