package data

import (
	"context"
	"log"

	"github.com/martinpaz/restfulapi/pkg/pokemon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PokemonRepository manages the operations with the database that
// correspond to the user model.
type PokemonRepository struct {
	Data *Data
}

// GetAll returns all users.
func (ur *PokemonRepository) GetAll(ctx context.Context) ([]pokemon.Pokemon, error) {
	findOptions := options.Find()
	findOptions.SetLimit(2)

	rows, err := Collection().Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	var pokemons []pokemon.Pokemon

	for rows.Next(ctx) {
		var p pokemon.Pokemon
		err := rows.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		pokemons = append(pokemons, p)
	}

	return pokemons, nil
}

// GetOne returns one user by id.
func (ur *PokemonRepository) GetOne(ctx context.Context, id uint) (pokemon.Pokemon, error) {
	/*q := `
	SELECT id, first_name, last_name, username, email, picture,
		created_at, updated_at
		FROM users WHERE id = $1;
	`

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	*/var u pokemon.Pokemon
	/*
		err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email,
			&u.Picture, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return pokemon.Pokemon{}, err
		}*/

	return u, nil
}

// GetByUsername returns one user by username.
func (ur *PokemonRepository) GetByName(ctx context.Context, username string) (pokemon.Pokemon, error) {
	/*q := `
	SELECT id, first_name, last_name, username, email, picture,
		password, created_at, updated_at
		FROM users WHERE username = $1;
	`

	row := ur.Data.DB.QueryRowContext(ctx, q, username)

	*/var u pokemon.Pokemon
	/*err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
		&u.Email, &u.Picture, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return pokemon.Pokemon{}, err
	}*/

	return u, nil
}

// Create adds a new user.
func (ur *PokemonRepository) Create(ctx context.Context, u *pokemon.Pokemon) error {
	/*q := `
	INSERT INTO users (first_name, last_name, username, email, picture, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	if u.Picture == "" {
		u.Picture = "https://placekitten.com/g/300/300"
	}

	if err := u.HashPassword(); err != nil {
		return err
	}

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.PasswordHash, time.Now(), time.Now(),
	)

	err = row.Scan(&u.ID)
	if err != nil {
		return err
	}*/

	return nil
}

// Update updates a user by id.
func (ur *PokemonRepository) Update(ctx context.Context, id uint, u pokemon.Pokemon) error {
	/*q := `
	UPDATE users set first_name=$1, last_name=$2, email=$3, picture=$4, updated_at=$5
		WHERE id=$6;
	`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.FirstName, u.LastName, u.Email,
		u.Picture, time.Now(), id,
	)
	if err != nil {
		return err
	}*/

	return nil
}

// Delete removes a user by id.
func (ur *PokemonRepository) Delete(ctx context.Context, id uint) error {
	/*q := `DELETE FROM users WHERE id=$1;`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}*/

	return nil
}
