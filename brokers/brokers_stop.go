package brokers

import (
	"fmt"

	"github.com/eddyzags/kafkactl/api/client"
)

func Stop(api client.APIClient, expr string) error {
	brokers, err := api.BrokerStop(expr)
	if err != nil {
		return err
	}

	if len(brokers) > 1 {
		fmt.Printf("The \"%s\" brokers were stopped successfully\n", expr)
	} else {
		fmt.Printf("The \"%s\" broker was stopped successfully\n", expr)
	}

	return nil
}
