package main

import "net/http"

func isConnected(next http.HandlerFunc, init http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if clientConnected {
			next.ServeHTTP(w, r)
		} else {
			init.ServeHTTP(w, r)
			next.ServeHTTP(w, r)
		}
	}
}
