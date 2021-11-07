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

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "  download /tmp/junk3 qa lib.tar.gz",
	Long: `
	
	      upload /workspaces/rabbitmq/lib 

	This will produce  "lib.tar.gz"
	
	      download /tmp/junk3 qa /tmp/lib.tar.gz
	


	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")

		if len(args) < 3 {
			fmt.Printf("\n: Need to provide the following:\n")
			fmt.Printf("\n:  1) directory to gz and dencode\n")
			fmt.Printf("\n:  2) consumer Q\n")
			fmt.Printf("\n:  3) handleFile\n")
			return
		}

		dir := args[0]

		g := gzencode.NewGZ()
		g.CertPath("./etc/certs")

		g.DestDir(dir)
		g.HandleFile(args[2], dir)
		g.ConsumerQ(args[1])

		err := g.Consume()
		if err != nil {
			fmt.Println(err)
			return
		}

		forever := make(chan bool)

		<-forever
		g.Shutdown()

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
