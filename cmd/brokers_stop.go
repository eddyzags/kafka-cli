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

	"github.com/spf13/cobra"
)

var (
	brokersStopForce   bool
	brokersStopTimeout string
)

var brokersStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a broker",
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

		if err := brokers.Stop(c, args[0], brokersStopTimeout, brokersStopForce); err != nil {
			fmt.Fprintf(os.Stderr, "kafkactl: Unexpected error occured \"%v\"\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	brokersCmd.AddCommand(brokersStopCmd)

	brokersStopCmd.PersistentFlags().BoolVarP(&brokersStopForce, "force", "", false, "Forcibly stop")
	brokersStopCmd.PersistentFlags().StringVarP(&brokersStopTimeout, "timeout", "", "", "Timeout 30s, 1m, 1h). 0s - no timeout")
}
