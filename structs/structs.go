package structs

type ConfigData struct {
	Host      string
	Port      int
	MocksPath string
}

type RequestData struct {
	Method   string
	Headers  map[string]string
	Path     string
	Query    string
	Body     string
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
