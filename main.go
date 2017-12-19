package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"gopkg.in/mgo.v2"

	"github.com/tutley/sk8/handlers"
)

func main() {

	serverPort := 3333
	// server port is a command line object so that you can reverse proxy with something
	// like NGINX and have a bunch of web apps running on one server
	var sPort string
	flag.StringVar(&sPort, "p", "3333", "The TCP Port for serving this webpage - default 3333")
	flag.Parse()

	sp, err := strconv.Atoi(sPort)
	if err != nil {
		log.Println("Error: port input at the command line failed to be applied: ", sPort)
	} else {
		serverPort = sp
	}

	// All of the other variables will be compiled into the app so set them here
	dbURL := "localhost"
	dbName := "sk8"

	// Init the Database
	session, err := mgo.Dial("mongodb://" + dbURL)
	if err != nil {
		log.Fatal("DB Connect error: ", err)
	}

	defer session.Close()
	db := session.DB(dbName)

	// Setup JWT strategy
	var tokenAuth *jwtauth.JwtAuth
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	r := chi.NewRouter()
	// Use Chi built-in middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultCompress)
	// Setup routes
	r.Route("/api", func(r chi.Router) {

		// TODO: Narrow this down so not everyone in the world can make API requests
		cors := cors.New(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(cors.Handler)
		r.Mount("/examples", handlers.Examples{
			Db: db,
			Ta: tokenAuth}.Routes())
		r.Mount("/auth", handlers.Auth{
			Db: db,
			Ta: tokenAuth}.Routes())
	})

	// This serves the static files
	box := rice.MustFindBox("sk8-dist")

	distFileServer := http.StripPrefix("/", http.FileServer(box.HTTPBox()))
	r.Mount("/", distFileServer)

	serveAddr := ":" + strconv.Itoa(serverPort)
	log.Println("sk8 Server listening on: ", strconv.Itoa(serverPort))
	http.ListenAndServe(serveAddr, r)
}
