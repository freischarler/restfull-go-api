package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/martinpaz/restfulapi/pkg/claim"
	"github.com/martinpaz/restfulapi/pkg/pokemon"
	"github.com/martinpaz/restfulapi/pkg/response"
)

// PokemonRouter is the router of the users.
type PokemonRouter struct {
	Repository pokemon.Repository
}

// Routes returns user router with each endpoint.
func (ur *PokemonRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ur.GetAllHandler)

	r.Post("/", ur.CreateHandler)

	r.Get("/{id}", ur.GetOneHandler)

	//r.With(middleware.Authorizator).Put("/{id}", ur.UpdateHandler)
	r.Put("/{id}", ur.UpdateHandler)

	//r.With(middleware.Authorizator).Delete("/{id}", ur.DeleteHandler)
	r.Delete("/{id}", ur.DeleteHandler)

	r.Post("/login/", ur.LoginHandler)

	return r
}

// LoginHandler search user and return a jwt.
func (ur *PokemonRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u pokemon.Pokemon

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	storedUser, err := ur.Repository.GetByName(ctx, u.Name)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	if !storedUser.PasswordMatch(u.Password) {
		response.HTTPError(w, r, http.StatusBadRequest, "password don't match")
		return
	}

	c := claim.Claim{ID: int(storedUser.ID)}
	//token, err := c.GetToken(os.Getenv("SIGNING_STRING"))
	token, err := c.GetToken("SECRET")
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"token": token})
}

// CreateHandler Create a new user.
func (ur *PokemonRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var u pokemon.Pokemon
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	u.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), u.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"user": u})
}

// GetAllHandler response all the users.
func (ur *PokemonRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"users": users})
}

// GetOneHandler response one user by id.
func (ur *PokemonRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	u, err := ur.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"pokemon": u})
}

// UpdateHandler update a stored user by id.
func (ur *PokemonRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var u pokemon.Pokemon
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, uint(id), u)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a user by ID.
func (ur *PokemonRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}
