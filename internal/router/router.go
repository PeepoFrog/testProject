package router

import (
	"github.com/PeepoFrog/testProject/iternal/controller"
	"github.com/PeepoFrog/testProject/iternal/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	h := handler.NewHanlders()
	c := controller.NewControllerRepository(h)

	router.HandleFunc("/uploadfile", c.LoadFromCSVToDB).Methods("PUT")
	router.HandleFunc("/search", c.QrlSearch).Methods("GET")
	router.HandleFunc("/searchcsv", c.QrlSearchToCSV).Methods("GET")

	return router
}
