package data

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	return router
}
