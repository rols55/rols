package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// struct which contains artist data from Api
type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relatoins    string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type ParsedArtists struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
}

const apiUrl = "https://groupietrackers.herokuapp.com/api"

// inits a variable which hold an array of artist's data
var apiArtists []Artists

var artistLocations Locations

var locationCollection [52]Locations

var artistConcerts Dates

var concertCollection [52]Dates

var CleanedUpArtists []ParsedArtists

//TODO FOR SOME REASON append on line 73 and 76 replaces values of every element with new value
//instead of appending new value to a sclice

func parseData() bool {
	artist()
	parseLocations()
	parseConcerts()
	buildArtists(apiArtists, locationCollection, concertCollection)
	return true
}

func parseLocations() {
	ptr := &locationCollection
	ptr2 := &artistLocations
	ptr3 := &apiArtists
	for i := range *ptr3 {
		locations(i)
		cleanUpLocations(*ptr2)
		ptr[i] = *ptr2
		*ptr2 = Locations{
			Id:        0,
			Locations: []string{},
			Dates:     "",
		}
	}
}

func parseConcerts() {
	ptr := &apiArtists
	ptr2 := &artistConcerts
	ptr3 := &concertCollection
	for i := range *ptr {
		concerts(i)
		cleanUpConcerts(*ptr2)
		ptr3[i] = *ptr2
		*ptr2 = Dates{
			Id:    0,
			Dates: []string{},
		}
	}
}

func buildArtists(a []Artists, l [52]Locations, d [52]Dates) {
	for i := 0; i <= 51; i++ {
		temp := ParsedArtists{
			Id:           a[i].Id,
			Image:        a[i].Image,
			Name:         a[i].Name,
			Members:      a[i].Members,
			CreationDate: a[i].CreationDate,
			FirstAlbum:   a[i].FirstAlbum,
			Locations:    l[i].Locations,
			ConcertDates: d[i].Dates,
		}
		CleanedUpArtists = append(CleanedUpArtists, temp)
	}
}

func artist() error {
	//gets data from API
	resp, err := http.Get(apiUrl + "/artists")
	if err != nil {
		return err
	}

	//closes data after reading it
	defer resp.Body.Close()

	//reads data from API
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//saves data from API into an array of structs which contain artists data
	err = json.Unmarshal(bytes, &apiArtists)
	if err != nil {
		return err
	}
	return nil
}

func locations(id int) {
	resp, err := http.Get(apiArtists[id].Locations)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//UNMARSHALLING MUST BE DONE INTO A VARIABLE
	err = json.Unmarshal(bytes, &artistLocations)
	if err != nil {
		panic(err)
	}
}

func concerts(id int) Dates {
	resp, err := http.Get(apiArtists[id].ConcertDates)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//UNMARSHALLING MUST BE DONE INTO A VARIABLE
	err = json.Unmarshal(bytes, &artistConcerts)
	if err != nil {
		panic(err)
	}
	return artistConcerts
}

func cleanUpLocations(l Locations) Locations {
	for i, v := range l.Locations {
		//separate city from country
		a, b, _ := strings.Cut(v, "-")
		//if city name consists of more than one word we need to capitalise every word
		if strings.Contains(a, "_") {
			words := strings.Split(a, "_")
			a = ""
			for i := 0; i < len(words); i++ {
				words[i] = capLetter(words[i])
				a += words[i] + " "
			}
		} else {
			a = capLetter(a)
		}
		//if country name consists of two words we need to capitalise each words,
		//but check for three letter countries
		if len(b) == 2 || len(b) == 3 {
			b = strings.ToUpper(b)
		} else if strings.Contains(b, "_") {
			w1, w2, _ := strings.Cut(b, "_")
			w1 = capLetter(w1)
			w2 = capLetter(w2)
			b = w1 + " " + w2
		} else {
			b = capLetter(b)
		}
		v = a + "- " + b
		l.Locations[i] = v
	}
	return l
}

func cleanUpConcerts(d Dates) Dates {
	for i, v := range d.Dates {
		v = strings.Replace(v, "*", "", -1)
		d.Dates[i] = v
	}
	return d
}

func capLetter(s string) string {
	temp := []byte(s)
	temp[0] = temp[0] - 32
	s = string(temp)
	return s
}
