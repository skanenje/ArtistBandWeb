// models.go

package main

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsData struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"` 
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesData struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Name      string   `json:"name"` 
	Dates []string `json:"dates"`
}

type RelationsData struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	Name      string   `json:"name"` 
	DatesLocations map[string][]string `json:"datesLocations"`
}