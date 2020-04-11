/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"errors"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push",
	Long:  `push`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ConfigPush()
		if err != nil {
			color.Red(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func ConfigPush() error {
	if !unit.IsDir(TmpDirPath) {
		return errors.New("not found dir")
	} else {
		err := unit.GitPush(TmpDirPath)
		if err != nil {
			return err
		}
	}

	maps := viper.GetStringMapString("maps")

	for src, dist := range maps {
		err := unit.Copy(dist, TmpDirPath+src)
		if err != nil {
			return err
		}
	}

	color.Green("Configs push finish!")
	return nil
}
