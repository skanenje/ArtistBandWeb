package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    artists, err := loadArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := struct {
        Artists []Artist
    }{
        Artists: artists,
    }

    templates.ExecuteTemplate(w, "index.html", data)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := loadArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "artists.html", artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artists, err := loadArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var artist Artist
	for _, a := range artists {
		if a.ID == id {
			artist = a
			break
		}
	}

	if artist.ID == 0 {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	templates.ExecuteTemplate(w, "artist.html", artist)
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := loadLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artists, err := loadArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a map of artist IDs to names
	artistNames := make(map[int]string)
	for _, artist := range artists {
		artistNames[artist.ID] = artist.Name
	}

	// Add artist names to locations
	for i, location := range locations.Index {
		locations.Index[i].Name = artistNames[location.ID]
	}

	templates.ExecuteTemplate(w, "locations.html", locations)
}

func datesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := loadDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artists, err := loadArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistNames := make(map[int]string)
	for _, artist := range artists {
		artistNames[artist.ID] = artist.Name
	}

	for i, date := range dates.Index {
		dates.Index[i].Name = artistNames[date.ID]
	}

	templates.ExecuteTemplate(w, "dates.html", dates)
}

func relationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := loadRelations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artists, err := loadArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artistNames := make(map[int]string)
	for _, artist := range artists {
		artistNames[artist.ID] = artist.Name
	}

	for i, relation := range relations.Index {
		relations.Index[i].Name = artistNames[relation.ID]
	}

	templates.ExecuteTemplate(w, "relations.html", relations)
}

func loadArtists() ([]Artist, error) {
	file, err := os.Open("artists.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var artists []Artist
	err = json.NewDecoder(file).Decode(&artists)
	return artists, err
}

func loadLocations() (LocationsData, error) {
	file, err := os.Open("locations.json")
	if err != nil {
		return LocationsData{}, err
	}
	defer file.Close()

	var locations LocationsData
	err = json.NewDecoder(file).Decode(&locations)
	return locations, err
}

func loadDates() (DatesData, error) {
	file, err := os.Open("dates.json")
	if err != nil {
		return DatesData{}, err
	}
	defer file.Close()

	var dates DatesData
	err = json.NewDecoder(file).Decode(&dates)
	return dates, err
}

func loadRelations() (RelationsData, error) {
	file, err := os.Open("relations.json")
	if err != nil {
		return RelationsData{}, err
	}
	defer file.Close()

	var relations RelationsData
	err = json.NewDecoder(file).Decode(&relations)
	return relations, err
}

func artistLocationsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/artist/locations/"))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := getArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := fetchLocations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find the locations for this specific artist
	artistLocations := []string{}
	for _, loc := range locations.Index {
		if loc.ID == artist.ID {
			artistLocations = loc.Locations
			break
		}
	}
	data := struct {
		Name      string
		Locations []string
	}{
		Name:      artist.Name,
		Locations: artistLocations,
	}

	templates.ExecuteTemplate(w, "artist_locations.html", data)
}

func artistRelationsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/artist/relations/"))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := getArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relations, err := fetchRelations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Find the relations for this specific artist
	var artistRelations map[string][]string
	for _, rel := range relations.Index {
		if rel.ID == artist.ID {
			artistRelations = rel.DatesLocations
			break
		}
	}
	data := struct {
		Name      string
		Relations map[string][]string
	}{
		Name:      artist.Name,
		Relations: artistRelations,
	}

	templates.ExecuteTemplate(w, "artist_relations.html", data)
}

func artistDatesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/artist/dates/"))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := getArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dates, err := fetchDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var artistDates []string
	for _, date := range dates.Index {
		if date.ID == artist.ID {
			artistDates = date.Dates
			break
		}
	}
	data := struct {
		Name  string
		Dates []string
	}{
		Name:  artist.Name,
		Dates: artistDates,
	}

	templates.ExecuteTemplate(w, "artist_dates.html", data)
}

func getArtistByID(id int) (Artist, error) {
	artists, err := loadArtists()
	if err != nil {
		return Artist{}, err
	}

	for _, artist := range artists {
		if artist.ID == id {
			return artist, nil
		}
	}

	return Artist{}, fmt.Errorf("artist not found")
}

// func SearchHandler(w http.ResponseWriter, r *http.Request) {
//     query := r.URL.Query().Get("query")
//     artists, err := loadArtists()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     var searchResults []Artist
//     for _, artist := range artists {
//         if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
//             searchResults = append(searchResults, artist)
//         }
//     }
//     data := map[string]interface{}{
//         "Query":   query,
//         "Artists": searchResults,
//     }
//     templates.ExecuteTemplate(w, "artists.html", data)
// }
