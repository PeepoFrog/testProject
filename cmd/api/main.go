package main

import (
	"log"
	"net/http"

	"github.com/PeepoFrog/testProject/internal/router"
)

func main() {
	//handler.LoadFromCSVTest()

	r := router.Router()
	srv := &http.Server{
		Addr:    "0.0.0.0:4000",
		Handler: r, //
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
