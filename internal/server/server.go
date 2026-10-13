package server

import (
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/config"
	"github.com/fateevanastusha/blogs-go-backend/internal/middleware"
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

func Start(cfg config.Config) error {
	server, err := New(cfg)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":"+cfg.Port, server)
}

func New(cfg config.Config) (*Server, error) {
	server := &Server{
		Router: mux.NewRouter(),
		logger: logrus.New(),
	}
	server.configureRouter(cfg)
	return server, nil
}

func (s *Server) configureRouter(cfg config.Config) {
	auth := middleware.BasicAuth(cfg.AuthLogin, cfg.AuthPassword)

	blogsRepo := storage.NewBlogsRepo()
	postsRepo := storage.NewPostsRepo()

	blogsService := service.NewBlogsService(blogsRepo, postsRepo)
	postsService := service.NewPostsService(postsRepo, blogsRepo)

	if err := router.ConfigureBlogsRouter(s.Router, s.logger, blogsService, auth); err != nil {
		s.logger.Fatal(err)
	}
	if err := router.ConfigurePostsRouter(s.Router, s.logger, postsService, auth); err != nil {
		s.logger.Fatal(err)
	}

}
