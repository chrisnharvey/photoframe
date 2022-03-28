package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chrisnharvey/photoframe/api"
)

func main() {
	s, err := api.GetSettings()

	if err != nil {
		fmt.Printf("%v", err)
		fmt.Println()
		os.Exit(1)
	}

	api := api.Api{
		Settings: s,
	}

	http.HandleFunc("/api/photos", api.ServePhotos)
	http.HandleFunc("/api/settings", api.ServeSettings)
	http.HandleFunc("/photos/", api.ServePhoto)

	fs := http.FileServer(http.Dir(s.UIPath))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Printf("Starting server on %v", s.Listen)
	fmt.Println()

	err = http.ListenAndServe(s.Listen, nil)

	if err != nil {
		fmt.Printf("%v", err)
		fmt.Println()
		os.Exit(2)
	}
}
