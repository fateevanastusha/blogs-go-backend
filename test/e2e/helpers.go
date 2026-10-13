package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fateevanastusha/blogs-go-backend/internal/config"
	"github.com/stretchr/testify/require"
)

var testConfig = config.Config{
	Port:         "3005",
	AuthLogin:    "admin",
	AuthPassword: "qwerty",
}

func DoRequest(t *testing.T, h http.Handler, method, path string, body any, auth bool) *httptest.ResponseRecorder {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		require.NoError(t, json.NewEncoder(&buf).Encode(body))
	}
	req := httptest.NewRequest(method, path, &buf)
	if auth {
		req.SetBasicAuth(testConfig.AuthLogin, testConfig.AuthPassword)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

type blogCase struct {
	body         map[string]string
	description  string
	expectStatus int
	id           int // для запросов по id; 0 — взять id созданного в тесте блога
	after        func(t *testing.T, h http.Handler, rec *httptest.ResponseRecorder)
}
type postCase struct {
	body         map[string]any
	description  string
	expectStatus int
	after        func(t *testing.T, h http.Handler, rec *httptest.ResponseRecorder)
}
