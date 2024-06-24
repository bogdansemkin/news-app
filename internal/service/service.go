package service

import (
	"context"
	"news-app/internal/repository/post"
	"news-app/pkg/model"
)

type PostTool interface {
	Create(ctx context.Context, post model.Post) error
	GetAll(ctx context.Context, args model.PaginationArgs) ([]model.Post, error)
	GetById(ctx context.Context, args model.PaginationByID) (*model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	PostTool
}

func NewService(repos *post.Repository) *Service {
	return &Service{PostTool: NewPostService(repos.PostTool)}
}
