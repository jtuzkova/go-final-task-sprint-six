package server

import (
	"log"
	"net/http"
	"time"

	h "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateRouter(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", h.MainHandle)
	router.HandleFunc("/upload", h.UploadHandle)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: httpServer,
	}
}