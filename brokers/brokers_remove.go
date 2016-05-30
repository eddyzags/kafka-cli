package brokers

import (
	"fmt"
	"strings"

	"github.com/eddyzags/kafkactl/api/client"
)

func Remove(api client.APIClient, expr string) error {
	ids, err := api.BrokerRemove(expr)
	if err != nil {
		return err
	}

	if strings.Contains(ids, ",") {
		fmt.Printf("The brokers \"%s\" were stopped successfully.\n", ids)
	} else {
		fmt.Printf("The broker \"%s\" was stopped successfully.\n", ids)
	}

	return nil
}
