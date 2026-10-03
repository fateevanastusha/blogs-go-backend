package storage

import (
	"context"
	"sync"

	"github.com/fateevanastusha/blogs-go-backend/internal/model"
)

type PostsRepo struct {
	posts  []model.Post
	mu     sync.Mutex
	nextID int
}

func NewPostsRepo() *PostsRepo {
	return &PostsRepo{nextID: 1}
}

func (pr *PostsRepo) findIndexByID(ID int) (int, error) {
	for i, b := range pr.posts {
		if b.ID == ID {
			return i, nil
		}
	}
	return -1, model.ErrNotFound

}

func (pr *PostsRepo) Create(ctx context.Context, title, shortDescription, content string, blogID int) (model.Post, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	post := model.Post{
		ID:               pr.nextID,
		Title:            title,
		ShortDescription: shortDescription,
		Content:          content,
		BlogID:           blogID,
	}

	pr.posts = append(pr.posts, post)
	pr.nextID++
	return post, nil
}

func (pr *PostsRepo) Update(ctx context.Context, ID int, title, shortDescription, content string) (model.Post, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	index, err := pr.findIndexByID(ID)
	if err != nil {
		return model.Post{}, err
	}
	post := &pr.posts[index]
	post.Title = title
	post.ShortDescription = shortDescription
	post.Content = content
	return *post, nil
}

func (pr *PostsRepo) Delete(ctx context.Context, ID int) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	index, err := pr.findIndexByID(ID)
	if err != nil {
		return err
	}
	pr.posts = append(pr.posts[:index], pr.posts[index+1:]...)
	return nil
}

func (pr *PostsRepo) GetByID(ctx context.Context, ID int) (model.Post, error) {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	index, err := pr.findIndexByID(ID)
	if err != nil {
		return model.Post{}, err
	}

	post := pr.posts[index]
	return post, nil
}

func (pr *PostsRepo) GetAll(ctx context.Context) []model.Post {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	res := make([]model.Post, len(pr.posts))
	copy(res, pr.posts)
	return res
}

func (pr *PostsRepo) GetByBlogID(ctx context.Context, blogID int) []model.Post {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	res := []model.Post{}
	for _, p := range pr.posts {
		if p.BlogID == blogID {
			res = append(res, p)
		}
	}
	return res
}
