package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type PopulationModel struct {
	Rank       int
	Name       string
	Pop2022    string
	Pop2021    string
	Growthrate string
	Area       int
	Density    string
}

func ShowWorldPopulation() []PopulationModel {

	data, err := ioutil.ReadFile("/home/salem/Desktop/worlddata/src/data/data.json")
	if err != nil {
		fmt.Println(err)
	}

	var dataStore []PopulationModel

	json.Unmarshal([]byte(data), &dataStore)
	return dataStore
}

func Allpopulation(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(ShowWorldPopulation())
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(data)
}

func GetCountryPopulationByName(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	countryName := param["name"]
	w.Header().Set("Content-Type", "application/json")
	for _, v := range ShowWorldPopulation() {

		if v.Name == countryName {
			data, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
			} else {
				w.Write(data)
			}
		}

	}

	fmt.Println(countryName)
}

func TopFiveMostPopulated(w http.ResponseWriter, r *http.Request) {

	ls := ShowWorldPopulation()[:6]
	fmt.Println(ls)

	data, err := json.Marshal(ls)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ButtomFiveMostPopulated(w http.ResponseWriter, r *http.Request) {

	ls := ShowWorldPopulation()[len(ShowWorldPopulation())-6 : len(ShowWorldPopulation())-1]
	fmt.Println(ls)

	data, err := json.Marshal(ls)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
