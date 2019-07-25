// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

const urlPrefix = "https://download.docker.com/linux/static/stable/x86_64/docker-%s.tgz"

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use: "print",
	Run: func(cmd *cobra.Command, args []string) {
		versions := []string{
			"docker-17.03.0-ce.tgz",
			"docker-17.03.1-ce.tgz",
			"docker-17.03.2-ce.tgz",
			"docker-17.06.0-ce.tgz",
			"docker-17.06.1-ce.tgz",
			"docker-17.06.2-ce.tgz",
			"docker-17.09.0-ce.tgz",
			"docker-17.09.1-ce.tgz",
			"docker-17.12.0-ce.tgz",
			"docker-17.12.1-ce.tgz",
			"docker-18.03.0-ce.tgz",
			"docker-18.03.1-ce.tgz",
			"docker-18.06.0-ce.tgz",
			"docker-18.06.1-ce.tgz",
			"docker-18.06.2-ce.tgz",
			"docker-18.06.3-ce.tgz",
			"docker-18.09.0.tgz",
			"docker-18.09.1.tgz",
			"docker-18.09.2.tgz",
			"docker-18.09.3.tgz",
			"docker-18.09.4.tgz",
			"docker-18.09.5.tgz",
			"docker-18.09.6.tgz",
			"docker-18.09.7.tgz",
			"docker-18.09.8.tgz",
			"docker-19.03.0.tgz",
		}

		for _, v := range versions {
			println(fmt.Sprintf(urlPrefix, v))
		}

	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
