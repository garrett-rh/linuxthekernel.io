// Package handlers contains API http handlers
package main

import (
	"encoding/json"
	"linuxthekernel.io/internal"
	"net/http"
)

type locality struct {
	Name string `json:"name"`
}

var localityMapping = map[string]func(*internal.CarInfo){
	"Arlington County": (*internal.CarInfo).ArlingtonTaxCalculator,
	"Fairfax County":   (*internal.CarInfo).FairfaxTaxCalculator,
	"Alexandria City":  (*internal.CarInfo).AlexandriaTaxCalculator,
}

// PostsHandler returns the listing of all posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	posts, err := internal.GetAllPosts()
	internal.SortPostsByDate(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PostHandler handler for fetching a single post by ID
func PostHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	content, err := internal.GetPostContent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LocalitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	supportedLocalities := []locality{
		{Name: "Arlington County"},
		{Name: "Fairfax County"},
		{Name: "Alexandria City"},
	}
	err := json.NewEncoder(w).Encode(supportedLocalities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CarTaxHandler(w http.ResponseWriter, r *http.Request) {
	var car internal.CarInfo
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if fn := localityMapping[car.Locality]; fn != nil {
		fn(&car)
	}
	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
