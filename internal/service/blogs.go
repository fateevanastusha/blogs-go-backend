package service

import (
	"context"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type blogsRepo interface {
	Create(ctx context.Context, name, description, WebsiteURL string) (model.Blog, error)
	Update(ctx context.Context, id int, name, description, websiteURL string) (model.Blog, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (model.Blog, error)
	GetAll(ctx context.Context) []model.Blog
}

type blogsService struct {
	repo blogsRepo
}

func NewBlogsService(repo blogsRepo) *blogsService {
	return &blogsService{repo: repo}
}

func (bs *blogsService) Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error) {
	return bs.repo.Create(ctx, name, description, websiteURL)
}
func (bs *blogsService) Update(ctx context.Context, id int, name, description, websiteURL string) (model.Blog, error) {
	return bs.repo.Update(ctx, id, name, description, websiteURL)
}
func (bs *blogsService) Delete(ctx context.Context, id int) error {
	return bs.repo.Delete(ctx, id)
}
func (bs *blogsService) GetByID(ctx context.Context, id int) (model.Blog, error) {
	return bs.repo.GetByID(ctx, id)
}
func (bs *blogsService) GetAll(ctx context.Context) []model.Blog {
	return bs.repo.GetAll(ctx)
}
