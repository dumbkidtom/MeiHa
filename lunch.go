package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	// "encoding/json"
	// "regexp"
	// "net/url"
)

// controller REST API URL
var controller_url string = "http://localhost:5000/restaurants"

type Results struct {
	Businesses []Business `json:"businesses"`
	Total      int        `json:"total"`
}

// business properties
type Business struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	IMGURL string `json:"image_url"`
	Phone  string `json:"phone"`
}

func index_handler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "input.gtpl")

}

func img_handler(w http.ResponseWriter, r *http.Request) {
	// myimg := r.URL.Path
	http.ServeFile(w, r, "images/gps.png")
}

func search_handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	myzip := r.Form.Get("zip")
	myradius := r.Form.Get("radius")
	mylatitude := r.Form.Get("lat")
	mylongitude := r.Form.Get("long")
	url := ""

	// call Controller REST API for results
	if len(myzip) > 5 {
		url = fmt.Sprintf("%s?zip=0&latitude=%s&longitude=%s&radius=%s", controller_url, mylatitude, mylongitude, myradius)
	} else {
		url = fmt.Sprintf("%s?zip=%s&radius=%s", controller_url, myzip, myradius)
	}

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// convert []byte(body) to json
	res := &Results{}
	err := json.Unmarshal([]byte(body), res)
	if err != nil {
		log.Fatal(err)
	}

	// generate random number within list of restaurants
	tt := len(res.Businesses)
	rn := rand.Intn(tt)

	// display restaurant info using template
	t, _ := template.ParseFiles("card.gtpl")
	t.Execute(w, res.Businesses[rn])

	//fmt.Fprintf(w, "total: %d, rand: %d", tt, rn)

}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/images/", img_handler)
	http.HandleFunc("/search", search_handler)
	http.ListenAndServe(":8080", nil)

}
