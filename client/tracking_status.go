package client

import (
	"errors"
	"net/http"

	"github.com/phil-inc/go-shippo/models"
)

// GetTrackingUpdate requests the tracking status of a shipment.
func (c *Client) GetTrackingUpdate(carrier, trackingNumber string) (*models.TrackingStatus, error) {
	if carrier == "" {
		return nil, errors.New("Empty carrier")
	}
	if trackingNumber == "" {
		return nil, errors.New("Empty tracking number")
	}

	output := &models.TrackingStatus{}
	err := c.do(http.MethodGet, "/tracks/"+carrier+"/"+trackingNumber, nil, output)
	return output, err
}

// RegisterTrackingWebhook registers a tracking webhook.
// TODO: documentation on this API endpoint is not clear.
// https://goshippo.com/docs/reference#track-create
func (c *Client) RegisterTrackingWebhook(carrier, trackingNumber string) (*models.TrackingStatus, error) {
	if carrier == "" {
		return nil, errors.New("Empty carrier")
	}
	if trackingNumber == "" {
		return nil, errors.New("Empty tracking number")
	}

	output := &models.TrackingStatus{}
	err := c.do(http.MethodPost, "/tracks/", &models.TrackingStatusInput{
		Carrier:        carrier,
		TrackingNumber: trackingNumber,
	}, output)
	return output, err
}
