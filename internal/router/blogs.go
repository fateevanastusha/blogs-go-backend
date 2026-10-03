package router

import (
	"context"
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type blogsService interface {
	Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error)
	Update(ctx context.Context, ID int, name, description, websiteURL string) (model.Blog, error)
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (model.Blog, error)
	GetAll(ctx context.Context) []model.Blog
}

type Blogs struct {
	logger  *logrus.Logger
	service blogsService
}

func ConfigureBlogsRouter(r *mux.Router, logger *logrus.Logger, service blogsService) error {
	blogs := &Blogs{
		logger:  logger,
		service: service,
	}
	sub := r.PathPrefix("/blogs").Subrouter()
	sub.HandleFunc("", blogs.getAll()).Methods("GET")
	sub.HandleFunc("/{ID}", blogs.getByID()).Methods("GET")
	sub.HandleFunc("", blogs.create()).Methods("POST")
	sub.HandleFunc("/{ID}", blogs.putByID()).Methods("PUT")
	sub.HandleFunc("/{ID}", blogs.deleteByID()).Methods("DELETE")
	return nil
}

func (b *Blogs) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Blogs) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Blogs) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Blogs) putByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Blogs) deleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
