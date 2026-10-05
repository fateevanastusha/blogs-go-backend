package e2e

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
	"github.com/fateevanastusha/blogs-go-backend/internal/server"
	"github.com/stretchr/testify/require"
)

func getAllEmptyBlogsCheck(t *testing.T, h http.Handler, _ *httptest.ResponseRecorder) {
	rec := DoRequest(t, h, http.MethodGet, "/blogs", nil)
	require.Equal(t, http.StatusOK, rec.Code)
	var blogs []model.Blog
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blogs))
	require.Equal(t, len(blogs), 0)
}

// func notFoundGetCheck(t *testing.T, h http.Handler, id int) {
// 	rec := DoRequest(t, h, http.MethodGet, "/blogs/"+strconv.Itoa(id), nil)
// 	require.Equal(t, http.StatusNotFound, rec.Code)
// }

// func create(t *testing.T, h http.Handler) model.Blog {
// 	rec := DoRequest(t, h, http.MethodPost, "/blogs", map[string]string{"name": "test blog", "description": "test description", "websiteURL": "https://x.com"})
// 	require.Equal(t, http.StatusCreated, rec.Code)
// 	var blog model.Blog
// 	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blog))
// 	return blog
// }

// func delete(t *testing.T, h http.Handler, id int) {
// 	rec := DoRequest(t, h, http.MethodDelete, "/blogs/"+strconv.Itoa(id), nil)
// 	require.Equal(t, http.StatusNoContent, rec.Code)
// }

func TestCreateBlogs(t *testing.T) {

	cases := []Case{
		{
			body:         map[string]string{"name": "test blog", "description": "test description", "websiteURL": "https://x.com"},
			description:  "success case",
			expectStatus: http.StatusCreated,
			after: func(t *testing.T, s http.Handler, rec *httptest.ResponseRecorder) {
				var blog model.Blog
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blog))
				require.Positive(t, blog.ID)
				require.Equal(t, "test blog", blog.Name)
				require.Equal(t, "test description", blog.Description)
				require.Equal(t, "https://x.com", blog.WebsiteURL)

				rec = DoRequest(t, s, http.MethodGet, "/blogs/"+strconv.Itoa(blog.ID), nil)
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			body:         map[string]string{"description": "d", "websiteURL": "https://x.com"},
			description:  "no name",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyBlogsCheck,
		},
		{
			body:         map[string]string{"name": "lala", "description": "d", "websiteURL": "hts/x.com"},
			description:  "invalid url",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyBlogsCheck,
		},
		{
			body:         map[string]string{"name": "lalalalalalalalalalalalalaalalalala", "description": "d", "websiteURL": "https://x.com"},
			description:  "too long name",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyBlogsCheck,
		},
		{
			body:         map[string]string{},
			description:  "empty body",
			expectStatus: http.StatusBadRequest,
			after:        getAllEmptyBlogsCheck,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			s, err := server.New()
			require.NoError(t, err)

			rec := DoRequest(t, s, http.MethodPost, "/blogs", c.body)
			require.Equal(t, c.expectStatus, rec.Code)
			if c.after != nil {
				c.after(t, s, rec)
			}

		})
	}
}

func TestPutBlogs(t *testing.T) {
	cases := []Case{
		{},
		{},
		{},
		{},
	}

}

func TestDeleteBlogs(t *testing.T) {

}
