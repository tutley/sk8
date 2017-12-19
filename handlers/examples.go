package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/tutley/sk8/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Examples serves as the anchor for all the handlers based on examples routes
type Examples struct {
	Db *mgo.Database
	Ta *jwtauth.JwtAuth
}

// Routes creates a REST router for the examples resource
// /examples
func (rs Examples) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /examples - read a list of examples
	r.Post("/", rs.Create) // POST /examples - create a new example and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /examples/{id} - read a single example by :id
		r.Put("/", rs.Update)    // PUT /examples/{id} - update a single example by :id
		r.Delete("/", rs.Delete) // DELETE /examples/{id} - delete a single example by :id
	})
	return r
}

// List lists all the examples
func (rs Examples) List(w http.ResponseWriter, r *http.Request) {
	// Grab examples from DB
	examples, err := models.ListExamples(rs.Db)
	if err != nil {
		log.Println("Error getting examples: ", err)
		http.Error(w, "Error: Couldn't fetch the examples from the database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(&examples)
	if e != nil {
		log.Println("Error encoding json response: ", e)
		http.Error(w, "Error: we got a bunch of gobbldygook from the database", http.StatusInternalServerError)
		return
	}

}

// Create will make a new example
func (rs Examples) Create(w http.ResponseWriter, r *http.Request) {
	// go ahead and push it to the db and return success or error
	example := models.Example{}

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		log.Println("error parsing json request ", err)
		http.Error(w, "There was an error with your submission message", http.StatusBadRequest)
		return
	}
	example.ID = bson.NewObjectId()

	err = example.Insert(rs.Db)
	if err != nil {
		log.Println("error inserting example: ", err)
		http.Error(w, "The database crapped out when trying to save the example", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(&example)
	if e != nil {
		http.Error(w, "The example was saved but we messed up with the return message", http.StatusInternalServerError)
		return
	}
}

// Get will get an example
func (rs Examples) Get(w http.ResponseWriter, r *http.Request) {
	// Grab example from DB
	id := chi.URLParam(r, "id")
	if id == "post" {
		w.WriteHeader(204)
		w.Write([]byte("\n"))
		return
	}
	example, err := models.FindExample(id, rs.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println("Error finding example: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(&example)
	if e != nil {
		http.Error(w, "The example was found but we messed up with the return message", http.StatusInternalServerError)
		return
	}
}

// Update will update an example
func (rs Examples) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// go ahead and push it to the db and return success or error
	example := models.Example{}

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		log.Println("error parsing json request ", err)
		http.Error(w, "There was something wrong with the update message", http.StatusBadRequest)
		return
	}
	example.ID = bson.ObjectIdHex(id)

	err = example.Update(rs.Db)
	if err != nil {
		log.Println("error update example: ", err)
		http.Error(w, "The database crapped out when we tried to update the example", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
	w.Write([]byte("\n"))
}

// Delete will delete an example
func (rs Examples) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	example := models.Example{ID: bson.ObjectIdHex(id)}
	e := example.Delete(rs.Db)
	if e != nil {
		log.Println("error deleting example: ", e)
		http.Error(w, "We found the example but for some reason the delete crapped out", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
	w.Write([]byte("\n"))
}
