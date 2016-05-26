package client

import (
	"github.com/eddyzags/kafkactl/models"
)

type APIClient interface {
	BrokerList() ([]*models.Broker, error)
	BrokerAdd()
}
