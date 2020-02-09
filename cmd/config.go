/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure atcoder-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		var username, password, language string
		fmt.Printf("username: ")
		fmt.Scanf("%s", &username)
		fmt.Printf("password: ")
		fmt.Scanf("%s", &password)
		fmt.Printf("language (list supported languages by `atcoder languages`): ")
		fmt.Scanf("%s", &language)

		viper.Set("username", username)
		viper.Set("password", password)
		viper.Set("language", language)

		dir, err := homedir.Dir()
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(dir, confName), os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
		file.Close()

		return viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
