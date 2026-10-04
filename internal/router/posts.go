package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type postsService interface {
	Create(ctx context.Context, title, shortDescription, content string, blogID int) (model.Post, error)
	Update(ctx context.Context, ID int, title, shortDescription, content string) (model.Post, error)
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
}

type Posts struct {
	logger  *logrus.Logger
	service postsService
}

func ConfigurePostsRouter(r *mux.Router, logger *logrus.Logger, service postsService) error {
	posts := &Posts{
		logger:  logger,
		service: service,
	}
	sub := r.PathPrefix("/posts").Subrouter()
	sub.HandleFunc("", posts.getAll()).Methods("GET")
	sub.HandleFunc("/{ID}", posts.getByID()).Methods("GET")
	sub.HandleFunc("", posts.create()).Methods("POST")
	sub.HandleFunc("/{ID}", posts.putByID()).Methods("PUT")
	sub.HandleFunc("/{ID}", posts.deleteByID()).Methods("DELETE")
	return nil
}

func (b *Posts) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := b.service.GetAll(r.Context())
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusOK, res)
		return
	}
}
func (b *Posts) getByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondError(w, r, http.StatusBadRequest, errors.New("bad ID"))
			return
		}

		res, err := b.service.GetByID(r.Context(), id)
		if err != nil {
			respondError(w, r, http.StatusNotFound, err)
			return
		}
		respond(w, r, http.StatusOK, res)
		return
	}
}
func (b *Posts) create() http.HandlerFunc {
	type request struct {
		Title            string `json:"title" validate:"required,max=30"`
		ShortDescription string `json:"shortDescription" validate:"required,max=100"`
		Content          string `json:"content" validate:"required,max=1000"`
		BlogID           int    `json:"blogID" validate:"required,min=1"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		req.Title = strings.TrimSpace(req.Title)
		req.ShortDescription = strings.TrimSpace(req.ShortDescription)
		req.Content = strings.TrimSpace(req.Content)
		if err := validateStruct(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		res, err := b.service.Create(r.Context(), req.Title, req.ShortDescription, req.Content, req.BlogID)
		if err != nil {
			if errors.As(err, model.ErrBlogNotFound) {
				respondError(w, r, http.StatusNotFound, err)
				return
			}
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusCreated, res)
		return
	}
}
func (b *Posts) putByID() http.HandlerFunc {
	type request struct {
		Title            string `json:"title" validate:"required,max=30"`
		ShortDescription string `json:"shortDescription" validate:"required,max=100"`
		Content          string `json:"content" validate:"required,max=1000"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondError(w, r, http.StatusBadRequest, errors.New("bad ID"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		req.Title = strings.TrimSpace(req.Title)
		req.ShortDescription = strings.TrimSpace(req.ShortDescription)
		req.Content = strings.TrimSpace(req.Content)
		if err := validateStruct(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		res, err := b.service.Update(r.Context(), id, req.Title, req.ShortDescription, req.Content)
		if err != nil {
			if errors.As(err, model.ErrNotFound) {
				respondError(w, r, http.StatusNotFound, err)
				return
			}
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusCreated, res)
		return

	}
}
func (b *Posts) deleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondError(w, r, http.StatusBadRequest, errors.New("bad ID"))
			return
		}

		err = b.service.Delete(r.Context(), id)
		if err != nil {
			if errors.As(err, model.ErrPostsExists) {
				respondError(w, r, http.StatusBadRequest, err)
				return
			}
			respondError(w, r, http.StatusNotFound, err)
			return
		}
		respond(w, r, http.StatusOK, nil)
		return
	}
}
