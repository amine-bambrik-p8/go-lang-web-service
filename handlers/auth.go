package handlers

import (
	"encoding/json"
	"net/http"

	commonHTTP "github.com/amine-bambrik-p8/go-lang-web-service/common/http"
	"github.com/amine-bambrik-p8/go-lang-web-service/models"
	"github.com/amine-bambrik-p8/go-lang-web-service/services"
)

// TODO should hook up the Auth routes
func (c *AuthController) HookHandlers() http.Handler {
	return nil
}

// IAuthController interface of Authentication routes handler
type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	HookHandlers() http.Handler
}

// RoutesController struct that represents the Authentication routes handler
type AuthController struct {
}

// the Authentication routes handler
// NOTE: the AuthHandler can be easily mocked since it's made public as var
var (
	AuthHandler IAuthController
)

func init() {
	AuthHandler = &AuthController{}
}

// Login checks for User Existance and Password validity
// TODOD should return Invalid Username or Password response incase username or passwrod don't match
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	var user models.NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		commonHTTP.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}
	result := services.Users.FindUser(&user)
	tokenString, err := result.GenerateJWT()
	response := map[string]string{
		"token": tokenString,
	}
	commonHTTP.SendJSON(w, r, response, http.StatusOK)
}
