package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/martinpaz/restfulapi/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &PokemonRouter{
		Repository: &data.PokemonRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	return r
}
