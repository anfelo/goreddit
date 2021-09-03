package main

import (
	"log"
	"net/http"

	"github.com/anfelo/goreddit/postgres"
	"github.com/anfelo/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("012345678901234567890123456789")
	h := web.NewHandler(store, csrfKey)
	http.ListenAndServe(":3000", h)
}
