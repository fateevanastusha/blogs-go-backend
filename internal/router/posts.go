package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Posts struct {
	logger *logrus.Logger
}

func ConfigurePostsRouter(r *mux.Router, logger *logrus.Logger) *Posts {
	blogs := &Posts{
		logger: logger,
	}
	sub := r.PathPrefix("/blogs").Subrouter()
	sub.HandleFunc("", blogs.getAll()).Methods("GET")
	sub.HandleFunc("/{id}", blogs.getById()).Methods("GET")
	sub.HandleFunc("", blogs.create()).Methods("POST")
	sub.HandleFunc("/{id}", blogs.putById()).Methods("PUT")
	sub.HandleFunc("/{id}", blogs.deleteById()).Methods("DELETE")
	return blogs
}

func (b *Posts) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) putById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (b *Posts) deleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
