package main

import (
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"log"
)

var err error
var auth spotify.Authenticator
var token *oauth2.Token
var client spotify.Client
var clientConnected = false

func checkError(message string) {
	if err != nil {
		log.Fatalln(message, ":", err)
	}
}

func main() {
	err = godotenv.Load()
	checkError("Error while loading .env")

	handleHttp()
	serve(":8080")
}
