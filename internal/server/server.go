package server

import (
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/router"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	Router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func Start(address string) error {
	server := &Server{
		Router: mux.NewRouter(),
		logger: logrus.New(),
	}
	server.configureRouter()
	return http.ListenAndServe(address, server)
}

func (s *Server) configureRouter() {
	router.ConfigureBlogsRouter(s.Router, s.logger)
}
