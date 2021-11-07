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

	"github.com/cwxstat/rabbitmq/wrapper/gzencode"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: " upload <directory to gz> <certurl>",
	Long: `
	

	 upload /workspaces/rabbitmq/lib "amqps://pig:P033wor4@localhost:5671"
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")

		// FIXME: (mmc) this smells bad
		if len(args) < 1 {
			fmt.Printf("\n: Need to provide the following:\n")
			fmt.Printf("\n:  1) directory to gz and encode\n")
			return
		}

		dir := args[0]

		g := gzencode.NewGZ()
		g.CertPath("./etc/certs")
		g.DirIn(dir)

		err := g.Produce()
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
