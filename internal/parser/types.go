package parser

type Pattern struct {
	Pattern    string
	Browser    string
	Version    string
	Platform   string
	DeviceType string
	RawFields  map[string]string
}

type Result struct {
	Browser    string
	Version    string
	Platform   string
	DeviceType string
	Matched    string
}
