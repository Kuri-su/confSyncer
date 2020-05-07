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
	"fmt"

	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull",
	Long:  `pull`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ConfigPull()
		if err != nil {
			color.Red(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ConfigPull() error {
	if !unit.IsDir(TmpDirPath) {
		err := unit.GitClone(viper.GetString("gitRepo"), TmpDirPath)
		if err != nil {
			return err
		}
	} else {
		err := unit.GitPull(TmpDirPath)
		if err != nil {
			return err
		}
	}

	c := viper.Get("maps")
	if c == nil {
		return nil
	}

	str, err := jsoniter.MarshalToString(c)
	if err != nil {
		return err
	}

	var maps []Path
	err = jsoniter.UnmarshalFromString(str, &maps)
	if err != nil {
		return err
	}

	for _, copyMap := range maps {
		unit.MakeDirWithFilePath(copyMap.Dist)
		err = unit.Copy(TmpDirPath+copyMap.Src, copyMap.Dist)
		if err != nil {
			return err
		}
	}

	fmt.Println("Configs pull finish!")
	return nil
}
