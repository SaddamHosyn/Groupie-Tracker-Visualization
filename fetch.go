package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	fetchError bool
)

const (
	apiURL            = "https://groupietrackers.herokuapp.com/api"
	artistsEndpoint   = apiURL + "/artists"
	locationsEndpoint = apiURL + "/locations"
	datesEndpoint     = apiURL + "/dates"
	relationsEndpoint = apiURL + "/relation"
)

// Artist data structure remains the same
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Artist data structure remains the same
type ArtistPageData struct {
	ID            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	CreationDate  int      `json:"creationDate"`
	FirstAlbum    string   `json:"firstAlbum"`
	Locations     []string
	Dates         []string
	Relations     []string
	LocationCount int
}

type LocationAPIResponse struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesAPIResponse struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

/* type Relation struct {
	Index []struct {
		ID        int      `json:"id"`
		Dates     []string `json:"dates"`
		Locations []string `json:"locations"`
	} `json:"index"`
} */

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ErrorResponse struct {
	Code    int
	Message string
}

var (
	artists   []Artist
	locations []Location
	dates     []Dates
	relations Relation
)

func ToUpper(s string) string {
	str := strings.Fields(s)
	for i := range str {
		if strings.ToLower(str[i]) == "usa" || strings.ToLower(str[i]) == "uk" {
			str[i] = strings.ToUpper(str[i])
		} else {
			str[i] = strings.ToUpper(string(str[i][0])) + string(str[i][1:])
		}
	}
	return strings.Join(str, " ")
}

func FetchArtists() error {
	resp, err := http.Get(artistsEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &artists); err != nil {
		return err
	}
	log.Println("Artists fetched successfully")
	return nil
}

func FetchLocations() error {
	resp, err := http.Get(locationsEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var apiResponse LocationAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}
	locations = apiResponse.Index
	log.Println("Locations fetched successfully")
	return nil
}

func FetchDates() error {
	resp, err := http.Get(datesEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Use the wrapper struct to unmarshal the data
	var apiResponse DatesAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	// Assign the fetched dates
	dates = apiResponse.Index
	log.Println("Dates fetched successfully")
	return nil
}

func FetchRelations() error {
	resp, err := http.Get(relationsEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &relations); err != nil {
		return err
	}
	log.Println("Relations fetched successfully")
	return nil
}

// Fetch All Data at Once
func FetchAllDataError() {

	if err := FetchArtists(); err != nil {
		fetchError = true
	}
	if err := FetchLocations(); err != nil {
		fetchError = true

	}
	if err := FetchDates(); err != nil {
		fetchError = true

	}
	if err := FetchRelations(); err != nil {
		fetchError = true

	}
}

func FetchArtistData(id int) (Artist, []string, []string, map[string][]string, error) {
	// Fetch the artist by ID
	var selectedArtist Artist
	for _, artistvar := range artists {
		if artistvar.ID == id {
			selectedArtist = artistvar
			break
		}
	}

	var associatedLocations []string
	var associatedDates []string
	var associatedRelations map[string][]string

	for _, location := range locations {
		if location.ID == id {
			for _, loc := range location.Locations {
				cleanLoc := strings.ReplaceAll(loc, "_", " ")
				cleanLoc = strings.ReplaceAll(cleanLoc, "-", " ")
				cleanLoc = ToUpper(cleanLoc)
				associatedLocations = append(associatedLocations, cleanLoc)
			}

		}
	}

	for _, date := range dates {
		if date.ID == id {
			for _, d := range date.Dates {
				cleanDate := strings.ReplaceAll(d, "*", "")
				associatedDates = append(associatedDates, cleanDate)
			}
		}
	}

	for _, relation := range relations.Index {
		if relation.ID == id {
			associatedRelations = make(map[string][]string)
			for loc, relDates := range relation.DatesLocations {
				cleanLoc := strings.ReplaceAll(loc, "_", " ")
				cleanLoc = strings.ReplaceAll(cleanLoc, "-", " ")
				cleanLoc = ToUpper(cleanLoc)
				cleanRelDates := []string{}
				for _, d := range relDates {
					cleanDate := strings.ReplaceAll(d, "*", "")
					cleanRelDates = append(cleanRelDates, cleanDate)
				}
				associatedRelations[cleanLoc] = cleanRelDates
			}
		}
	}
	return selectedArtist, associatedLocations, associatedDates, associatedRelations, nil
}
