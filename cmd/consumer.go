/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/cwxstat/rabbitmq/lib/consumer"
	"github.com/cwxstat/rabbitmq/lib/flag"
	"github.com/cwxstat/rabbitmq/lib/setup"

	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

consumer

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumer called")
		queue := "queue"
		if len(args) >= 1 {
			queue = args[0]
		}

		setup := setup.NewSetup().CertPath("./etc/certs")
		f := flag.NewFlags()
		handler := &consumer.HS{}
		fmt.Println("exchangeType: ", f.ExchangeType)

		c, err := consumer.NewConsumer(f.Exchange,
			f.ExchangeType, queue,
			f.BindingKey, f.ConsumerTag, handler, setup.SetupConn)

		fmt.Println("pc called", err)
		forever := make(chan bool)

		<-forever
		c.Shutdown()
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
