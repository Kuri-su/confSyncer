/*
 * 	ConfSyncer - a little sync config files tool in the Linux.
 *     Copyright (C) 2020  Amatist_kurisu<misaki.zhcy@gmail.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package main

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
