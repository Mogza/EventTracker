package main

import "net/http"

func handleHttp() {
	http.HandleFunc("/login", authenticator)
	http.HandleFunc("/init", initClient)
	http.HandleFunc("/pause", isConnected(pause, initClient))
	http.HandleFunc("/play", isConnected(play, initClient))
}

func serve(port string) {
	err = http.ListenAndServe(port, nil)
	checkError("Error while to listen on :8080")
}
