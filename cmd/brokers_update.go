// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/eddyzags/kafkactl/api/client"
	"github.com/eddyzags/kafkactl/brokers"
	"github.com/eddyzags/kafkactl/types"

	"github.com/spf13/cobra"
)

var brokersUpdateParams types.BrokerAdd

var brokersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a broker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "The broker id or expr is missing\n")
			os.Exit(1)
		}

		if apiURL == "" {
			fmt.Println("Cannot find kafka scheduler url. Using default: 0.0.0.0:7070")
		}

		c := client.NewClient(apiURL)

		brokersUpdateParams = args[0]

		if err := brokers.Update(c, &brokersUpdateParams); err != nil {
			fmt.Fprintf(os.Stderr, "kafkactl: Unexpected error occured \"%v\"\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	brokersCmd.AddCommand(brokersUpdateCmd)

	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.BindAddress, "bind-address", "", "", "Broker bind address (broker0, 192.168.50.*, if:eth1). Default - auto")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Constraints, "constraints", "", "", "Constraints (hostname=like:master,rack=like:1.*)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Cpus, "cpus", "", "0.5", "Cpus amount (0.1, 1, 2)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.FailoverDelay, "failover-delay", "", "", "Failover delay")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.FailoverMaxDelay, "failover-max-delay", "", "", "Max failover delay")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.FailoverMaxTries, "failover-max-tries", "", "", "Max failover tries. Default - none")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Heap, "heap", "", "1024", "Heap amount in MB")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.JvmOptions, "jvm-options", "", "", "JVM options string (-Xms128m -XX:PermSize=48m)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Log4jOptions, "log4j-options", "", "", "log4j options or file. Examples: log4j.logger.kafka=DEBUG, kafkaAppender file:log4j.properties")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Mem, "mem", "", "2048", "Mem amount in MB. Default - 2048")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Options, "options", "", "", "Options or file. Examples: log.dirs=/tmp/kafka/$id,num.io.threads=16 file:server.properties")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Port, "port", "", "", "Port or range (31092, 31090..31100). Default - auto")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.StickinessPeriod, "stickiness-period", "", "", "Stickiness period to preserve same node for broker (5m, 10m, 1h)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersUpdateParams.Volume, "volume", "", "", "Pre-reserved persitent volume id")
}
