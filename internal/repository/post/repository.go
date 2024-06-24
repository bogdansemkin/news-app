package post

import (
	"context"
	"github.com/jmoiron/sqlx"
	"news-app/pkg/model"
)

type PostTool interface {
	Create(ctx context.Context, post model.Post) error
	GetAll(ctx context.Context, args model.PaginationArgs) ([]model.Post, error)
	GetById(ctx context.Context, args model.PaginationByID) (*model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	PostTool
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{PostTool: NewPostPostgres(db)}
}
