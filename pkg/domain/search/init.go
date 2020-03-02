package search

var (
	defaultGenerator   Generator
	defaultHTTPService HTTP
)

// Do all calls for generator initializations here
// to easily track changes to the default service.
// Generator init fn definitions, e.g. `initenCdGenerator`
// can be defined here or on their separate files
func init() {
	var err error
	if err = initDefaultGenerator(); err == nil {

	}
	err = initDefaultHTTPService()
}

// DefaultHTTPService returns default http service allowing ease of
// spawning service without directly exposing global variable
func DefaultHTTPService() HTTP {
	return defaultHTTPService
}

func initDefaultGenerator() {

}
