package model

import "time"

type PaginationArgs struct {
	Limit  int
	Offset int
}

type PaginationByID struct {
	ID int
	PaginationArgs
}

type Post struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

const (
	ErrSinglePostNotFound = "post not found"
	ErrPostsNotFound      = "posts not found"
	ErrPostInvalidModel   = "empty title or content"
	ErrPostInvalidId      = "invalid id "
)
