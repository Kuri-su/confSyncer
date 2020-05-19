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

package confsyncer

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".confSyncer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".confSyncer")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("yaml")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		color.Green("Using config file: %s", viper.ConfigFileUsed())
	} else {
		color.Yellow("Warning: Not found config, now reinit it! ")
		err := createConfigFile()
		if err != nil {
			log.Fatal(color.RedString(err.Error()))
		}
	}

	// found config file failed
	err := configCheck()
	if err != nil {
		log.Fatal(color.RedString(err.Error()))
	}
}

func configCheck() error {
	if viper.Get("maps") == nil || len(viper.Get("maps").([]interface{})) == 0 {
		color.Red("Warning! 'maps' fields is empty in configFile")
	}
	return nil
}

func createConfigFile() error {
	if !unit.IsFile(cfgFile) {

		// make dir
		err := os.MkdirAll(dirPath, 0644)
		if err != nil {
			log.Fatalln(err)
		}

		newFile, err := os.Create(cfgFile)
		if err != nil {
			log.Fatalln(err)
		}
		// Write config
		_, err = newFile.Write([]byte(DefaultConfigContext))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\x1b[32mSuccess! Create config file in %s \x1b[0m \n", cfgFile)
	}

	return nil
}

// ShowConfig
func ShowConfig(cmd *cobra.Command, args []string) {
	settingMap := viper.AllSettings()
	settingStr, err := jsoniter.MarshalIndent(settingMap, "", "    ")
	if err != nil {
		fmt.Printf("json marshal failed in ShowConfig! err: %s \n", err.Error())
	}
	color.Green("\nThis is your config: \n\n%s \n", settingStr)
}