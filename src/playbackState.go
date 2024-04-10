package main

import "net/http"

func pause(w http.ResponseWriter, r *http.Request) {
	err = client.Pause()
	checkError("Error while pausing the song")
}

func play(w http.ResponseWriter, r *http.Request) {
	err = client.Play()
	checkError("Error while playing the song")
}
