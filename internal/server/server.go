package server

import (
	"log/slog"
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/config"
	"github.com/fateevanastusha/blogs-go-backend/internal/middleware"
	"github.com/fateevanastusha/blogs-go-backend/internal/router"
	"github.com/fateevanastusha/blogs-go-backend/internal/service"
	"github.com/fateevanastusha/blogs-go-backend/internal/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	logger  *slog.Logger
	Router  *mux.Router
	handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func Start(cfg config.Config, logger *slog.Logger) error {
	server, err := New(cfg, logger)
	if err != nil {
		return err
	}
	logger.Info("server started", "port", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, server)
}

func New(cfg config.Config, logger *slog.Logger) (*Server, error) {
	server := &Server{
		Router: mux.NewRouter(),
		logger: logger,
	}
	if err := server.configureRouter(cfg); err != nil {
		return nil, err
	}
	server.handler = middleware.RequestID(
		middleware.Logging(logger)(
			middleware.Recover(server.Router),
		),
	)
	return server, nil
}

func (s *Server) configureRouter(cfg config.Config) error {
	auth := middleware.BasicAuth(cfg.AuthLogin, cfg.AuthPassword)

	blogsRepo := storage.NewBlogsRepo()
	postsRepo := storage.NewPostsRepo()

	blogsService := service.NewBlogsService(blogsRepo, postsRepo)
	postsService := service.NewPostsService(postsRepo, blogsRepo)

	if err := router.ConfigureBlogsRouter(s.Router, blogsService, auth); err != nil {
		return err
	}
	return router.ConfigurePostsRouter(s.Router, postsService, auth)
}
