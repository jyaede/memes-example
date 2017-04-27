package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/yaede7/memes-example/imgflip"
	"gopkg.in/mgo.v2"
)

func main() {
	godotenv.Load()

	imgflipClient := imgflip.New(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	sess, err := mgo.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mongodb")

	//Define routes
	r := httprouter.New()

	mh := MemeHandler{imgflipClient, sess}
	r.GET("/memes/available", mh.Available)
	r.GET("/memes/collection", mh.Collection)
	r.POST("/memes/collection", mh.Create)
	r.GET("/memes/collection/:id", mh.Get)
	r.DELETE("/memes/collection/:id", mh.Delete)

	//log every http request
	h := NewLogger()(r)

	//keep service alive if a panic happens
	h = handlers.RecoveryHandler()(h)

	//Set cors
	h = cors.New(cors.Options{
		MaxAge:           600,
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization", "Accept-Language"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PATCH", "PUT", "DELETE"},
	}).Handler(h)

	log.Println("Staring http server")
	log.Fatal(http.ListenAndServe(":8000", h))
}
