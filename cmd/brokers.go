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

// brokersCmd represents the brokers command
var brokersCmd = &cobra.Command{
	Use:   "brokers",
	Short: "Manage your brokers",
	Long: `A broker holds a number of partitions from differents topics.
You can manager them with this command`,
	Run: func(cmd *cobra.Command, args []string) {
		if apiURL == "" {
			fmt.Println("Cannot find kafka scheduler url. Using default: 0.0.0.0:7070")
		}

		c := client.NewClient(apiURL)

		if err := brokers.List(c, brokersListAll); err != nil {
			fmt.Fprintf(os.Stderr, "kafkactl: Unexpected error occured \"%v\"\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(brokersCmd)

	brokersCmd.PersistentFlags().BoolVarP(&brokersListAll, "all", "a", false, "Do not ignore unactive brokers")
}
