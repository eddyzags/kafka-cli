package brokers

import (
	"github.com/kafkactl/api/client"
)

type BrokerAdd struct{}

func Add(api client.APIClient, in *BrokerAdd) error {}
