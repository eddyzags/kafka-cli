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

var brokersAddParams types.BrokerAdd

// brokersAddCmd represents a brokersCmd subcommand
var brokersAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a broker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "The broker id or expr is missing")
			os.Exit(1)
		}

		if apiURL == "" {
			fmt.Println("Cannot find kafka scheduler url. Using default: 0.0.0.0:7070")
		}

		c := client.NewClient(apiURL)

		brokersAddParams.Expr = args[0]

		if err := brokers.Add(c, &brokersAddParams); err != nil {
			fmt.Fprintf(os.Stderr, "kafkactl: Unexpected error occured \"%v\"\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	brokersCmd.AddCommand(brokersAddCmd)

	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.BindAddress, "bind-address", "", "", "Broker bind address (broker0, 192.168.50.*, if:eth1). Default - auto")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Constraints, "constraints", "", "", "Constraints (hostname=like:master,rack=like:1.*)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Cpus, "cpus", "", "0.5", "Cpus amount (0.1, 1, 2)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.FailoverDelay, "failover-delay", "", "", "Failover delay")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.FailoverMaxDelay, "failover-max-delay", "", "", "Max failover delay")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.FailoverMaxTries, "failover-max-tries", "", "", "Max failover tries. Default - none")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Heap, "heap", "", "1024", "Heap amount in MB")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.JvmOptions, "jvm-options", "", "", "JVM options string (-Xms128m -XX:PermSize=48m)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Log4jOptions, "log4j-options", "", "", "log4j options or file. Examples: log4j.logger.kafka=DEBUG, kafkaAppender file:log4j.properties")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Mem, "mem", "", "2048", "Mem amount in MB. Default - 2048")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Options, "options", "", "", "Options or file. Examples: log.dirs=/tmp/kafka/$id,num.io.threads=16 file:server.properties")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Port, "port", "", "", "Port or range (31092, 31090..31100). Default - auto")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.StickinessPeriod, "stickiness-period", "", "", "Stickiness period to preserve same node for broker (5m, 10m, 1h)")
	brokersAddCmd.PersistentFlags().StringVarP(&brokersAddParams.Volume, "volume", "", "", "Pre-reserved persitent volume id")
}
