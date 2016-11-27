package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var Conf ConfigData

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "received request to [%s]%s?%s\n", r.Method, r.URL.Path, r.URL.RawQuery)
}

func loadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error reading conf file.\n%s\n", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	Conf = ConfigData{}
	err = decoder.Decode(&Conf)
	if err != nil {
		fmt.Printf("error decoding conf file.\n%s\n", err)
		os.Exit(1)
	}
}

func main() {
	loadConfig("config.json")
	fmt.Printf("listing on %s:%d\n", Conf.Host, Conf.Port)
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(Conf.Host+":"+strconv.Itoa(Conf.Port), nil)
}
