package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	configFile = "leedsgeeks.json"
	port       = flag.String("port", "5000", "http port")
)

type Config struct {
	Groups []Group
}

type Group struct {
	Name string
}

func main() {
	flag.Parse()
	http.HandleFunc("/", index)
	http.ListenAndServe(":"+*port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	config, err := readConfig()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, config.Groups[0].Name)
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
