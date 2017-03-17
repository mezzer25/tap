package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func BeerIndex(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(beers); err != nil {
		panic(err)
	}
}

func BeerShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var beerId int
	var err error
	if beerId, err = strconv.Atoi(vars["beerId"]); err != nil {
		panic(err)
	}
	beer := RepoFindBeer(beerId)
	if beer.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(beer); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
