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
	"github.com/cuisongliu/container-install/pkg"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall container and delete lib files",
	Run: func(cmd *cobra.Command, args []string) {
		install := install.NewInstaller()
		install.UnInstall()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
	uninstallCmd.Flags().StringVar(&pkg.SSHConfig.User, "user", "root", "servers user name for ssh")
	uninstallCmd.Flags().StringVar(&pkg.SSHConfig.Password, "passwd", "admin", "servers user password for ssh")
	uninstallCmd.Flags().StringVar(&pkg.SSHConfig.PkFile, "pk", "", "servers user private key file for ssh")
	uninstallCmd.Flags().StringVar(&pkg.SSHConfig.PkPassword, "pk-passwd", "", "private key password for ssh")

	uninstallCmd.Flags().StringSliceVar(&install.Hosts, "host", []string{}, "install hosts")
	uninstallCmd.Flags().StringVar(&command.Lib, "lib", "", "store location,default : docker is /var/lib/docker , containerd is /var/lib/containerd ")
	// Here you will define your flags and configuration settings.
	//uninstallCmd.Flags().StringVar(&install.User, "user", "root", "servers user name for ssh")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
