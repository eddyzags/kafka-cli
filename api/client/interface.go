package client

import (
	"github.com/eddyzags/kafkactl/models"
)

type APIClient interface {
	// Broker handling functions
	BrokerAdd(in map[string]string) ([]*models.Broker, error)
	BrokerStart(expr, timeout string) ([]*models.Broker, error)
	BrokerList() ([]*models.Broker, error)
	BrokerStop(expr, timeout string, stop bool) ([]*models.Broker, error)
	BrokerRemove(expr string) (string, error)
}
