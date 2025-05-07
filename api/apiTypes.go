package api

import "net/http"

type server struct {
	Server *http.Server
}

type MyWriter struct {
	Data string
}
