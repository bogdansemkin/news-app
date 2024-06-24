package service

import (
	"context"
	"github.com/pkg/errors"
	"news-app/internal/repository/post"
	"news-app/pkg/model"
)

type ServicePost struct {
	repo post.PostTool
}

func NewPostService(repo post.PostTool) *ServicePost {
	return &ServicePost{repo: repo}
}

const (
	defaultLimit  = 10
	defaultOffset = 0
)

func (s *ServicePost) GetAll(ctx context.Context, args model.PaginationArgs) ([]model.Post, error) {
	validArgs(&args)

	result, err := s.repo.GetAll(ctx, args)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New(model.ErrPostsNotFound)
	}

	return result, nil
}

func (s *ServicePost) GetById(ctx context.Context, args model.PaginationByID) (*model.Post, error) {
	validArgs(&args.PaginationArgs)
	if args.ID < 0 {
		return nil, errors.New(model.ErrPostInvalidId)
	}

	result, err := s.repo.GetById(ctx, args)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New(model.ErrSinglePostNotFound)
	}

	return result, nil
}

func (s *ServicePost) Create(ctx context.Context, post model.Post) error {
	if !validJSON(&post) {
		return errors.New(model.ErrPostInvalidModel)
	}

	err := s.repo.Create(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServicePost) Update(ctx context.Context, post model.Post) (model.Post, error) {
	if !validJSON(&post) {
		return model.Post{}, errors.New(model.ErrPostInvalidModel)
	}

	result, err := s.repo.Update(ctx, post)
	if err != nil {
		return model.Post{}, err
	}

	return result, nil
}

func (s *ServicePost) Delete(ctx context.Context, id int) error {
	if id < 0 {
		return errors.New(model.ErrPostInvalidId)
	}
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func validArgs(args *model.PaginationArgs) {
	if args.Limit <= 0 {
		args.Limit = defaultLimit
	}
	if args.Offset < 0 {
		args.Offset = defaultOffset
	}
}

func validJSON(post *model.Post) bool {
	if post != nil {
		if post.Title == "" || post.Content == "" {
			return false
		}
	} else {
		return false
	}

	return true
}
