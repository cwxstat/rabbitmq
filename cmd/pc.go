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
	"time"

	"github.com/cwxstat/rabbitmq/lib/consumer"
	"github.com/cwxstat/rabbitmq/lib/flag"
	"github.com/cwxstat/rabbitmq/lib/producer"
	"github.com/cwxstat/rabbitmq/lib/setup"
	"github.com/spf13/cobra"
)

// pcCmd represents the pc command
var pcCmd = &cobra.Command{
	Use:   "pc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

pc

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		queue := "queue"
		if len(args) >= 1 {
			queue = args[0]
		}

		setup := setup.NewSetup().CertPath("./etc/certs")
		f := flag.NewFlags()
		handler := &consumer.HS{}

		c, err := consumer.NewConsumer(f.Exchange,
			f.ExchangeType, queue,
			f.BindingKey, f.ConsumerTag, handler, setup.SetupConn)

		fmt.Println("pc called", err)
		fmt.Println("Exchange Type: ", f.ExchangeType)
		forever := make(chan bool)

		for i := 0; i < 100; i++ {
			producer.NewPublish(f.Exchange, f.ExchangeType,
				"test-key", fmt.Sprintf("        message body: %d", i), true, producer.ConfirmOne,
				setup.SetupConn)
			time.Sleep(1 * time.Second)

		}
		<-forever
		c.Shutdown()
	},
}

func init() {
	rootCmd.AddCommand(pcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
