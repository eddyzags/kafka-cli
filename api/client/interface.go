package client

import (
	"github.com/eddyzags/kafkactl/models"
)

type APIClient interface {
	// Broker handling functions
	BrokerAdd(in map[string]string) ([]*models.Broker, error)
	BrokerStart(expr, timeout string) (string, error)
	BrokerList() ([]*models.Broker, error)
}
