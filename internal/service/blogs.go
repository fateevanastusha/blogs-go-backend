package service

import (
	"context"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type blogsRepo interface {
	Create(ctx context.Context, name, description, WebsiteURL string) (model.Blog, error)
	Update(ctx context.Context, ID int, name, description, websiteURL string) (model.Blog, error)
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (model.Blog, error)
	GetAll(ctx context.Context) []model.Blog
}

type postReader interface {
	GetByBlogID(ctx context.Context, blogID int) []model.Post
}

type BlogsService struct {
	repo       blogsRepo
	postReader postReader
}

func NewBlogsService(repo blogsRepo, postReader postReader) *BlogsService {
	return &BlogsService{repo: repo, postReader: postReader}
}

func (bs *BlogsService) Create(ctx context.Context, name, description, websiteURL string) (model.Blog, error) {
	return bs.repo.Create(ctx, name, description, websiteURL)
}
func (bs *BlogsService) Update(ctx context.Context, ID int, name, description, websiteURL string) (model.Blog, error) {
	return bs.repo.Update(ctx, ID, name, description, websiteURL)
}
func (bs *BlogsService) Delete(ctx context.Context, ID int) error {
	existingPosts := bs.postReader.GetByBlogID(ctx, ID)
	if len(existingPosts) > 0 {
		return model.ErrPostsExists
	}
	return bs.repo.Delete(ctx, ID)
}
func (bs *BlogsService) GetByID(ctx context.Context, ID int) (model.Blog, error) {
	return bs.repo.GetByID(ctx, ID)
}
func (bs *BlogsService) GetAll(ctx context.Context) []model.Blog {
	return bs.repo.GetAll(ctx)
}
