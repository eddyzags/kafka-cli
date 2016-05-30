package brokers

import (
	"fmt"

	"github.com/eddyzags/kafkactl/api/client"
)

func Start(api client.APIClient, expr, timeout string) error {
	ids, err := api.BrokerStart(expr, timeout)
	if err != nil {
		return err
	}

	fmt.Printf("The \"%s\" broker(s) were started successfully\n", ids)

	return nil
}
