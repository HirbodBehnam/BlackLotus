package main

import (
	"log/slog"
	"net/http"
)

func GateEndpoint(w http.ResponseWriter, r *http.Request) {
	// Request must be post
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	// Body must be parsed
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Log
	userAgent := r.Header.Get("User-Agent")
	ip := r.RemoteAddr
	slog.Debug("Gate:", ip, userAgent, r.Form)
	// Send a command
	// TODO

}
