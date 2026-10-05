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
	server, err := New()
	if err != nil {
		server.logger.Fatal(err)
	}
	return http.ListenAndServe(address, server)
}

func New() (*Server, error) {
	server := &Server{
		Router: mux.NewRouter(),
		logger: logrus.New(),
	}
	server.configureRouter()
	return server, nil
}

func (s *Server) configureRouter() {

	blogsRepo := storage.NewBlogsRepo()
	postsRepo := storage.NewPostsRepo()

	blogsService := service.NewBlogsService(blogsRepo, postsRepo)
	postsService := service.NewPostsService(postsRepo, blogsRepo)

	if err := router.ConfigureBlogsRouter(s.Router, s.logger, blogsService); err != nil {
		s.logger.Fatal(err)
	}
	if err := router.ConfigurePostsRouter(s.Router, s.logger, postsService); err != nil {
		s.logger.Fatal(err)
	}

}
