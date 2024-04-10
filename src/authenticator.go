package main

import (
	"github.com/zmb3/spotify"
	"net/http"
)

func authenticator(w http.ResponseWriter, r *http.Request) {
	auth = spotify.NewAuthenticator("http://localhost:8080/", spotify.ScopeUserLibraryRead, spotify.ScopeUserFollowRead, spotify.ScopeStreaming, spotify.ScopeUserModifyPlaybackState)
	http.Redirect(w, r, auth.AuthURL("state-string"), http.StatusFound)

}

func initClient(w http.ResponseWriter, r *http.Request) {
	token, err = auth.Token("state-string", r)
	checkError("Error while fetching the token")
	client = auth.NewClient(token)
	clientConnected = true
}
