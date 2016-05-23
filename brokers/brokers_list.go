package brokers

import (
	"fmt"
	"os"

	"github.com/eddyzags/kafkactl/api/client"

	"github.com/olekukonko/tablewriter"
)

func List(api client.APIClient, all bool) error {
	brokers, err := api.BrokerList()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(false)
	if all {
		table.SetHeader([]string{"ID", "ACTIVE", "STATE", "CPUS", "MEM", "HEAP", "HOSTNAME", "ENDPOINT"})
		for _, broker := range brokers {
			table.Append([]string{
				broker.ID,
				broker.Active,
				broker.State,
				fmt.Sprintf("%g", broker.Cpus),
				fmt.Sprintf("%g", broker.Mem),
				fmt.Sprintf("%g", broker.Heap),
				broker.Task.Hostname,
				broker.Task.Endpoint,
			})
		}
	} else {
		table.SetHeader([]string{"ID", "STATE", "CPUS", "MEM", "HEAP", "HOSTNAME", "ENDPOINT"})
		for _, broker := range brokers {
			if broker.Active != true {
				continue
			}

			table.Append([]string{
				broker.ID,
				broker.State,
				fmt.Sprintf("%g", broker.Cpus),
				fmt.Sprintf("%g", broker.Mem),
				fmt.Sprintf("%g", broker.Heap),
				broker.Task.Hostname,
				broker.Task.Endpoint,
			})
		}
	}

	table.Render()

	return nil
}
