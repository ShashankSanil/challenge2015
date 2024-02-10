// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// type Involvers struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// 	Role string `json:"role"`
// }

// type Movie struct {
// 	Name string       `json:"name"`
// 	Url  string       `json:"url"`
// 	Role string       `json:"role"`
// 	Cast []*Involvers `json:"cast"`
// 	Crew []*Involvers `json:"crew"`
// }

// type Person struct {
// 	Name   string   `json:"name"`
// 	Url    string   `json:"url"`
// 	Movies []*Movie `json:"movies"`
// }

// const moviebuffURL = "http://data.moviebuff.com/"

// var result []map[string]string

// func getActorURL(actorName string) string {
// 	return strings.ReplaceAll(strings.ToLower(actorName), " ", "-")
// }

// func fetchData(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request to %s failed with status code: %d", url, resp.StatusCode)
// 	}

// 	contentType := resp.Header.Get("Content-Type")
// 	if !strings.Contains(contentType, "application/json") {
// 		return nil, fmt.Errorf("API returned unexpected content type: %s for URL: %s", contentType, url)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

// func getActorData(actorURL string) (*Person, error) {
// 	var actorData Person
// 	res, err := fetchData(moviebuffURL + actorURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &actorData)
// 	return &actorData, nil
// }

// func getMovieData(movieURL string) (*Movie, error) {
// 	var movieData Movie
// 	res, err := fetchData(moviebuffURL + movieURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &movieData)
// 	return &movieData, nil
// }

// var actorFailure bool = false

// func findDegreesOfSeparation(actor1, actor2 string, viewed, viewedMovie map[string]bool) {
// 	fmt.Println("------------", actor1, actor2)
// 	actor1Data, err := getActorData(actor1)
// 	if err != nil {
// 		fmt.Println("Error fetching actor1 data:", err)
// 		actorFailure = true
// 		return
// 	}

// 	for _, movieEntry := range actor1Data.Movies {
// 		movieData, err := getMovieData(movieEntry.Url)
// 		if err != nil {
// 			fmt.Println("Error fetching movie data:", err)
// 			continue
// 		}
// 		total := make([]*Involvers, 0)
// 		total = append(total, movieData.Cast...)
// 		total = append(total, movieData.Crew...)
// 		if item, ok := contains(actor2, total); ok {
// 			res := make(map[string]string)
// 			res["Movie"] = movieEntry.Name
// 			res[movieEntry.Role] = actor1Data.Name
// 			res[item.Role] = item.Name
// 			result = append(result, res)
// 			return
// 		}
// 	}
// 	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
// 	for _, movieEntry := range actor1Data.Movies {
// 		movieData, err := getMovieData(movieEntry.Url)
// 		if err != nil {
// 			fmt.Println("Error fetching movie data:", err)
// 			continue
// 		}
// 		total := make([]*Involvers, 0)
// 		total = append(total, movieData.Cast...)
// 		total = append(total, movieData.Crew...)
// 		for _, psn := range total {
// 			if psn.Url == actor2 {
// 				fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@", psn.Url)
// 				res := make(map[string]string)
// 				res["Movie"] = movieEntry.Name
// 				res[movieEntry.Role] = actor1Data.Name
// 				res[psn.Role] = psn.Name
// 				result = append(result, res)
// 				return
// 			}
// 			if !viewed[psn.Url] {
// 				viewed[psn.Url] = true
// 				res := make(map[string]string)
// 				res["Movie"] = movieEntry.Name
// 				res[movieEntry.Role] = actor1Data.Name
// 				res[psn.Role] = psn.Name
// 				result = append(result, res)
// 				findDegreesOfSeparation(psn.Url, actor2, viewed, viewedMovie)
// 				if !actorFailure {
// 					return
// 				}
// 				actorFailure = false
// 			}
// 		}
// 	}
// }

// func contains(target string, slice []*Involvers) (*Involvers, bool) {
// 	for _, item := range slice {
// 		// fmt.Println("item.url", item.Url)
// 		if item.Url == target {
// 			return item, true
// 		}
// 	}
// 	return nil, false
// }

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: go run degrees_of_separation.go <actor1_name> <actor2_name>")
// 		return
// 	}

// 	actor1Name := os.Args[1]
// 	actor2Name := os.Args[2]

// 	actor1URL := getActorURL(actor1Name)
// 	actor2URL := getActorURL(actor2Name)
// 	viewed := make(map[string]bool, 0)
// 	viewedMovie := make(map[string]bool, 0)
// 	viewed[actor1URL] = true
// 	findDegreesOfSeparation(actor1URL, actor2URL, viewed, viewedMovie)
// 	fmt.Println("Degrees of Separation: ", len(result))
// 	fmt.Println(result)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// type Involvers struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// 	Role string `json:"role"`
// }

// type Movie struct {
// 	Name string       `json:"name"`
// 	Url  string       `json:"url"`
// 	Role string       `json:"role"`
// 	Cast []*Involvers `json:"cast"`
// 	Crew []*Involvers `json:"crew"`
// }

// type Movie1 struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// 	Role string `json:"role"`
// }
// type Person struct {
// 	Name   string   `json:"name"`
// 	Url    string   `json:"url"`
// 	Movies []*Movie `json:"movies"`
// }

// const moviebuffURL = "http://data.moviebuff.com/"

// var result []map[string]string

// var degreeOfSeperation int
// var minDegree int

// func getActorURL(actorName string) string {
// 	return strings.ReplaceAll(strings.ToLower(actorName), " ", "-")
// }

// func fetchData(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request to %s failed with status code: %d", url, resp.StatusCode)
// 	}

// 	contentType := resp.Header.Get("Content-Type")
// 	if !strings.Contains(contentType, "application/json") {
// 		return nil, fmt.Errorf("API returned unexpected content type: %s for URL: %s", contentType, url)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

// func getActorData(actorURL string) (*Person, error) {
// 	var actorData Person
// 	res, err := fetchData(moviebuffURL + actorURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &actorData)
// 	return &actorData, nil
// }

// func getMovieData(movieURL string) (*Movie, error) {
// 	var movieData Movie
// 	res, err := fetchData(moviebuffURL + movieURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &movieData)
// 	return &movieData, nil
// }

// func findDegreesOfSeparation(actor1, actor2 string, viewed, viewedMovie map[string]bool) {
// 	actor1Data, err := getActorData(actor1)
// 	if err != nil {
// 		fmt.Println("Error fetching actor1 data:", err)
// 		return
// 	}

// 	for _, movieEntry := range actor1Data.Movies {
// 		movieData, err := getMovieData(movieEntry.Url)
// 		if err != nil {
// 			fmt.Println("Error fetching movie data:", err)
// 			continue
// 		}
// 		total := make([]*Involvers, 0)
// 		total = append(total, movieData.Cast...)
// 		total = append(total, movieData.Crew...)
// 		for _, psn := range total {
// 			if psn.Url == actor2 {
// 				if degreeOfSeperation < minDegree {
// 					minDegree = degreeOfSeperation
// 				}
// 				break
// 			} else {
// 				degreeOfSeperation++
// 				if degreeOfSeperation > minDegree {
// 					break
// 				}
// 				findDegreesOfSeparation(psn.Url, actor2, viewed, viewedMovie)
// 				return
// 			}
// 		}

// 	}

// }

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: go run degrees_of_separation.go <actor1_name> <actor2_name>")
// 		return
// 	}

// 	actor1Name := os.Args[1]
// 	actor2Name := os.Args[2]

// 	actor1URL := getActorURL(actor1Name)
// 	actor2URL := getActorURL(actor2Name)
// 	viewed := make(map[string]bool, 0)
// 	viewedMovie := make(map[string]bool, 0)
// 	viewed[actor1URL] = true
// 	findDegreesOfSeparation(actor1URL, actor2URL, viewed, viewedMovie)
// 	fmt.Println("Degrees of Separation: ", minDegree)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// type Involvers struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// 	Role string `json:"role"`
// }

// type Movie struct {
// 	Name string       `json:"name"`
// 	Url  string       `json:"url"`
// 	Role string       `json:"role"`
// 	Cast []*Involvers `json:"cast"`
// 	Crew []*Involvers `json:"crew"`
// }

// type Person struct {
// 	Name   string   `json:"name"`
// 	Url    string   `json:"url"`
// 	Movies []*Movie `json:"movies"`
// }

// const moviebuffURL = "http://data.moviebuff.com/"

// var minDegree int = -1

// func getActorURL(actorName string) string {
// 	return strings.ReplaceAll(strings.ToLower(actorName), " ", "-")
// }

// func fetchData(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("API request to %s failed with status code: %d", url, resp.StatusCode)
// 	}

// 	contentType := resp.Header.Get("Content-Type")
// 	if !strings.Contains(contentType, "application/json") {
// 		return nil, fmt.Errorf("API returned unexpected content type: %s for URL: %s", contentType, url)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

// func getActorData(actorURL string) (*Person, error) {
// 	var actorData Person
// 	res, err := fetchData(moviebuffURL + actorURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &actorData)
// 	return &actorData, nil
// }

// func getMovieData(movieURL string) (*Movie, error) {
// 	var movieData Movie
// 	res, err := fetchData(moviebuffURL + movieURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	json.Unmarshal(res, &movieData)
// 	return &movieData, nil
// }

// func findDegreesOfSeparation(actor1, actor2 string, degree int, viewed map[string]bool) {
// 	if actor1 == actor2 {
// 		if minDegree == -1 || degree < minDegree {
// 			minDegree = degree
// 		}
// 		return
// 	}

// 	if minDegree != -1 && degree >= minDegree {
// 		return
// 	}

// 	if viewed[actor1] {
// 		return
// 	}
// 	viewed[actor1] = true

// 	actor1Data, err := getActorData(actor1)
// 	if err != nil {
// 		fmt.Println("Error fetching actor data:", err)
// 		return
// 	}

// 	for _, movieEntry := range actor1Data.Movies {
// 		movieData, err := getMovieData(movieEntry.Url)
// 		if err != nil {
// 			fmt.Println("Error fetching movie data:", err)
// 			continue
// 		}

// 		for _, person := range append(movieData.Cast, movieData.Crew...) {
// 			findDegreesOfSeparation(person.Url, actor2, degree+1, viewed)
// 		}
// 	}
// }

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: go run degrees_of_separation.go <actor1_name> <actor2_name>")
// 		return
// 	}

// 	actor1Name := os.Args[1]
// 	actor2Name := os.Args[2]

// 	actor1URL := getActorURL(actor1Name)
// 	actor2URL := getActorURL(actor2Name)
// 	viewed := make(map[string]bool, 0)
// 	findDegreesOfSeparation(actor1URL, actor2URL, 0, viewed)

// 	if minDegree == -1 {
// 		fmt.Println("No connection found between the actors")
// 	} else {
// 		fmt.Println("Degrees of Separation:", minDegree)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var checkList []map[string]interface{}

type statusControl struct {
	StatusCode int
}

var counter int

func separation(actorFrom string, actorTo string, degree int, count int, history []map[string]interface{}, checkList []map[string]interface{}, mainCheck int) (bool, int, []map[string]interface{}, []map[string]interface{}) {
	if count > degree {
		return false, count, history, checkList
	}

	var check []map[string]interface{}
	for _, x := range checkList {
		if x["url"] == actorFrom {
			check = append(check, x)
		}
	}

	if len(check) > 0 {
		return false, count, history, checkList
	}

	counter++

	r1, err := http.Get("http://data.moviebuff.com/" + actorFrom)
	if err != nil || r1.StatusCode != http.StatusOK {
		return false, count, history, checkList
	}

	var data map[string]interface{}
	body, _ := ioutil.ReadAll(r1.Body)
	json.Unmarshal(body, &data)

	check = nil
	for _, x := range checkList {
		if x["url"] == actorTo {
			check = append(check, x)
		}
	}

	var r2 *http.Response
	var data1 map[string]interface{}

	if len(check) > 0 {
		data1 = check[0]["data"].(map[string]interface{})
		r2 = &http.Response{StatusCode: 200}
	} else {
		r2, err = http.Get("http://data.moviebuff.com/" + actorTo)
		if err != nil || r2.StatusCode != http.StatusOK {
			return false, count, history, checkList
		}

		body, _ := ioutil.ReadAll(r2.Body)
		json.Unmarshal(body, &data1)

		checkList = append(checkList, map[string]interface{}{"url": actorTo, "data": data1})
	}

	if r1.StatusCode != http.StatusOK || r2.StatusCode != http.StatusOK {
		return false, count, history, checkList
	}

	if data["type"].(string) == "Person" && data1["type"].(string) == "Person" && mainCheck == 0 {
		mainCheck++
		movies := data["movies"].([]interface{})
		movieMap := make(map[string]interface{})
		for _, v := range movies {
			movie := v.(map[string]interface{})
			movieMap[movie["url"].(string)] = movie
		}

		for _, v := range data1["movies"].([]interface{}) {
			movie := v.(map[string]interface{})
			if _, ok := movieMap[movie["url"].(string)]; ok {
				r2, _ := http.Get("http://data.moviebuff.com/" + movie["url"].(string))
				body, _ := ioutil.ReadAll(r2.Body)
				json.Unmarshal(body, &data)
				merge := append(data["crew"].([]interface{}), data["cast"].([]interface{})...)
				mergeMap := make(map[string]interface{})
				for _, v := range merge {
					member := v.(map[string]interface{})
					mergeMap[member["url"].(string)] = member
				}
				res := []map[string]interface{}{}
				for _, v := range mergeMap {
					res = append(res, v.(map[string]interface{}))
				}
				for _, i := range res {
					output := map[string]interface{}{"Movie": movie["url"], i["url"].(string): res[0]["role"], "Type": "P"}
					history = append(history, output)
				}
				return true, count, history, checkList
			}
		}
		if len(data["movies"].([]interface{})) > len(data1["movies"].([]interface{})) {
			checkList = append(checkList, map[string]interface{}{"url": actorTo, "data": data1})
			data = data1
			actorFrom, actorTo = actorTo, actorFrom
		}
	}

	if data["type"].(string) == "Person" {
		movies := data["movies"].([]interface{})
		for _, movieInterface := range movies {
			movie := movieInterface.(map[string]interface{})
			count++
			output := map[string]interface{}{"Movie": movie["url"], actorFrom: movie["role"], "Type": "P"}
			history = append(history, output)
			a, b, c, d := separation(movie["url"].(string), actorTo, degree, count, history, checkList, mainCheck)
			checkList = d
			if a {
				return a, b, c, checkList
			} else {
				history = history[:len(history)-1]
				count--
			}
		}
	} else if data["type"].(string) == "Movie" {
		crew := []map[string]interface{}{}
		for _, v := range data["crew"].([]interface{}) {
			member := v.(map[string]interface{})
			crew = append(crew, member)
		}
		for _, v := range crew {
			if v["url"].(string) == actorTo {
				output := map[string]interface{}{"Movie": actorFrom, actorTo: v["role"], "Type": "M"}
				history = append(history, output)
				return true, count, history, checkList
			}
		}

		cast := []map[string]interface{}{}
		for _, v := range data["cast"].([]interface{}) {
			member := v.(map[string]interface{})
			cast = append(cast, member)
		}
		for _, v := range cast {
			if v["url"].(string) == actorTo {
				output := map[string]interface{}{"Movie": actorFrom, actorTo: v["role"], "Type": "M"}
				history = append(history, output)
				return true, count, history, checkList
			}
		}

		merge := append(data["crew"].([]interface{}), data["cast"].([]interface{})...)
		mergeMap := make(map[string]interface{})
		for _, v := range merge {
			member := v.(map[string]interface{})
			mergeMap[member["url"].(string)] = member
		}
		for _, v := range mergeMap {
			count++
			output := map[string]interface{}{"Movie": actorFrom, v.(map[string]interface{})["url"].(string): v.(map[string]interface{})["role"], "Type": "M"}
			history = append(history, output)
			a, b, c, d := separation(v.(map[string]interface{})["url"].(string), actorTo, degree, count, history, checkList, mainCheck)
			checkList = d
			if a {
				return a, b, c, checkList
			} else {
				history = history[:len(history)-1]
				count--
			}
		}
	}

	return false, count, history, checkList
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <actorFrom> <actorTo>")
		return
	}

	fr := args[0]
	to := args[1]

	treeDepth := 3
	result, degree, connections, _ := separation(fr, to, treeDepth, 0, []map[string]interface{}{}, checkList, 0)
	fmt.Printf("Degree of Separation - %d\n", degree)
	if result {
		for _, connection := range connections {
			fmt.Println(connection)
		}
	} else {
		fmt.Println("not found")
	}
}
