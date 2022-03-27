package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/photos", servePhotos)
	http.HandleFunc("/photos/", servePhoto)

	fs := http.FileServer(http.Dir(os.Getenv("UI_PATH")))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(os.Getenv("LISTEN"), nil)
}
