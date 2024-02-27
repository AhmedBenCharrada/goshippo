package shippo

import (
	"context"
	"fmt"
)

type AddressService service

// AddressInput ..
//
// See https://goshippo.com/docs/reference#addresses
type AddressInput struct {
	Name          string `json:"name"`
	Company       string `json:"company,omitempty"`
	StreetNo      string `json:"street_no,omitempty"`
	Street1       string `json:"street1"`
	Street2       string `json:"street2,omitempty"`
	Street3       string `json:"street3,omitempty"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Country       string `json:"country"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	IsResidential bool   `json:"is_residential"`
	Validate      bool   `json:"validate"`
	Metadata      string `json:"metadata,omitempty"`
}

// CheckRequiredFields checks the existence of required fields.
func (in *AddressInput) CheckRequiredFields() error {
	if len(in.Name) == 0 {
		return fmt.Errorf("name is required")
	}
	if len(in.Street1) == 0 {
		return fmt.Errorf("street1 is required")
	}
	if len(in.City) == 0 {
		return fmt.Errorf("city is required")
	}

	if len(in.State) == 0 {
		return fmt.Errorf("state is required")
	}
	if len(in.Zip) == 0 {
		return fmt.Errorf("zip is required")
	}
	if len(in.Country) != 2 {
		return fmt.Errorf("country is required")
	}

	return nil
}

// Address ..
//
// See https://goshippo.com/docs/reference#addresses
type Address struct {
	ObjectFields
	ValidationResults *ValidationResults `json:"validation_results"`
	Name              string             `json:"name"`
	Company           string             `json:"company,omitempty"`
	StreetNo          string             `json:"street_no,omitempty"`
	Street1           string             `json:"street1"`
	Street2           string             `json:"street2,omitempty"`
	Street3           string             `json:"street3,omitempty"`
	City              string             `json:"city"`
	State             string             `json:"state"`
	Zip               string             `json:"zip"`
	Country           string             `json:"country"`
	Phone             string             `json:"phone,omitempty"`
	Email             string             `json:"email,omitempty"`
	IsResidential     bool               `json:"is_residential"`
	IsCompleted       bool               `json:"is_complete"`
	Metadata          string             `json:"metadata,omitempty"`
	Test              bool               `json:"test,omitempty"`
}

// ValidationResults ..
type ValidationResults struct {
	IsValid  bool             `json:"is_valid"`
	Messages []*OutputMessage `json:"messages"`
}

// Create creates a new address object.
func (s *AddressService) Create(ctx context.Context, in *AddressInput) (*Address, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input")
	}
	if err := in.CheckRequiredFields(); err != nil {
		return nil, err
	}
	// make sure the address is valid
	in.Validate = true

	return create[AddressInput, Address](ctx, s.client, "/addresses/", in)
}

// Get lists all addresses.
func (s *AddressService) Get(ctx context.Context, page, size uint) ([]Address, error) {
	return list[Address](ctx, s.client, "/addresses/", nil, page, size)
}

// GetByID retrieve an existing address by object id.
func (s *AddressService) GetByID(ctx context.Context, id string) (*Address, error) {
	return get[Address](ctx, s.client, "/addresses/", id)
}

// Validate validates an existing address by object id.
func (s *AddressService) Validate(ctx context.Context, id string) (*Address, error) {
	return get[Address](ctx, s.client, "/addresses/", id, "/validate/")
}
