package shippo

import (
	"context"
)

type TransactionService service

type LabelFileType string

const (
	PNG         LabelFileType = "PNG"
	PDF         LabelFileType = "PDF"
	PNG_2_3x7_5 LabelFileType = "PNG_2.3x7.5"
	PDF4x6      LabelFileType = "PDF_4x6"
	PDF4x8      LabelFileType = "PDF_4x8"
	PDF_A4      LabelFileType = "PDF_A4"
	PDF_A6      LabelFileType = "PDF_A6"
	ZPLII       LabelFileType = "ZPLII"
)

type TransactionCreateRequest struct {
	Rate          string        `json:"rate"`
	Metadata      string        `json:"metadata"`
	LabelFileType LabelFileType `json:"label_file_type"`
	Async         bool          `json:"async"`
}

type InstantTransactionRequest struct {
	CarrierAccount    string        `json:"carrier_account"`
	LabelFileType     LabelFileType `json:"label_file_type"`
	Metadata          string        `json:"metadata"`
	ServiceLevelToken string        `json:"servicelevel_token"`
	Shipment          Shipment      `json:"shipment"`
}

type Transaction struct {
	ObjectFields
	CommercialInvoiceURL string          `json:"commercial_invoice_url"`
	Eta                  string          `json:"eta"`
	LabelFileType        LabelFileType   `json:"label_file_type"`
	LabelURL             string          `json:"label_url"`
	Messages             []OutputMessage `json:"messages"`
	Metadata             string          `json:"metadata"`
	QrCodeURL            string          `json:"qr_code_url"`
	Rate                 string          `json:"rate"`
	Status               string          `json:"status"`
	Test                 bool            `json:"test"`
	TrackingNumber       string          `json:"tracking_number"`
	TrackingStatus       string          `json:"tracking_status"`
	TrackingURLProvider  string          `json:"tracking_url_provider"`
}

// CreateWithRate creates a new transaction object and purchases the shipping label using a rate object that has previously been created.
func (s *TransactionService) CreateWithRate(ctx context.Context, in *TransactionCreateRequest) (*Transaction, error) {
	return create[TransactionCreateRequest, Transaction](ctx, s.client, "/transactions", in)
}

// CreateWithShipment creates a new transaction object and purchases the shipping label instantly using shipment details, an existing carrier account, and an existing service level token.
func (s *TransactionService) CreateWithShipment(ctx context.Context, in *InstantTransactionRequest) (*Transaction, error) {
	return create[InstantTransactionRequest, Transaction](ctx, s.client, "/transactions", in)
}

// Get returns a list of all transaction objects.
func (s *TransactionService) Get(ctx context.Context, page, size uint) ([]Transaction, error) {
	return list[Transaction](ctx, s.client, "/transactions", nil, page, size)
}
