package search

var defaultHTTPService HTTP

func init() {
	defaultHTTPService = HTTP{
		service: &server{
			mapper: make(map[string]Generator),
		},
	}
}

// DefaultHTTPService returns default http service allowing ease of
// spawning service without directly exposing global variable
func DefaultHTTPService() HTTP {
	return defaultHTTPService
}
