package service

import (
	"context"
	"errors"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type postsRepo interface {
	Create(ctx context.Context, title, shortDescription, content string, blogId int) (model.Post, error)
	Update(ctx context.Context, id int, title, shortDescription, content string) (model.Post, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (model.Post, error)
	GetAll(ctx context.Context) []model.Post
}

type blogsRepo interface {
	GetByID(ctx context.Context, id int) (model.Post, error)
}

type postsService struct {
	repo      postsRepo
	blogsRepo blogsRepo
}

func NewPostsService(repo postsRepo, blogsRepo blogsRepo) *postsService {
	return &postsService{repo: repo, blogsRepo: blogsRepo}
}

func (bs *postsService) Create(ctx context.Context, title, shortDescription, content string, blogId int) (model.Post, error) {
	if _, err := bs.blogsRepo.GetByID(ctx, blogId); errors.Is(err, model.ErrNotFound) {
		return model.Post{}, model.ErrBlogNotFound
	}
	return bs.repo.Create(ctx, title, shortDescription, content, blogId)
}
func (bs *postsService) Update(ctx context.Context, id int, title, shortDescription, content string) (model.Post, error) {
	return bs.repo.Update(ctx, id, title, shortDescription, content)
}
func (bs *postsService) Delete(ctx context.Context, id int) error {
	return bs.repo.Delete(ctx, id)
}
func (bs *postsService) GetByID(ctx context.Context, id int) (model.Post, error) {

	return bs.repo.GetByID(ctx, id)
}
func (bs *postsService) GetAll(ctx context.Context) []model.Post {
	return bs.repo.GetAll(ctx)
}
