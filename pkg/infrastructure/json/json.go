package json

import (
	"encoding/json"

	"github.com/json-iterator/go"
)

var (
	// JSON initialise json iter lib
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal initialise marshal lib
	Marshal = JSON.Marshal
	// Unmarshal initialise unmarshal lib
	Unmarshal = JSON.Unmarshal
	// NewDecoder wrap decoder of json iter
	NewDecoder = json.NewDecoder
	// NewEncoder wrap encoder of json iter
	NewEncoder = json.NewEncoder
)
