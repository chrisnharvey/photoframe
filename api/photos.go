package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func servePhotos(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("../photos")

	if err != nil {
		log.Fatal(err)
	}

	var photos []string

	for _, f := range files {
		photos = append(photos, "http://"+r.Host+"/photos/"+f.Name())
	}

	data, _ := json.Marshal(photos)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
