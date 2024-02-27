package shippo

import (
	"context"
)

// ParcelService ..
type ParcelService service

// Parcel ..
//
// See https://goshippo.com/docs/reference#parcels
type Parcel struct {
	ObjectFields
	Length       string       `json:"length"`
	Width        string       `json:"width"`
	Height       string       `json:"height"`
	DistanceUnit string       `json:"distance_unit"`
	Weight       string       `json:"weight"`
	MassUnit     string       `json:"mass_unit"`
	Template     string       `json:"template,omitempty"`
	Extra        *ParcelExtra `json:"extra,omitempty"`
	Metadata     string       `json:"metadata,omitempty"`
	ObjectState  string       `json:"object_state,omitempty"`
	Test         bool         `json:"test"`
}

// ParcelInput ..
//
// See https://goshippo.com/docs/reference#parcels
type ParcelInput struct {
	Length       string       `json:"length"`
	Width        string       `json:"width"`
	Height       string       `json:"height"`
	DistanceUnit string       `json:"distance_unit"`
	Weight       string       `json:"weight"`
	MassUnit     string       `json:"mass_unit"`
	Test         bool         `json:"test"`
	Metadata     string       `json:"metadata,omitempty"`
	Template     string       `json:"template,omitempty"`
	Extra        *ParcelExtra `json:"extra,omitempty"`
}

// ParcelExtra ..
type ParcelExtra struct {
	COD        *ParcelCOD       `json:"COD,omitempty"`
	Insurance  *ParcelInsurance `json:"insurance,omitempty"`
	Reference1 string           `json:"reference_1,omitempty"`
	Reference2 string           `json:"reference_2,omitempty"`
}

// ParcelCOD ..
type ParcelCOD struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}

// ParcelInsurance ..
type ParcelInsurance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Content  string `json:"content"`
}

// Create creates a new parcel object.
func (s *ParcelService) Create(ctx context.Context, in *ParcelInput) (*Parcel, error) {
	return create[ParcelInput, Parcel](ctx, s.client, "/parcels/", in)
}

// Get returns the user created parcels
func (s *ParcelService) Get(ctx context.Context, page, size uint) ([]Parcel, error) {
	return list[Parcel](ctx, s.client, "/parcels", nil, page, size)
}

// GetByID returns a parcel by its ID
func (s *ParcelService) GetByID(ctx context.Context, id string) (*Parcel, error) {
	return get[Parcel](ctx, s.client, "/shipments/", id)
}

// GetTemplates returns templates for carriers which the user has added and enabled.
func (s *ParcelService) GetTemplates(ctx context.Context, include string) ([]CarrierParcelTemplate, error) {
	var params *string
	if include != "" {
		p := "include=" + include
		params = &p
	}

	return list[CarrierParcelTemplate](ctx, s.client, "/parcel-templates", params, 0, 0)
}
