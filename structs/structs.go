package structs

type ConfigData struct {
	Host      string
	Port      int
	MocksPath string
}

type MockData struct {
	Headers string
	Body    string
}
