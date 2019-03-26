package post

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"test-api/internal/model"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Posts() (*[]model.Post, error) {
	data := make([]model.Post, 0)

	err := r.db.Select(&data, "SELECT * FROM posts WHERE 1")

	if err == sql.ErrNoRows {
		return &data, nil
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Post(id int32) (*model.Post, error) {
	data := model.Post{}

	err := r.db.Get(&data, "SELECT * FROM posts WHERE id = $1", id)

	if err == sql.ErrNoRows {
		return &data, nil
	}

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) CreatePost(post *model.Post) error {
	var lastInsertId int32

	result := r.db.QueryRow(
		"INSERT INTO posts (email, text, status) VALUES ($1, $2, $3) RETURNING id",
		post.Email,
		post.Text,
		post.Status,
	).Scan(&lastInsertId)

	if result == sql.ErrNoRows {
		return errors.New("failed to create new post")
	}

	if result != nil {
		return result
	}

	post.ID = lastInsertId

	return nil
}
