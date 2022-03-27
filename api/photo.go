package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/disintegration/imaging"
)

func servePhoto(w http.ResponseWriter, r *http.Request) {
	file := path.Base(r.URL.Path)

	src, err := imaging.Open(os.Getenv("PHOTOS_PATH")+"/"+file, imaging.AutoOrientation(true))

	if err != nil {
		fmt.Fprintf(w, "Error loading photo: %d", err)

		return
	}

	width := 800
	height := 600

	query := r.URL.Query()

	queryWidth, hasWidth := query["w"]
	queryHeight, hasHeight := query["h"]

	if hasWidth {
		width, _ = strconv.Atoi(queryWidth[0])
	}

	if hasHeight {
		height, _ = strconv.Atoi(queryHeight[0])
	}

	img := imaging.Fit(src, width, height, imaging.Lanczos)

	w.Header().Set("Content-Type", "image/jpeg")
	imaging.Encode(w, img, imaging.JPEG)
}
