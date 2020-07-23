package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	commonHTTP "github.com/amine-bambrik-p8/go-lang-web-service/common/http"
	"github.com/amine-bambrik-p8/go-lang-web-service/models"
	"github.com/amine-bambrik-p8/go-lang-web-service/services"
	"github.com/gorilla/mux"
)

// TODO Should refector to generic CRUDController
// Hooks up all the api's routes for the UserController
func (c *UserController) HookHandlers() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", c.Index).Methods("GET")
	router.HandleFunc("/{id}", c.Show).Methods("GET")
	router.HandleFunc("/", c.Create).Methods("POST")
	router.HandleFunc("/{id}", c.Update).Methods("PUT")
	router.HandleFunc("/{id}", c.Delete).Methods("DELETE")
	return router
}

// IUserController interface of the Users handler
type IUserController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	HookHandlers() http.Handler
}

// UserController struct that represents the Users routes handler
type UserController struct {
}

// the Users routes handler
// NOTE: the UserHandler can be easily mocked since it's made public as var
var (
	UserHandler IUserController
)

func init() {
	UserHandler = &UserController{}
}

// Index returns the list of all users
func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users := services.Users.AllUsers()
	commonHTTP.SendJSON(w, r, users, http.StatusOK)
}

// Show return the user by selected by the id set as URL param
// TODO Should check if no user was found and return NotFound status in that case
func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Print("Error Parsing User id")
		commonHTTP.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}
	user := services.Users.FindUserById(idUint)
	commonHTTP.SendJSON(w, r, user, http.StatusOK)

}

// Create creates a new User and returns the new user
// TODO should implement check for username already taken
func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		commonHTTP.SendJSON(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	result := services.Users.NewUser(&user)
	commonHTTP.SendJSON(w, r, result, http.StatusOK)
}

// Delete deletes a user and returns the deleted user
// TODO should implement check for user Not Found and return appropriate status code
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Print("Error Parsing User id")
		commonHTTP.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}
	user := services.Users.DeleteUser(idUint)
	commonHTTP.SendJSON(w, r, user, http.StatusOK)
}

// Update updates the user that matches the username
// TODO should implement check for user not found and retrun status NotFound
func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var user models.NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		commonHTTP.SendJSON(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	result := services.Users.UpdateUser(&user)
	commonHTTP.SendJSON(w, r, result, http.StatusOK)
}
