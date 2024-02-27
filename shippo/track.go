package shippo

import (
	"context"
	"time"
)

// TrackingService ..
type TrackingService service

// TrackRequest ..
//
// See https://goshippo.com/docs/reference#tracks
type TrackRequest struct {
	Carrier        string `json:"carrier"`
	TrackingNumber string `json:"tracking_number"`
	Messages       string `json:"messages"`
}

// TrackingResponse ..
//
// See https://goshippo.com/docs/reference#tracks
type TrackingResponse struct {
	AddressFrom     TrackingAddress  `json:"address_from"`
	AddressTo       TrackingAddress  `json:"address_to"`
	Carrier         string           `json:"carrier"`
	Eta             *time.Time       `json:"eta"`
	Messages        []string         `json:"messages"`
	Metadata        string           `json:"metadata"`
	OriginalEta     *time.Time       `json:"original_eta"`
	ServiceLevel    ServiceLevel     `json:"servicelevel"`
	TrackingHistory []TrackingStatus `json:"tracking_history"`
	TrackingNumber  string           `json:"tracking_number"`
	TrackingStatus  TrackingStatus   `json:"tracking_status"`
	Transaction     string           `json:"transaction"`
}

// ShipmentTrackingStatus ..
type ShipmentTrackingStatus string

// Shipment tracking statuses.
const (
	UNKNOWN     ShipmentTrackingStatus = "UNKNOWN"
	PRE_TRANSIT ShipmentTrackingStatus = "PRE_TRANSIT"
	TRANSIT     ShipmentTrackingStatus = "TRANSIT"
	DELIVERED   ShipmentTrackingStatus = "DELIVERED"
	RETURNED    ShipmentTrackingStatus = "RETURNED"
	FAILURE     ShipmentTrackingStatus = "FAILURE"
)

// TrackingStatus ..
type TrackingStatus struct {
	ObjectFields
	Location      TrackingAddress        `json:"location"`
	Status        ShipmentTrackingStatus `json:"status"`
	StatusDate    time.Time              `json:"status_date"`
	StatusDetails string                 `json:"status_details"`
}

// TrackingAddress ..
type TrackingAddress struct {
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// Register registers a webhook that will send HTTP notifications to you when the status of your tracked package changes.
func (s *TrackingService) Register(ctx context.Context, in *TrackRequest) (*TrackingResponse, error) {
	return create[TrackRequest, TrackingResponse](ctx, s.client, "/tracks/", in)
}

// Get request the tracking status of a shipment by sending a GET request.
func (s *TrackingService) Get(ctx context.Context, carrier, id string) (*TrackingResponse, error) {
	return get[TrackingResponse](ctx, s.client, "/tracks/", carrier+"/", id)
}
