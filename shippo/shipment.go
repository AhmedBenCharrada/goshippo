package shippo

import (
	"context"
	"time"
)

type ShipmentService service

// Shipmen ..
//
// See https://goshippo.com/docs/reference#shipments
type Shipment struct {
	ObjectFields
	CustomsDeclaration string           `json:"customs_declaration"`
	Extra              ShipmentExtra    `json:"extra"`
	Metadata           string           `json:"metadata"`
	Date               time.Time        `json:"shipment_date"`
	From               ShipmentAddress  `json:"address_from"`
	Return             ShipmentAddress  `json:"address_return"`
	To                 ShipmentAddress  `json:"address_to"`
	CarrierAccounts    []string         `json:"carrier_accounts"`
	Messages           []OutputMessage  `json:"messages"`
	Parcels            []ParcelTemplate `json:"parcels"`
	Rates              []ShipmentRate   `json:"rates"`
	Status             ShipmentStatus   `json:"status"`
	Test               bool             `json:"test"`
}

// ShipmentRequest ..
//
// See https://goshippo.com/docs/reference#shipments
type ShipmentRequest struct {
	CustomsDeclaration string           `json:"customs_declaration"`
	Date               *time.Time       `json:"shipment_date,omitempty"`
	From               ShipmentAddress  `json:"address_from"`
	To                 ShipmentAddress  `json:"address_to"`
	Return             ShipmentAddress  `json:"address_return,omitempty"`
	Extra              *ShipmentExtra   `json:"extra,omitempty"`
	Parcels            []ShipmentParcel `json:"parcels"`
	Metadata           string           `json:"metadata"`
	Async              bool             `json:"async"`
}

// ShipmentAddress ..
type ShipmentAddress struct {
	City          string `json:"city"`
	Company       string `json:"company"`
	Country       string `json:"country,required"`
	Email         string `json:"email"`
	IsResidential bool   `json:"is_residential"`
	Metadata      string `json:"metadata"`
	Name          string `json:"name,required"`
	Phone         string `json:"phone"`
	State         string `json:"state"`
	Street1       string `json:"street1,required"`
	Street2       string `json:"street2"`
	Street3       string `json:"street3"`
	StreetNo      string `json:"street_no"`
	Zip           string `json:"zip,required"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Validate      bool   `json:"validate"`
}

type ShipmentParcel struct {
	DistanceUnit string `json:"distance_unit"`
	Height       string `json:"height"`
	Length       string `json:"length"`
	MassUnit     string `json:"mass_unit"`
	Template     string `json:"template,omitempty"`
	Weight       string `json:"weight"`
	Width        string `json:"width"`
	Metadata     string `json:"metadata"`
}

type AncillaryServiceEndorsement string

const (
	FORWARDING_SERVICE_REQUESTED AncillaryServiceEndorsement = "FORWARDING_SERVICE_REQUESTED"
	RETURN_SERVICE_REQUESTED     AncillaryServiceEndorsement = "RETURN_SERVICE_REQUESTED"
)

type ShipmentExtra struct {
	AncillaryEndorsement       AncillaryServiceEndorsement `json:"ancillary_endorsement"`
	AuthorityToLeave           bool                        `json:"authority_to_leave"`
	Alcohol                    *Alcohol                    `json:"alcohol,omitempty"`
	Billing                    *Billing                    `json:"billing,omitempty"`
	BypassAddressValidation    bool                        `json:"bypass_address_validation"`
	CarbonNeutral              bool                        `json:"carbon_neutral"`
	CarrierHubID               string                      `json:"carrier_hub_id"`
	CarrierHubTravelTime       int                         `json:"carrier_hub_travel_time"`
	Cod                        *Cod                        `json:"COD,omitempty"`
	ContainerType              string                      `json:"container_type"`
	CriticalPullTime           string                      `json:"critical_pull_time"`
	CustomerBranch             string                      `json:"customer_branch"`
	CustomerReference          *Reference                  `json:"customer_reference,omitempty"`
	DangerousGoodsCode         string                      `json:"dangerous_goods_code"`
	DangerousGoods             *DangerousGoods             `json:"dangerous_goods,omitempty"`
	DeliveryInstructions       string                      `json:"delivery_instructions"`
	DeptNumber                 *Reference                  `json:"dept_number,omitempty"`
	DryIce                     *DryIce                     `json:"dry_ice,omitempty"`
	FulfillmentCenter          string                      `json:"fulfillment_center"`
	Insurance                  *Insurance                  `json:"insurance,omitempty"`
	InvoiceNumber              *InvoiceNumber              `json:"invoice_number,omitempty"`
	IsReturn                   bool                        `json:"is_return"`
	LasershipAttrs             string                      `json:"lasership_attrs"`
	LasershipDeclaredValue     string                      `json:"lasership_declared_value"`
	PoNumber                   *Reference                  `json:"po_number,omitempty"`
	PreferredDeliveryTimeframe string                      `json:"preferred_delivery_timeframe"`
	Premium                    bool                        `json:"premium"`
	QrCodeRequested            bool                        `json:"qr_code_requested"`
	Reference1                 string                      `json:"reference_1"`
	Reference2                 string                      `json:"reference_2"`
	RequestRetailRates         bool                        `json:"request_retail_rates"`
	ReturnServiceType          string                      `json:"return_service_type"`
	RmaNumber                  *Reference                  `json:"rma_number,omitempty"`
	SaturdayDelivery           bool                        `json:"saturday_delivery"`
	SignatureConfirmation      string                      `json:"signature_confirmation"`
}

type RecipientType string

const (
	LICENSEE RecipientType = "licensee"
	CONSUMER RecipientType = "consumer"
)

type Alcohol struct {
	Included  bool          `json:"contains_alcohol"`
	Recipient RecipientType `json:"recipient_type"`
}
type Billing struct {
	Account           string `json:"account"`
	Country           string `json:"country"`
	ParticipationCode string `json:"participation_code"`
	Type              string `json:"type"`
	Zip               string `json:"zip"`
}

type PaymentMethod string

const (
	SECURED_FUNDS PaymentMethod = "SECURED_FUNDS"
	CASH          PaymentMethod = "CASH"
	ANY           PaymentMethod = "ANY"
)

type Cod struct {
	Amount   string        `json:"amount"`
	Currency string        `json:"currency"`
	Method   PaymentMethod `json:"payment_method"`
}

type DangerousGoods struct {
	Contains            bool `json:"contains"`
	BiologicalMaterials *struct {
		Contains bool `json:"contains"`
	} `json:"biological_materials"`
	LithiumBatteries struct {
		Contains bool `json:"contains"`
	} `json:"lithium_batteries"`
}

type DryIce struct {
	Contains bool   `json:"contains"`
	Weight   string `json:"weight"`
}

type InsuranceProvider string

const (
	FEDEX  InsuranceProvider = "FEDEX"
	UPS    InsuranceProvider = "UPS"
	ONTRAC InsuranceProvider = "ONTRAC"
)

type Insurance struct {
	Amount   string            `json:"amount"`
	Currency string            `json:"currency"`
	Content  string            `json:"content"`
	Provider InsuranceProvider `json:"provider"`
}
type InvoiceNumber struct {
	Prefix string `json:"prefix"`
	Value  string `json:"value"`
}

type Reference struct {
	Prefix string `json:"prefix"`
	Value  string `json:"value"`
}

type ShipmentStatus = string

const (
	WAITING ShipmentStatus = "WAITING"
	QUEUED  ShipmentStatus = "QUEUED"
	SUCCESS ShipmentStatus = "SUCCESS"
	ERROR   ShipmentStatus = "ERROR"
)

type ParcelTemplate struct {
	DistanceUnit string `json:"distance_unit"`
	Height       string `json:"height"`
	Default      bool   `json:"is_default"`
	Length       string `json:"length"`
	MassUnit     string `json:"mass_unit"`
	Name         string `json:"name"`
	ObjectID     string `json:"object_id"`
	Template     string `json:"template"`
	Weight       string `json:"weight"`
	Width        string `json:"width"`
}
type RateAttribute string

const (
	BESTVALUE RateAttribute = "BESTVALUE"
	CHEAPEST  RateAttribute = "CHEAPEST"
	FASTEST   RateAttribute = "FASTEST"
)

// ShipmentRate ..
//
// See https://goshippo.com/docs/reference#rates
type ShipmentRate struct {
	Amount                 string          `json:"amount"`
	AmountLocal            string          `json:"amount_local"`
	ArrivesBy              any             `json:"arrives_by"`
	Attributes             []RateAttribute `json:"attributes"`
	CarrierAccount         string          `json:"carrier_account"`
	Currency               string          `json:"currency"`
	CurrencyLocal          string          `json:"currency_local"`
	DurationTerm           string          `json:"duration_terms"`
	EstimatedDays          int64           `json:"estimated_days"`
	IncludedInsurancePrice int64           `json:"included_insurance_price"`
	Messages               []OutputMessage `json:"messages"`
	ObjectFields
	Provider         string        `json:"provider"`
	ProviderImage75  string        `json:"provider_image_75"`
	ProviderImage200 string        `json:"provider_image_200"`
	ServiceLevel     *ServiceLevel `json:"servicelevel"`
	Shipment         string        `json:"shipment"`
	Test             bool          `json:"test"`
	Zone             string        `json:"zone"`
}

type ServiceLevel struct {
	Name          string `json:"name"`
	Terms         string `json:"terms"`
	Token         string `json:"service_level_token"`
	ExtendedToken string `json:"extended_token"`
	Parent        string `json:"parent_servicelevel"`
}

// Create creates a new shipment object.
func (s *ShipmentService) Create(ctx context.Context, in *ShipmentRequest) (*Shipment, error) {
	return create[ShipmentRequest, Shipment](ctx, s.client, "/shipments/", in)
}

// Get list all shipment objects.
func (s *ShipmentService) Get(ctx context.Context, dateFilter *DateFilter, page, size uint) ([]Shipment, error) {
	var params *string
	if dateFilter != nil {
		p := dateFilter.ToPathParam()
		params = &p
	}

	return list[Shipment](ctx, s.client, "/shipments/", params, page, size)
}

// GetByID retrieve an existing shipment by object id.
func (s *ShipmentService) GetByID(ctx context.Context, id string) (*Shipment, error) {
	return get[Shipment](ctx, s.client, "/shipments/", id)
}

// RatesFor returns the rates for the shipment.
func (s *ShipmentService) RatesFor(ctx context.Context, id string, currency string) ([]ShipmentRate, error) {
	url := "/shipments/" + id + "/rates/" + currency
	return list[ShipmentRate](ctx, s.client, url, nil, 0, 0)
}

// GetRate retrieve an existing rate by object id.
func (s *ShipmentService) GetRate(ctx context.Context, id string) (*ShipmentRate, error) {
	return get[ShipmentRate](ctx, s.client, "/rates/", id)
}
