package brokers

import (
	"fmt"
	"os"

	"github.com/eddyzags/kafkactl/api/client"
	"github.com/eddyzags/kafkactl/types"

	"github.com/olekukonko/tablewriter"
)

func Update(api client.APIClient, params *types.BrokerAdd) error {
	in := map[string]string{
		"bindAddress":      params.BindAddress,
		"constraints":      params.Constraints,
		"cpus":             params.Cpus,
		"broker":           params.Expr,
		"failoverDelay":    params.FailoverDelay,
		"failoverMaxTries": params.FailoverMaxTries,
		"heap":             params.Heap,
		"jvmOptions":       params.JvmOptions,
		"log4jOptions":     params.Log4jOptions,
		"mem":              params.Mem,
		"options":          params.Options,
		"port":             params.Port,
		"stickinessPeriod": params.StickinessPeriod,
		"volume":           params.Volume,
	}

	brokers, err := api.BrokerUpdate(in)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"ID", "ACTIVE", "CPUS", "MEM", "HEAP", "PORT"})

	for _, broker := range brokers {
		table.Append([]string{
			broker.ID,
			fmt.Sprintf("%t", broker.Active),
			fmt.Sprintf("%g", broker.Cpus),
			fmt.Sprintf("%g", broker.Mem),
			fmt.Sprintf("%g", broker.Heap),
			broker.Port,
		})
	}

	table.Render()

	return nil
}
