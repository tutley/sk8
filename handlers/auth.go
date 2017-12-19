package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tutley/sk8/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Auth serves as the anchor for all the handlers based on examples routes
type Auth struct {
	Db *mgo.Database
	Ta *jwtauth.JwtAuth
}

// Routes creates a REST router for the Auth resource
// /auth
// /api/auth/...
func (rs Auth) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.PostLogin)      // POST /api/auth
	r.Post("/new", rs.PostNewUser) // POST /api/auth/new
	return r
}

// PostLogin will return a JWT token for the user that signed in.
func (rs Auth) PostLogin(w http.ResponseWriter, r *http.Request) {
	var authTry = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&authTry)
	if err != nil {
		log.Println("error parsing json request for login. ", err)
		http.Error(w, "Something was wrong with the request body.", http.StatusBadRequest)
		return
	}

	user, err := models.FindUserByUsername(authTry.Username, rs.Db)
	if err != nil {
		log.Println("User Not Found: ", authTry.Username)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// Check the password

	err = user.CheckPassword(authTry.Password)
	if err != nil {
		log.Printf("Invalid password for User: %+v.", authTry.Username)
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := rs.Ta.Encode(jwtauth.Claims{"user_id": user.ID})

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"token\": \"%s\"}", tokenString)
}

// PostNewUser will create a new user
func (rs Auth) PostNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser = models.User{}
	// {"username":"something","password":"something","email":"something"}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println("error with new user input: ", err)
		http.Error(w, "Something was wrong with the new user request.", http.StatusBadRequest)
		return
	}
	newUser.ID = bson.NewObjectId()
	log.Println("password: ", newUser.Password)
	err = newUser.Insert(rs.Db)
	if err != nil {
		log.Println("error adding new user: ", err)
		http.Error(w, "Something went wrong while adding the new user.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result":"success"}`))
}
