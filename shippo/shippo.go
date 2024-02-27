package shippo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.goshippo.com"
	authKey = "ShippoToken "
	version = "2018-02-08"
)

type Client interface {
	NewRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error)
	Do(req *http.Request, v interface{}) (*http.Response, error)

	Address() *AddressService
	Parcel() *ParcelService
	Shipment() *ShipmentService
	Rate() *RateService
	Carrier() *CarrierService
	Transaction() *TransactionService
}

type service struct {
	client Client
}

type client struct {
	client  *http.Client
	baseURL *url.URL
	token   string

	address     *AddressService
	parcel      *ParcelService
	shipment    *ShipmentService
	rate        *RateService
	carrier     *CarrierService
	transaction *TransactionService
	tracking    *TrackingService
}

var defaultHTTPClient = &http.Client{Timeout: time.Second * 5}

func NewClient(token string, options ...ClientOption) (Client, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &client{
		client:  defaultHTTPClient,
		baseURL: url,
		token:   token,
	}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	svc := &service{client: client}
	client.address = (*AddressService)(svc)
	client.parcel = (*ParcelService)(svc)
	client.shipment = (*ShipmentService)(svc)
	client.rate = (*RateService)(svc)
	client.carrier = (*CarrierService)(svc)
	client.transaction = (*TransactionService)(svc)
	client.tracking = (*TrackingService)(svc)
	return client, nil
}

func (c *client) Address() *AddressService {
	return c.address
}

func (c *client) Parcel() *ParcelService {
	return c.parcel
}

func (c *client) Shipment() *ShipmentService {
	return c.shipment
}

func (c *client) Rate() *RateService {
	return c.rate
}

func (c *client) Carrier() *CarrierService {
	return c.carrier
}

func (c *client) Transaction() *TransactionService {
	return c.transaction
}

func (c *client) Tracking() *TrackingService {
	return c.tracking
}

func (c *client) NewRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	fullUrl, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, fullUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	req.Header.Set("Authorization", authKey+c.token)
	req.Header.Set("SHIPPO-API-VERSION", version)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkForErrors(resp)
	if err != nil {
		return resp, err
	}

	if resp.Body != nil && v != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}

		err = json.Unmarshal(body, v)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func checkForErrors(resp *http.Response) error {
	if c := resp.StatusCode; c >= 200 && c < 400 {
		return nil
	}

	return &ErrorResponse{Response: resp}
}
