package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Blogs struct {
	logger *logrus.Logger
}

func ConfigureBlogsRouter(r *mux.Router, logger *logrus.Logger) *Blogs {
	blogs := &Blogs{
		logger: logger,
	}
	sub := r.PathPrefix("/blogs").Subrouter()
	sub.HandleFunc("", blogs.getAll()).Methods("GET")
	sub.HandleFunc("/{id}", blogs.getByID()).Methods("GET")
	sub.HandleFunc("", blogs.create()).Methods("POST")
	sub.HandleFunc("/{id}", blogs.putByID()).Methods("PUT")
	sub.HandleFunc("/{id}", blogs.deleteByID()).Methods("DELETE")
	return blogs
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
