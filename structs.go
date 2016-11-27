package main

type ConfigData struct {
	Host          string
	Port          int
	RequestsPath  string
	ResponsesPath string
}

type RequestData struct {
	Method   string
	Headers  map[string]string
	Path     string
	Query    string
	Body     map[string]string
	RawBody  string
	Response ResponseData
}

type ResponseData struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type MockData struct {
	Requests []RequestData
}
