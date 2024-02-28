package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AhmedBenCharrada/goshippo/shippo"
)

func main() {
	client := newClient("your_goshippo_token")

	run(client, getAddresses)
}

func newClient(token string) shippo.Client {
	client, err := shippo.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func run[T any](client shippo.Client, fn func(shippo.Client) (T, error)) {
	res, err := fn(client)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%+v\n", res)
}

func getAddresses(client shippo.Client) ([]shippo.Address, error) {
	return client.Address().Get(context.Background(), 1, 25)
}

func validateAddress(client shippo.Client) (*shippo.Address, error) {
	return client.Address().Validate(context.Background(), "address_id")
}

func getShipmentByID(client shippo.Client) (*shippo.Shipment, error) {
	return client.Shipment().GetByID(context.Background(), "shipment_id")
}

func listShipments(client shippo.Client) ([]shippo.Shipment, error) {
	return client.Shipment().Get(context.Background(), nil, 1, 25)
}

func createShipment(client shippo.Client) (*shippo.Shipment, error) {
	// valid test address
	// 123 SW Heritage Pkwy, Beaverton, OR 97006, United States
	to := shippo.ShipmentAddress{
		City:    "Beaverton",
		Country: "US",
		State:   "OR",
		Zip:     "97006",
		Street1: "123 SW Heritage Pkwy",
	}
	// valid test address
	// 12215 SW main St, Tigard OR 97223
	from := shippo.ShipmentAddress{
		City:    "Tigard",
		Country: "US",
		State:   "OR",
		Zip:     "97223",
		Street1: "12215 SW main St",
	}

	parcels := []shippo.ShipmentParcel{
		{
			DistanceUnit: "in",
			Height:       "1",
			Length:       "1",
			Width:        "1",
			MassUnit:     "lb",
			Weight:       "1",
		},
	}

	in := &shippo.ShipmentRequest{
		To:      to,
		From:    from,
		Return:  from,
		Parcels: parcels,
	}

	return client.Shipment().Create(context.Background(), in)
}

func getShipmentRates(client shippo.Client) ([]shippo.ShipmentRate, error) {
	return client.Shipment().RatesFor(context.Background(), "shipment_id", "USD")
}

func getRate(client shippo.Client) (*shippo.ShipmentRate, error) {
	return client.Rate().Get(context.Background(), "rate_id")
}

func getCarriers(client shippo.Client) ([]shippo.Carrier, error) {
	return client.Carrier().Get(context.Background(), 1, 25)
}

func getTemplates(client shippo.Client) ([]shippo.CarrierParcelTemplate, error) {
	return client.Parcel().GetTemplates(context.Background(), "enabled")
}
