package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	configFile = "leedsgeeks.json"
	port       = flag.String("port", "5000", "http port")
	templates  = template.Must(template.ParseGlob("templates/*.html"))
)

type Config struct {
	Groups []Group
}

type Group struct {
	Name        string
	Description string
	Links       []Link
}

type Link struct {
	Label    string
	URL      string
	LinkText string
}

func main() {
	flag.Parse()
	http.HandleFunc("/", index)
	http.Handle("/_/", http.StripPrefix("/_/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":"+*port, nil)
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

	f, err := os.Open(configFile)
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
