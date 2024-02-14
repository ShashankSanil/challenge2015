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
					output := map[string]interface{}{"Movie": movie["url"], i["role"].(string): res[0]["url"]}
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
			output := map[string]interface{}{"Movie": movie["url"], movie["role"].(string): actorFrom}
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
				output := map[string]interface{}{"Movie": actorFrom, v["role"].(string): actorTo}
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
				output := map[string]interface{}{"Movie": actorFrom, v["role"].(string): actorTo}
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
			output := map[string]interface{}{"Movie": actorFrom, v.(map[string]interface{})["role"].(string): v.(map[string]interface{})["url"]}
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
	_, degree, _, _ := separation(fr, to, treeDepth, 0, []map[string]interface{}{}, checkList, 0)
	fmt.Printf("Degree of Separation - %d\n", degree)
}
