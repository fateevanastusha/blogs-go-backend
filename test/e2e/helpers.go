package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func DoRequest(t *testing.T, h http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		require.NoError(t, json.NewEncoder(&buf).Encode(body))
	}
	req := httptest.NewRequest(method, path, &buf)
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
