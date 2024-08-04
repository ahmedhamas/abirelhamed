package server

import (
	"abir-el-hamd/internal/handlers"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/api", handlers.HomeApi)
	mux.HandleFunc("POST /cases/add", handlers.AddCase)
	mux.HandleFunc("DELETE /case/{id}", handlers.DeleteCase)
	mux.HandleFunc("GET /case/{id}", handlers.GetCase)
	mux.HandleFunc("/api/case/{id}", handlers.CaseApi)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	return http.ListenAndServe(s.listenAddr, mux)
}
