package main

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

func fetchArtists() ([]Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

func fetchLocations() (LocationsData, error) {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return LocationsData{}, err
	}
	defer resp.Body.Close()

	var locations LocationsData
	err = json.NewDecoder(resp.Body).Decode(&locations)
	return locations, err
}

func fetchDates() (DatesData, error) {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return DatesData{}, err
	}
	defer resp.Body.Close()

	var dates DatesData
	err = json.NewDecoder(resp.Body).Decode(&dates)
	return dates, err
}

func fetchRelations() (RelationsData, error) {
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return RelationsData{}, err
	}
	defer resp.Body.Close()

	var relations RelationsData
	err = json.NewDecoder(resp.Body).Decode(&relations)
	return relations, err
}