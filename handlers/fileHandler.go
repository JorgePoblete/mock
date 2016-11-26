package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JorgePoblete/mock/structs"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	file := strings.Split(r.URL.Path, "/")[1]
	requests, err := loadRequestsFile("data/" + file + ".json")
	if err != nil {
		return
	}
	for _, request := range requests.Requests {
		if compareRequests(r, request) {
			for key, value := range request.Response.Headers {
				w.Header().Set(key, value)
			}
			w.WriteHeader(request.Response.StatusCode)
			fmt.Fprint(w, request.Response.Body)
			return
		}
	}
}

func compareRequests(a *http.Request, b structs.RequestData) bool {
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
	//	rawBody, _ := ioutil.ReadAll(a.Body)
	//	a.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
	//	fmt.Printf("%s\n\n", rawBody)
	//	regexValidator := regexp.MustCompile("")
	//	formDataString := `[a-zA-Z-:;=" ]{38}%s["\n]{3}%s`
	//	fmt.Printf("%+v\n", regexValidator.MatchString(formDataString))
	return true
}

func loadRequestsFile(filePath string) (structs.MockData, error) {
	// Check if file already exists
	if _, err := os.Stat(filePath); err != nil {
		return structs.MockData{}, errors.New("FILE_DOESNT_EXISTS")
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading %s file.\n%s\n", filePath, err)
		os.Exit(1)
	}
	requests := make([]structs.RequestData, 0)
	err = json.Unmarshal(file, &requests)
	if err != nil {
		fmt.Printf("error decoding %s file.\n%s\n", filePath, err)
		os.Exit(1)
	}
	return structs.MockData{requests}, nil
}
