package post

import (
	"context"
	"github.com/jmoiron/sqlx"
	"news-app/pkg/model"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

const (
	QueryGetAll = `
SELECT 
    id, 
    title, 
    content, 
    created_at,
    updated_at
FROM post
    ORDER BY id
    LIMIT $1 OFFSET $2;`

	QueryGetById = `
SELECT 
    id, 
    title, 
    content, 
    created_at,
    updated_at
FROM post
	WHERE id = $1
    ORDER BY id
    LIMIT $2 OFFSET $3;
	`

	QueryInsert = `
INSERT INTO post (title, content)
VALUES ($1, $2);`

	QueryUpdate = `
 UPDATE post
        SET title = $1, content = $2, updated_at = NOW()
        WHERE id = $3
        RETURNING id, title, content, created_at, updated_at;`

	QueryDelete = `
 DELETE FROM post
        WHERE id = $1;`
)

// добавить тайм-аут в контекст
func (p *PostPostgres) GetAll(ctx context.Context, args model.PaginationArgs) ([]model.Post, error) {
	var result []model.Post

	err := p.db.SelectContext(ctx, &result, QueryGetAll, args.Limit, args.Offset)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PostPostgres) GetById(ctx context.Context, args model.PaginationByID) (*model.Post, error) {
	var (
		posts  []model.Post
		result *model.Post
	)

	err := p.db.SelectContext(ctx, &posts, QueryGetById, args.ID, args.Limit, args.Offset)
	if err != nil {
		return nil, err
	}
	if len(posts) > 0 {
		result = &posts[0]
	}

	return result, nil
}

func (p *PostPostgres) Create(ctx context.Context, post model.Post) error {
	_, err := p.db.ExecContext(ctx, QueryInsert, post.Title, post.Content)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostPostgres) Update(ctx context.Context, post model.Post) (model.Post, error) {
	row := p.db.QueryRowContext(ctx, QueryUpdate, post.Title, post.Content, post.ID)

	var updatedPost model.Post

	err := row.Scan(
		&updatedPost.ID,
		&updatedPost.Title,
		&updatedPost.Content,
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt,
	)
	if err != nil {
		return model.Post{}, err
	}

	return updatedPost, nil
}

func (p *PostPostgres) Delete(ctx context.Context, id int) error {
	_, err := p.db.ExecContext(ctx, QueryDelete, id)
	if err != nil {
		return err
	}

	return nil
}
