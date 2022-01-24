package comment

import (
	"context"
)

// Repository handle the CRUD operations with Users.
type Repository interface {
	Create(ctx context.Context, comment *Comment) error
	GetOne(ctx context.Context, id uint) ([]Comment, error)
	Delete(ctx context.Context, id uint) error
}
