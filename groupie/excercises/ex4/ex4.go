package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type WrapLocation struct {
	Index []Location `json:"index"`
}

type Date struct {
	ID    int `json:"id"`
	Dates []string
}

type WrapDate struct {
	Index []Date `json:"index"`
}

type Relation struct {
	ID           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"`
}

type WrapRelation struct {
	Index []Relation `json:"index"`
}

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
	artist     []Artist
	location   []Location
	date       []Date
	relation   []Relation
)

func fetchAll() error {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	//urlRelations := "https://groupietrackers.herokuapp.com/api/relation"

	data, err := httpClient.Get(urlArtists)
	if err != nil {
		return fmt.Errorf("failed to fatch artists: %w", err)
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		if err := json.NewDecoder(data.Body).Decode(&artist); err != nil {
			return fmt.Errorf("failed to decode artists: %w", err)
		}
	}

	data, err = httpClient.Get(urlLocations)
	if err != nil {
		return fmt.Errorf("failed to fatch locations: %w", err)
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		var loct WrapLocation
		if err = json.NewDecoder(data.Body).Decode(&location); err != nil {
			return fmt.Errorf("failed to decode locations: %w", err)
		}
		location = loct.Index
	}

	data, err = httpClient.Get(urlDates)
	if err != nil {
		return fmt.Errorf("failed to fatch dates: %w", err)
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		var loct WrapLocation
		if err = json.NewDecoder(data.Body).Decode(&date); err != nil {
			return fmt.Errorf("failed to decode dates: %w", err)
		}
		location = loct.Index
	}

	data, err = httpClient.Get(urlLocations)
	if err != nil {
		return fmt.Errorf("failed to fatch locations: %w", err)
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		var loct WrapLocation
		if err = json.NewDecoder(data.Body).Decode(&location); err != nil {
			return fmt.Errorf("failed to decode locations: %w", err)
		}
		location = loct.Index
	}
	return nil
}

func main() {}
