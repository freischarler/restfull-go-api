package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/martinpaz/restfulapi/pkg/pokemon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PokemonRepository manages the operations with the database that
// correspond to the pokemon model.
type PokemonRepository struct {
	Data *Data
}

// GetAll returns all users.
func (ur *PokemonRepository) GetAll(ctx context.Context) ([]pokemon.Pokemon, error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)

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

// GetOne returns one pokemon by id.
func (ur *PokemonRepository) GetOne(ctx context.Context, id uint) (pokemon.Pokemon, error) {

	//filter := bson.D{primitive.E{Key: "health", Value: id}}
	filter := bson.M{"health": id}
	var p pokemon.Pokemon

	err := Collection().FindOne(ctx, filter).Decode(&p)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	return p, nil
}

// GetByUsername returns one user by username.
func (ur *PokemonRepository) GetByName(ctx context.Context, pokemon_name string) (pokemon.Pokemon, error) {
	filter := bson.M{"name": pokemon_name}
	var p pokemon.Pokemon

	err := Collection().FindOne(ctx, filter).Decode(&p)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	return p, nil
}

// Create adds a new user.
func (ur *PokemonRepository) Create(ctx context.Context, p *pokemon.Pokemon) error {

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	insertResult, err := Collection().InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}

// Update updates a user by id.
func (ur *PokemonRepository) Update(ctx context.Context, id uint, p pokemon.Pokemon) error {
	filter := bson.M{"id": id}

	replacementObj := p

	//This is not an update, this function reemplace the whole object/document
	updateResult, err := Collection().ReplaceOne(ctx, filter, replacementObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

// Delete removes a user by id.
func (ur *PokemonRepository) Delete(ctx context.Context, id uint) error {
	filter := bson.M{"id": id}

	deleteResult, err := Collection().DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

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
