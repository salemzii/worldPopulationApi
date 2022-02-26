package main

import (
	"net/http"

	"github.com/salemzii/worldPopulationApi/src/data"
)

func main() {
	router := data.Router()

	http.Handle("/", router)

	router.HandleFunc("/api/all", data.Allpopulation)
	router.HandleFunc("/api/{name}", data.GetCountryPopulationByName)
	router.HandleFunc("/api/top/pop", data.TopFiveMostPopulated)
	router.HandleFunc("/api/buttom/pop", data.ButtomFiveMostPopulated)

	http.ListenAndServe(":8080", router)
}
