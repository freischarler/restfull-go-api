package data

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *mongo.Client
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	/*err = MakeMigration(db)
	if err != nil {
		log.Panic(err)
	}*/

	data = &Data{
		DB: db,
	}
	
}

func New() *Data {
	once.Do(initDB)

	return data
}

// Close closes the resources used by data.
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Disconnect(context.TODO())
}

func Collection() *mongo.Collection {
	return data.DB.Database("pokeworld").Collection("pokemons")
}
