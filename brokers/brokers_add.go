package brokers

import (
	"github.com/eddyzags/kafkactl/api/client"
)

type BrokerAdd struct{}

func Add(api client.APIClient, in *BrokerAdd) error {
	return nil
}
