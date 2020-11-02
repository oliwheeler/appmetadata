package main

import (
	"log"
	"net/http"

	"github.com/oliwheeler/appmetadata/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", api.New()))
}
