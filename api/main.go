package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/photos", servePhotos)

	pics := http.FileServer(http.Dir(os.Getenv("PHOTOS_PATH")))
	http.Handle("/photos/", http.StripPrefix("/photos/", pics))

	fs := http.FileServer(http.Dir(os.Getenv("UI_PATH")))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(os.Getenv("LISTEN"), nil)
}
