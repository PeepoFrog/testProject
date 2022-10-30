package router

import (
	"github.com/PeepoFrog/testProject/internal/controller"
	postgre "github.com/PeepoFrog/testProject/internal/database/postgres"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	h := postgre.NewPostgre()
	c := controller.NewController(h)

	router.HandleFunc("/uploadfile", c.LoadFromCSVToDB).Methods("PUT")
	router.HandleFunc("/search", c.QrlSearch).Methods("GET")
	router.HandleFunc("/searchcsv", c.QrlSearchToCSV).Methods("GET")

	return router
}
