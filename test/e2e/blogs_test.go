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

func TestCreateBlogs(t *testing.T) {

	cases := []blogCase{
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
	cases := []blogCase{
		{
			body:         map[string]string{"name": "edited", "description": "edited", "websiteURL": "https://x.com?edited=true"},
			description:  "success case",
			expectStatus: http.StatusCreated,
		},
		{
			body:         map[string]string{"description": "d", "websiteURL": "https://x.com"},
			description:  "no name",
			expectStatus: http.StatusBadRequest,
		},
		{
			body:         map[string]string{"name": "lala", "description": "d", "websiteURL": "hts/x.com"},
			description:  "invalid url",
			expectStatus: http.StatusBadRequest,
		},
		{
			body:         map[string]string{"name": "lalalalalalalalalalalalalaalalalala", "description": "d", "websiteURL": "https://x.com"},
			description:  "too long name",
			expectStatus: http.StatusBadRequest,
		},
		{
			body:         map[string]string{},
			description:  "empty body",
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {

			s, err := server.New()
			require.NoError(t, err)

			newBlog := createBlog(t, s)

			rec := DoRequest(t, s, http.MethodPut, "/blogs/"+strconv.Itoa(newBlog.ID), c.body)
			require.Equal(t, c.expectStatus, rec.Code)
			if c.expectStatus == http.StatusBadRequest {
				notEditedBlog := getBlog(t, s, newBlog.ID)
				require.Equal(t, newBlog.Description, notEditedBlog.Description)
				require.Equal(t, newBlog.Name, notEditedBlog.Name)
				require.Equal(t, newBlog.WebsiteURL, notEditedBlog.WebsiteURL)
			}
			if c.expectStatus == http.StatusCreated {
				editedBlog := getBlog(t, s, newBlog.ID)
				require.Equal(t, c.body["description"], editedBlog.Description)
				require.Equal(t, c.body["name"], editedBlog.Name)
				require.Equal(t, c.body["websiteURL"], editedBlog.WebsiteURL)
			}

			if c.after != nil {
				c.after(t, s, rec)
			}

		})
	}

}

func TestDeleteBlogs(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		newBlog := createBlog(t, s)

		rec := DoRequest(t, s, http.MethodDelete, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusNoContent, rec.Code)
	})
	t.Run("delete already deleted", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		newBlog := createBlog(t, s)

		rec := DoRequest(t, s, http.MethodDelete, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusNoContent, rec.Code)

		rec = DoRequest(t, s, http.MethodDelete, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("delete not existing", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		rec := DoRequest(t, s, http.MethodDelete, "/blogs/434444", nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("delete with existing post", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		newBlog := createBlog(t, s)
		createPost(t, s, newBlog.ID)

		rec := DoRequest(t, s, http.MethodDelete, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestGetByIdBlogs(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		newBlog := createBlog(t, s)

		rec := DoRequest(t, s, http.MethodGet, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusOK, rec.Code)
		var blog model.Blog
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blog))
		require.Equal(t, newBlog.Description, blog.Description)
		require.Equal(t, newBlog.WebsiteURL, blog.WebsiteURL)
		require.Equal(t, newBlog.Name, blog.Name)
	})

	t.Run("get not existing", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		rec := DoRequest(t, s, http.MethodGet, "/blogs/434444", nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("get deleted", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)
		newBlog := createBlog(t, s)
		deleteBlog(t, s, newBlog.ID)

		rec := DoRequest(t, s, http.MethodGet, "/blogs/"+strconv.Itoa(newBlog.ID), nil)
		require.Equal(t, http.StatusNotFound, rec.Code)
	})

}

func TestGetBlogs(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		s, err := server.New()
		require.NoError(t, err)

		var blogs []model.Blog

		rec := DoRequest(t, s, http.MethodGet, "/blogs", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blogs))
		require.Equal(t, 0, len(blogs))

		newBlogs := []model.Blog{createBlog(t, s), createBlog(t, s), createBlog(t, s)}

		rec = DoRequest(t, s, http.MethodGet, "/blogs", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blogs))
		require.Equal(t, 3, len(newBlogs))

		for i := 0; i < len(newBlogs); i++ {
			b1, b2 := newBlogs[i], blogs[i]
			require.Equal(t, b1.Description, b2.Description)
			require.Equal(t, b1.WebsiteURL, b2.WebsiteURL)
			require.Equal(t, b1.Name, b2.Name)
		}

		for i := 0; i < len(newBlogs); i++ {
			deleteBlog(t, s, newBlogs[i].ID)
		}

		rec = DoRequest(t, s, http.MethodGet, "/blogs", nil)
		require.Equal(t, http.StatusOK, rec.Code)
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &blogs))
		require.Equal(t, 0, len(blogs))
	})

}
