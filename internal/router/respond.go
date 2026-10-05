package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

func respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	var ve *model.ValidationError
	if errors.As(err, &ve) {
		respond(w, r, code, map[string]any{"errors": ve.Fields})
		return
	}
	respond(w, r, code, map[string]string{"error": err.Error()})
}
