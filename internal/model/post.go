package model

type Post struct {
	ID     int32  `db:"id"`
	Email  string `db:"email,omitempty"`
	Text   string `db:"text,omitempty"`
	Status int    `db:"status,omitempty"`
}
