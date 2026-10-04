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

type blogsService interface {
	Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error)
	Update(ctx context.Context, ID int, name, description, websiteURL string) (model.Blog, error)
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (model.Blog, error)
	GetAll(ctx context.Context) ([]model.Blog, error)
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
func (b *Blogs) getByID() http.HandlerFunc {
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
func (b *Blogs) create() http.HandlerFunc {
	type request struct {
		Name        string `json:"name" validate:"required,max=15"`
		Description string `json:"description" validate:"required,max=500"`
		WebsiteURL  string `json:"websiteURL" validate:"required,max=100,url,startswith=https://"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		req.Name = strings.TrimSpace(req.Name)
		req.Description = strings.TrimSpace(req.Description)
		req.WebsiteURL = strings.TrimSpace(req.WebsiteURL)
		if err := validateStruct(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		res, err := b.service.Create(r.Context(), req.Name, req.Description, req.WebsiteURL)
		if err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}
		respond(w, r, http.StatusCreated, res)
		return
	}
}
func (b *Blogs) putByID() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		WebsiteURL  string `json:"websiteURL"`
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

		req.Name = strings.TrimSpace(req.Name)
		req.Description = strings.TrimSpace(req.Description)
		req.WebsiteURL = strings.TrimSpace(req.WebsiteURL)
		if err := validateStruct(req); err != nil {
			respondError(w, r, http.StatusBadRequest, err)
			return
		}

		res, err := b.service.Update(r.Context(), id, req.Name, req.Description, req.WebsiteURL)
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
func (b *Blogs) deleteByID() http.HandlerFunc {
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
		respond(w, r, http.StatusNoContent, nil)
		return
	}
}
