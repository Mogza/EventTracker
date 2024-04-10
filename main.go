package main

import (
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var err error
var auth spotify.Authenticator
var token *oauth2.Token
var client spotify.Client

func checkError(message string) {
	if err != nil {
		log.Fatalf(message, ": %v", err)
	}
}

func main() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .end : %v", err)
	}

	http.HandleFunc("/login", authenticator)
	http.HandleFunc("/pause", pause)
	http.HandleFunc("/play", play)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error while listening on 8080 : %v", err)
	}
}

func authenticator(w http.ResponseWriter, r *http.Request) {
	auth = spotify.NewAuthenticator("http://localhost:8080/", spotify.ScopeUserLibraryRead, spotify.ScopeUserFollowRead, spotify.ScopeStreaming, spotify.ScopeUserModifyPlaybackState)
	http.Redirect(w, r, auth.AuthURL("state-string"), http.StatusFound)

}

func pause(w http.ResponseWriter, r *http.Request) {
	token, err = auth.Token("state-string", r)
	if err != nil {
		log.Fatalf("Error while fetching the token : %v", err)
	}
	client = auth.NewClient(token)
	err = client.Pause()
	if err != nil {
		log.Fatalf("Error while pausing the song : %v", err)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	token, err = auth.Token("state-string", r)
	if err != nil {
		log.Fatalf("Error while fetching the token : %v", err)
	}
	client = auth.NewClient(token)
	err = client.Play()
	if err != nil {
		log.Fatalf("Error while playing the song : %v", err)
	}
}
