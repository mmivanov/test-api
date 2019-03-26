package post

import (
	"test-api/internal/model"
)

type IPostRepository interface {
	Posts() (*[]model.Post, error)
	Post(id int32) (*model.Post, error)
	CreatePost(post *model.Post) error
}

type service struct {
	repo IPostRepository
}

type post struct {
	ID     int32  `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Text   string `json:"text,omitempty"`
	Status int    `json:"status,omitempty"`
}

func New(r IPostRepository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) Posts() (*[]model.Post, error) {
	return s.repo.Posts()
}

func (s *service) Post(id int32) (*model.Post, error) {
	return s.repo.Post(id)
}

func (s *service) CreatePost(post *model.Post) error {
	return s.repo.CreatePost(post)
}
