package shippo

import (
	"context"
)

type RateService service

// Get retrieve an existing rate by object id.
func (s *RateService) Get(ctx context.Context, id string) (*ShipmentRate, error) {
	return get[ShipmentRate](ctx, s.client, "/rates/", id)
}
