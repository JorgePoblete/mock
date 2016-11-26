package main

import (
	"encoding/json"
	"fmt"
	"github.com/JorgePoblete/mock/handlers"
	"github.com/JorgePoblete/mock/structs"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "received request to [%s]%s?%s\n", r.Method, r.URL.Path, r.URL.RawQuery)
}

func loadConfig(filePath string) structs.ConfigData {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error reading conf file.\n%s\n", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	conf := structs.ConfigData{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Printf("error decoding conf file.\n%s\n", err)
		os.Exit(1)
	}
	return conf
}

func main() {
	conf := loadConfig("config/config.json")
	fmt.Printf("listing on %s:%d\n", conf.Host, conf.Port)
	http.HandleFunc("/", handlers.RequestHandler)
	http.ListenAndServe(conf.Host+":"+strconv.Itoa(conf.Port), nil)
}
