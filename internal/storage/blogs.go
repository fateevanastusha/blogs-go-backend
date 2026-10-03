package storage

import (
	"context"
	"sync"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type blogsRepo struct {
	blogs  []model.Blog
	mu     sync.Mutex
	nextID int
}

func NewBlogsRepo() *blogsRepo {
	return &blogsRepo{nextID: 1}
}

func (br *blogsRepo) findIndexByID(id int) (int, error) {
	for i, b := range br.blogs {
		if b.ID == id {
			return i, nil
		}
	}
	return -1, model.ErrNotFound

}

func (br *blogsRepo) Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	blog := model.Blog{
		ID:          br.nextID,
		Name:        name,
		Description: description,
		WebsiteURL:  websiteURL,
	}

	br.blogs = append(br.blogs, blog)
	br.nextID++
	return blog, nil
}

func (br *blogsRepo) Update(ctx context.Context, id int, name, description, websiteURL string) (model.Blog, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return model.Blog{}, err
	}
	blog := &br.blogs[index]
	blog.Name = name
	blog.Description = description
	blog.WebsiteURL = websiteURL
	return *blog, nil
}

func (br *blogsRepo) Delete(ctx context.Context, id int) error {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return err
	}
	br.blogs = append(br.blogs[:index], br.blogs[index+1:]...)
	return nil
}

func (br *blogsRepo) GetByID(ctx context.Context, id int) (model.Blog, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return model.Blog{}, err
	}

	blog := br.blogs[index]
	return blog, nil
}

func (br *blogsRepo) GetAll(ctx context.Context) []model.Blog {
	br.mu.Lock()
	defer br.mu.Unlock()

	res := make([]model.Blog, len(br.blogs))
	copy(res, br.blogs)
	return res
}
