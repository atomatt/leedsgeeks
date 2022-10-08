package main // import "github.com/emgee/leedsgeeks"

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
)

var (
	//go:embed leedsgeeks.json
	//go:embed static/*
	//go:embed templates/*
	embedFS embed.FS

	configFile = "leedsgeeks.json"
	templates  = template.Must(template.ParseFS(embedFS, "templates/*.html"))
)

type Config struct {
	Maintainer    string
	RepositoryURL string
	Groups        []Group
}

type Group struct {
	Name        string
	Description string
	Schedule    string
	Links       []Link
}

type Link struct {
	Label    string
	URL      string
	LinkText string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	staticFS, _ := fs.Sub(embedFS, "static")

	http.HandleFunc("/", index)
	http.Handle("/_/", http.StripPrefix("/_/", http.FileServer(http.FS(staticFS))))

	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	config, err := readConfig()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = templates.ExecuteTemplate(w, "index.html", config)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func readConfig() (*Config, error) {

	f, err := embedFS.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("%s not found", configFile)
	}
	defer f.Close()

	config := &Config{}
	err = json.NewDecoder(f).Decode(config)
	if err != nil {
		return nil, fmt.Errorf("%s is invalid", configFile)
	}

	return config, nil
}
