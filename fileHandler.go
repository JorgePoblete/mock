package main

import (
	_ "bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "regexp"
	"strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	file := strings.Split(r.URL.Path, "/")[1]
	requests, err := loadRequestsFile(Conf.RequestsPath + file + ".json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, request := range requests.Requests {
		if !compareRequests(r, request) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		Body, err := loadFile(Conf.ResponsesPath + request.Response.Body + ".resp")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for key, value := range request.Response.Headers {
			w.Header().Set(key, value)
		}
		w.WriteHeader(request.Response.StatusCode)
		fmt.Fprint(w, Body)
		return
	}
}

func compareRequests(a *http.Request, b RequestData) bool {
	if a.Method != b.Method {
		return false
	}
	if a.URL.Path != b.Path {
		return false
	}
	if a.URL.RawQuery != b.Query {
		return false
	}
	for key, value := range b.Headers {
		if a.Header.Get(key) != value {
			return false
		}
	}
	// Use regex validator to validate if the body of the request corresponds withe the one the mock has.
	// posible regex [gata=1|gata:1|name="gata"\n\n1]
	//	rawBody, _ := ioutil.ReadAll(a.Body)
	//	a.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	//	fmt.Printf("%s\n\n", rawBody)
	//	regexValidator := regexp.MustCompile("")
	//	formDataString := `[a-zA-Z-:;=" ]{38}%s["\n]{3}%s`
	//	fmt.Printf("%+v\n", regexValidator.MatchString(formDataString))
	return true
}

func loadRequestsFile(filePath string) (MockData, error) {
	file, err := loadFile(filePath)
	if err != nil {
		fmt.Printf("Error loading request file %s.\n%s\n", filePath, err)
		return MockData{}, errors.New("ERROR_LODADING_FILE")
	}
	requests := make([]RequestData, 0)
	err = json.Unmarshal([]byte(file), &requests)
	if err != nil {
		fmt.Printf("error decoding %s file.\n%s\n", filePath, err)
		return MockData{}, errors.New("ERROR_DECODING_FILE")
	}

	return MockData{requests}, nil
}

func checkFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func loadFile(filePath string) (string, error) {
	if !checkFileExist(filePath) {
		return "", errors.New("FILE_DOESNT_EXISTS")
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("ERROR_READING_FILE")
	}
	return string(file), nil
}
