package router

import (
	"context"
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type postsService interface {
	Create(ctx context.Context, title, shortDescription, content string, blogId int) (model.Post, error)
	Update(ctx context.Context, id int, title, shortDescription, content string) (model.Post, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (model.Post, error)
	GetAll(ctx context.Context) []model.Post
}

type Posts struct {
	logger  *logrus.Logger
	service postsService
}

func ConfigurePostsRouter(r *mux.Router, logger *logrus.Logger, service postsService) *Posts {
	posts := &Posts{
		logger:  logger,
		service: service,
	}
	sub := r.PathPrefix("/posts").Subrouter()
	sub.HandleFunc("", posts.getAll()).Methods("GET")
	sub.HandleFunc("/{id}", posts.getByID()).Methods("GET")
	sub.HandleFunc("", posts.create()).Methods("POST")
	sub.HandleFunc("/{id}", posts.putByID()).Methods("PUT")
	sub.HandleFunc("/{id}", posts.deleteByID()).Methods("DELETE")
	return posts
}

func (b *Posts) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (b *Posts) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) putByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) deleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
