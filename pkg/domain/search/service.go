package search

import (
	"context"
	"sync"

	"github.com/Zumata/v3-common/data"
)

// Service is search domain interface
type Service interface {
	Search(context.Context, *data.SearchRequest) (*data.SearchBasic, error)
}

// server implements Search service. It maps properties
// to functions that simulate certain behaviour
type server struct {
	mapper map[string]Generator
}

// Search ...
func (s *server) Search(ctx context.Context, req *data.SearchRequest) (*data.SearchBasic, error) {
	var wg sync.WaitGroup
	results := make(chan *data.SearchBasic, len(req.Params.PartnerHotelIDs))

	for _, hotelID := range req.Params.PartnerHotelIDs {
		generator, ok := s.mapper[hotelID]
		if !ok {
			generator = defaultGenerator
		}

		reqCopy := *req
		reqCopy.Params.PartnerHotelIDs = []string{hotelID}

		wg.Add(1)
		go func(results chan *data.SearchBasic, wg *sync.WaitGroup) {
			s, err := generator.Package(context.Background(), &reqCopy)
			if err != nil {
				// TODO: handle error; context cancellation maybe?
			}
			results <- s
			wg.Done()
		}(results, &wg)
	}

	// await results
	wg.Wait()

	// process results
	basic := new(data.SearchBasic)
	for {
		r, more := <-results
		if !more {
			// when done replace channel with nil instead of
			// simply closing channel to avoid busy loop
			results = nil
			break
		}
		for k, v := range r.Packages {
			basic.Packages[k] = v
		}
	}

	return basic, nil
}
