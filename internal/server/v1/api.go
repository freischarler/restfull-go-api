package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/martinpaz/restfulapi/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	cr := &CommentRouter{
		Repository: &data.CommentRepository{
			Data: data.New(),
		},
	}

	r.Mount("/comment", cr.Routes())

	return r
}
