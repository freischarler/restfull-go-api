package pokemon

import (
	"context"
)

// Repository handle the CRUD operations with Pokemons.
type Repository interface {
	GetAll(ctx context.Context) ([]Pokemon, error)
	GetOne(ctx context.Context, id uint) (Pokemon, error)
	GetByName(ctx context.Context, Pokemonname string) (Pokemon, error)
	Create(ctx context.Context, Pokemon *Pokemon) error
	Update(ctx context.Context, id uint, Pokemon Pokemon) error
	Delete(ctx context.Context, id uint) error
}
