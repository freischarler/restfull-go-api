package post

import "context" //the documentation is https://go.dev/blog/context

// Repository handle the CRUD operations with Posts.
type Repository interface {
	GetAll(ctx context.Context) ([]Post, error)
	GetOne(ctx context.Context, id uint) (Post, error)
	GetByUser(ctx context.Context, userID uint) ([]Post, error)
	GetByType(ctx context.Context, recipeTYPE string) ([]Post, error)
	GetRank(ctx context.Context) ([]Post, error)
	Create(ctx context.Context, post *Post) error
	UpdateLike(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, post Post) error
	Delete(ctx context.Context, id uint) error
}
