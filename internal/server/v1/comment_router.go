package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/martinpaz/restfulapi/pkg/comment"
	"github.com/martinpaz/restfulapi/pkg/response"
)

// UserRouter is the router of the users.
type CommentRouter struct {
	Repository comment.Repository
}

// Routes returns user router with each endpoint.
func (cr *CommentRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", cr.CreateHandler)

	r.Get("/{id}", cr.GetOneHandler)

	r.Delete("/{id}", cr.DeleteHandler)

	return r
}

// CreateHandler Create a new comment.
func (cr *CommentRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var u comment.Comment
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = cr.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), u.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"comment": u})
}

// GetOneHandler response one post by id.
func (cr *CommentRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := cr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"comment": p})
}

// DeleteHandler Remove a user by ID.
func (cr *CommentRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = cr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}
