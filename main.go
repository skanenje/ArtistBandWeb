package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

func main() {
	// Parse templates
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Fetch and store data
	artists, err := fetchArtists()
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	saveJSON("artists.json", artists)

	locations, err := fetchLocations()
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	saveJSON("locations.json", locations)

	dates, err := fetchDates()
	if err != nil {
		log.Fatalf("Error fetching dates: %v", err)
	}
	saveJSON("dates.json", dates)

	relations, err := fetchRelations()
	if err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
	saveJSON("relations.json", relations)

	// Set up routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artists", artistsHandler)
	http.HandleFunc("/artist/", artistHandler)
	http.HandleFunc("/locations", locationsHandler)
	http.HandleFunc("/dates", datesHandler)
	http.HandleFunc("/relations", relationsHandler)

	// New routes for artist-specific pages
	http.HandleFunc("/artist/locations/", artistLocationsHandler)
	http.HandleFunc("/artist/relations/", artistRelationsHandler)
	http.HandleFunc("/artist/dates/", artistDatesHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func saveJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}