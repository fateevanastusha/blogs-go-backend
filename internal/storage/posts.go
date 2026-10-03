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

func (br *PostsRepo) findIndexByID(id int) (int, error) {
	for i, b := range br.posts {
		if b.ID == id {
			return i, nil
		}
	}
	return -1, model.ErrNotFound

}

func (br *PostsRepo) Create(ctx context.Context, title, shortDescription, content string, blogId int) (model.Post, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	post := model.Post{
		ID:               br.nextID,
		Title:            title,
		ShortDescription: shortDescription,
		Content:          content,
		BlogID:           blogId,
	}

	br.posts = append(br.posts, post)
	br.nextID++
	return post, nil
}

func (br *PostsRepo) Update(ctx context.Context, id int, title, shortDescription, content string) (model.Post, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return model.Post{}, err
	}
	post := &br.posts[index]
	post.Title = title
	post.ShortDescription = shortDescription
	post.Content = content
	return *post, nil
}

func (br *PostsRepo) Delete(ctx context.Context, id int) error {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return err
	}
	br.posts = append(br.posts[:index], br.posts[index+1:]...)
	return nil
}

func (br *PostsRepo) GetByID(ctx context.Context, id int) (model.Post, error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	index, err := br.findIndexByID(id)
	if err != nil {
		return model.Post{}, err
	}

	post := br.posts[index]
	return post, nil
}

func (br *PostsRepo) GetAll(ctx context.Context) []model.Post {
	br.mu.Lock()
	defer br.mu.Unlock()

	res := make([]model.Post, len(br.posts))
	copy(res, br.posts)
	return res
}
