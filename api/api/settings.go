package api

import (
	"encoding/json"
	"net/http"

	"github.com/caarlos0/env/v6"
)

type Settings struct {
	Listen      string `env:"LISTEN,required,notEmpty" envDefault:":80"`
	PhotosPath  string `env:"PHOTOS_PATH,required,notEmpty"`
	UIPath      string `env:"UI_PATH,required,notEmpty"`
	RefreshTime int    `env:"REFRESH_TIME,required,notEmpty" envDefault:"30000"`
	PhotoTime   int    `env:"PHOTO_TIME,required,notEmpty" envDefault:"10000"`
}

func GetSettings() (*Settings, error) {
	s := Settings{}

	err := env.Parse(&s)

	return &s, err
}

func (api *Api) ServeSettings(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(api.Settings)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
