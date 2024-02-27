package shippo

import (
	"context"
	"time"
)

type PickupService service

// Pickup ..
//
// See https://goshippo.com/docs/reference#pickups
type Pickup struct {
	ObjectFields
	CarrierAccount     string          `json:"carrier_account"`
	Location           Location        `json:"location"`
	Metadata           string          `json:"metadata"`
	RequestedEndTime   *time.Time      `json:"requested_end_time"`
	RequestedStartTime *time.Time      `json:"requested_start_time"`
	Transactions       []string        `json:"transactions"`
	ConfirmedStartTime *time.Time      `json:"confirmed_start_time"`
	ConfirmedEndTime   *time.Time      `json:"confirmed_end_time"`
	CancelByTime       *time.Time      `json:"cancel_by_time"`
	Status             string          `json:"status"`
	ConfirmationCode   string          `json:"confirmation_code"`
	Timezone           string          `json:"timezone"`
	Messages           []OutputMessage `json:"messages"`
	IsTest             bool            `json:"is_test"`
}

// PickupRequest ..
//
// See https://goshippo.com/docs/reference#pickups
type PickupRequest struct {
	CarrierAccount     string     `json:"carrier_account"`
	Location           Location   `json:"location"`
	Metadata           string     `json:"metadata"`
	RequestedEndTime   *time.Time `json:"requested_end_time"`
	RequestedStartTime *time.Time `json:"requested_start_time"`
	Transactions       []string   `json:"transactions"`
}

// Location ..
type Location struct {
	Address              Address `json:"address"`
	BuildingLocationType string  `json:"building_location_type"`
	BuildingType         string  `json:"building_type"`
	Instructions         string  `json:"instructions"`
}

// Create creates a pickup object. This request is for a carrier to come to a specified location to take a package for shipping.
func (s *PickupService) Create(ctx context.Context, in *PickupRequest) (*Pickup, error) {
	return create[PickupRequest, Pickup](ctx, s.client, "/pickups/", in)
}
