package shippo

import (
	"context"
)

type CarrierService service

// Carrier ..
//
// See https://goshippo.com/docs/reference#carrier-accounts
type Carrier struct {
	AccountID       string           `json:"account_id"`
	Active          bool             `json:"active"`
	Carrier         string           `json:"carrier"`
	Parameters      CarrierParameter `json:"parameters"`
	IsShippoAccount bool             `json:"is_shippo_account"`
	Metadata        string           `json:"metadata"`
	ObjectID        string           `json:"object_id"`
	ObjectOwner     string           `json:"object_owner"`
	Test            bool             `json:"test"`
}

// CarrierParameter ..
type CarrierParameter struct {
	AccountNumber             string `json:"account_number"`
	AiaCountryIso2            string `json:"aia_country_iso2"`
	BillingAddressCity        string `json:"billing_address_city"`
	BillingAddressCountryIso2 string `json:"billing_address_country_iso2"`
	BillingAddressState       string `json:"billing_address_state"`
	BillingAddressStreet1     string `json:"billing_address_street1"`
	BillingAddressStreet2     string `json:"billing_address_street2"`
	BillingAddressZip         string `json:"billing_address_zip"`
	CollecCountryIso2         string `json:"collec_country_iso2"`
	CollecZip                 string `json:"collec_zip"`
	Company                   string `json:"company"`
	CurrencyCode              string `json:"currency_code"`
	Email                     string `json:"email"`
	FullName                  string `json:"full_name"`
	HasInvoice                bool   `json:"has_invoice"`
	InvoiceControlId          string `json:"invoice_controlid"`
	InvoiceDate               string `json:"invoice_date"`
	InvoiceNumber             string `json:"invoice_number"`
	InvoiceValue              string `json:"invoice_value"`
	Phone                     string `json:"phone"`
	Title                     string `json:"title"`
	UpsAgreements             bool   `json:"ups_agreements"`
}

// See https://goshippo.com/docs/reference#carrier-parcel-templates
type CarrierParcelTemplate struct {
	Name                 string `json:"name"`
	Token                string `json:"token"`
	Carrier              string `json:"carrier"`
	IsVariableDimensions bool   `json:"is_variable_dimensions"`
	Length               string `json:"length"`
	Width                string `json:"width"`
	Height               string `json:"height"`
	DistanceUnit         string `json:"distance_unit"`
}

// Get returns a list of all carrier accounts.
func (s *CarrierService) Get(ctx context.Context, page, size uint) ([]Carrier, error) {
	return list[Carrier](ctx, s.client, "/carrier_accounts", nil, page, size)
}
