package storage

import (
	"context"
	"sync"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type BlogsRepo struct {
	blogs  []model.Blog
	mu     sync.Mutex
	nextID int
}

func NewBlogsRepo() *BlogsRepo {
	return &BlogsRepo{nextID: 1}
}

func (br *BlogsRepo) findIndexByID(ID int) (int, error) {
	for i, b := range br.blogs {
		if b.ID == ID {
			return i, nil
		}
	}
	return -1, model.ErrNotFound

}

func (br *BlogsRepo) Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error) {
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

func (br *BlogsRepo) Update(ctx context.Context, ID int, name, description, websiteURL string) (model.Blog, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(ID)
	if err != nil {
		return model.Blog{}, err
	}
	blog := &br.blogs[index]
	blog.Name = name
	blog.Description = description
	blog.WebsiteURL = websiteURL
	return *blog, nil
}

func (br *BlogsRepo) Delete(ctx context.Context, ID int) error {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(ID)
	if err != nil {
		return err
	}
	br.blogs = append(br.blogs[:index], br.blogs[index+1:]...)
	return nil
}

func (br *BlogsRepo) GetByID(ctx context.Context, ID int) (model.Blog, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(ID)
	if err != nil {
		return model.Blog{}, err
	}

	blog := br.blogs[index]
	return blog, nil
}

func (br *BlogsRepo) GetAll(ctx context.Context) []model.Blog {
	br.mu.Lock()
	defer br.mu.Unlock()

	res := make([]model.Blog, len(br.blogs))
	copy(res, br.blogs)
	return res
}
