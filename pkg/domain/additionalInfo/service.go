package additionalinfo

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/Zumata/v3-common/data"
)

// Worker is function type simulating specific booking behaviour
type Worker func(context.Context, *data.AdditionalInfoRequest) (*data.AdditionalInfo, error)

// Server implements AdditionalInfo service. It maps property codes
// to functions that simulate certain behaviour
type Server struct {
	logger *zerolog.Logger
	mapper map[string]Worker
}

// AdditionalInfo ...
func (ap *AdditionalInfoProps) AdditionalInfo(ctx context.Context, req *data.AdditionalInfoRequest) (*data.AdditionalInfo, error) {
	
	return nil, nil
}
