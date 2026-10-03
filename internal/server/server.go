package server

import (
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/router"
	"github.com/fateevanastusha/blogs-go-backend/internal/service"
	"github.com/fateevanastusha/blogs-go-backend/internal/storage"
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
	s.configureBlogs()
	s.configurePosts()
}

func (s *Server) configureBlogs() {
	repo := storage.NewBlogsRepo()
	service := service.NewBlogsService(repo)
	router.ConfigureBlogsRouter(s.Router, s.logger, service)
}

func (s *Server) configurePosts() {
	repo := storage.NewPostsRepo()
	service := service.NewPostsService(repo)
	router.ConfigurePostsRouter(s.Router, s.logger, service)
}
