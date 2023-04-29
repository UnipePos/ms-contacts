package main

import (
	"log"
	"net/http"

	config "github.com/UnipePos/ms-contacts/config"
	"github.com/UnipePos/ms-contacts/internal/app/contact"
	"github.com/UnipePos/ms-contacts/internal/app/http/rest"
	"github.com/UnipePos/ms-contacts/internal/app/storage/boltdb"
	"github.com/UnipePos/ms-contacts/internal/pkg/driver"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("cannot load .env file %s", err)
	}

	conn, err := driver.BoltConnect(cfg)
	if err != nil {
		log.Printf("an error has occurred with database %s", err)
	}

	defer conn.SQL.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	cr := boltdb.NewContactRepository(conn.SQL)
	cs := contact.NewContactService(cr)
	ch := rest.NewContactHandler(cs)

	router.Route("/contacts", func(r chi.Router) {
		r.Get("/", ch.GetContacts)
		r.Get("/{id}", ch.GetContact)
		r.Post("/", ch.PostContact)
		r.Delete("/{id}", ch.DeleteContact)
	})

	log.Println("Server Started on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
