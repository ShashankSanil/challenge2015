package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Involvers struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Role string `json:"role"`
}

type Movie struct {
	Name string       `json:"name"`
	Url  string       `json:"url"`
	Role string       `json:"role"`
	Cast []*Involvers `json:"cast"`
	Crew []*Involvers `json:"crew"`
}

type Person struct {
	Name   string   `json:"name"`
	Url    string   `json:"url"`
	Movies []*Movie `json:"movies"`
}

const moviebuffURL = "http://data.moviebuff.com/"

var minDegree int = -1

func getActorURL(actorName string) string {
	return strings.ReplaceAll(strings.ToLower(actorName), " ", "-")
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request to %s failed with status code: %d", url, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return nil, fmt.Errorf("API returned unexpected content type: %s for URL: %s", contentType, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getActorData(actorURL string) (*Person, error) {
	var actorData Person
	res, err := fetchData(moviebuffURL + actorURL)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &actorData)
	return &actorData, nil
}

func getMovieData(movieURL string) (*Movie, error) {
	var movieData Movie
	res, err := fetchData(moviebuffURL + movieURL)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &movieData)
	return &movieData, nil
}

func findDegreesOfSeparation(actor1, actor2 string, degree int, viewed map[string]bool) {
	if actor1 == actor2 {
		if minDegree == -1 || degree < minDegree {
			minDegree = degree
		}
		return
	}

	if minDegree != -1 && degree >= minDegree {
		return
	}

	if viewed[actor1] {
		return
	}
	viewed[actor1] = true

	actor1Data, err := getActorData(actor1)
	if err != nil {
		fmt.Println("Error fetching actor data:", err)
		return
	}

	for _, movieEntry := range actor1Data.Movies {
		movieData, err := getMovieData(movieEntry.Url)
		if err != nil {
			fmt.Println("Error fetching movie data:", err)
			continue
		}

		for _, person := range append(movieData.Cast, movieData.Crew...) {
			findDegreesOfSeparation(person.Url, actor2, degree+1, viewed)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run degrees_of_separation.go <actor1_name> <actor2_name>")
		return
	}

	actor1Name := os.Args[1]
	actor2Name := os.Args[2]

	actor1URL := getActorURL(actor1Name)
	actor2URL := getActorURL(actor2Name)
	viewed := make(map[string]bool, 0)
	findDegreesOfSeparation(actor1URL, actor2URL, 0, viewed)

	if minDegree == -1 {
		fmt.Println("No connection found between the actors")
	} else {
		fmt.Println("Degrees of Separation:", minDegree)
	}
}
