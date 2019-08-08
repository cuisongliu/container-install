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
	"github.com/cuisongliu/container-install/install"
	"github.com/cuisongliu/container-install/install/command"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install docker for url",
	Run: func(cmd *cobra.Command, args []string) {
		install := install.NewInstaller()
		install.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.
	installCmd.Flags().StringVar(&command.User, "user", "root", "servers user name for ssh")
	installCmd.Flags().StringVar(&command.Passwd, "passwd", "admin", "servers user password for ssh")
	installCmd.Flags().StringVar(&command.PrivateKeyFile, "pk", "/root/.ssh/id_rsa", "servers user private key file for ssh")

	installCmd.Flags().StringSliceVar(&install.Hosts, "host", []string{}, "container install hosts")
	installCmd.Flags().StringSliceVar(&command.RegistryArr, "registry", []string{"127.0.0.1"}, "container's registry ip")
	installCmd.Flags().StringVar(&command.PkgUrl, "pkg-url", "", "https://download.docker.com/linux/static/stable/x86_64/docker-19.03.0.tgz download offline docker url, or file localtion ex. /root/docker.tgz")
	installCmd.Flags().StringVar(&command.Lib, "lib", "", "store location,default : docker is /var/lib/docker , containerd is /var/lib/containerd ")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
