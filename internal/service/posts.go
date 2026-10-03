package service

import (
	"context"
	"errors"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type postsRepo interface {
	Create(ctx context.Context, title, shortDescription, content string, blogID int) (model.Post, error)
	Update(ctx context.Context, ID int, title, shortDescription, content string) (model.Post, error)
	Delete(ctx context.Context, ID int) error
	GetByID(ctx context.Context, ID int) (model.Post, error)
	GetAll(ctx context.Context) []model.Post
	GetByBlogID(ctx context.Context, blogID int) []model.Post
}

type blogsReader interface {
	GetByID(ctx context.Context, ID int) (model.Blog, error)
}

type PostsService struct {
	repo        postsRepo
	blogsReader blogsReader
}

func NewPostsService(repo postsRepo, blogsReader blogsReader) *PostsService {
	return &PostsService{repo: repo, blogsReader: blogsReader}
}

func (ps *PostsService) Create(ctx context.Context, title, shortDescription, content string, blogID int) (model.Post, error) {
	if _, err := ps.blogsReader.GetByID(ctx, blogID); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return model.Post{}, model.ErrBlogNotFound
		}
		return model.Post{}, err
	}
	return ps.repo.Create(ctx, title, shortDescription, content, blogID)
}
func (ps *PostsService) Update(ctx context.Context, ID int, title, shortDescription, content string) (model.Post, error) {
	return ps.repo.Update(ctx, ID, title, shortDescription, content)
}
func (ps *PostsService) Delete(ctx context.Context, ID int) error {
	return ps.repo.Delete(ctx, ID)
}
func (ps *PostsService) GetByID(ctx context.Context, ID int) (model.Post, error) {
	return ps.repo.GetByID(ctx, ID)
}
func (ps *PostsService) GetAll(ctx context.Context) []model.Post {
	return ps.repo.GetAll(ctx)
}
func (ps *PostsService) GetByBlogID(ctx context.Context, blogID int) []model.Post {
	return ps.repo.GetByBlogID(ctx, blogID)
}
