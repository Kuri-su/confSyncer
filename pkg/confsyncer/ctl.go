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
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	// Root
	dirPath = "$HOME/.confSyncer"
	cfgFile = "$HOME/.confSyncer/config.yaml"
	// TmpDir(TmpDirPath) is a tmp git clone dir
	TmpDirPath = "/tmp/confSyncer-" + fmt.Sprint(time.Now().Format("20060102"))
	// version
	version = "0.0.1"
	// use it when need a default config
	DefaultConfigContext = `
---
gitRepo: git@gitlab.com:xxx/xxx.git
gitPullTimeInternal: 30 # second
configs:
  - src: /a.json
    dist: /home/kurisu/.config/a/config
  - src: /b.json
    dist: /home/kurisu/.config/b/config
`

	// ========================================================================
	// ========================================================================
	// ========================================================================

	// CMD
	rootCmd = &cobra.Command{
		Use:   "confSyncer",
		Short: "confSyncer",
		Long:  `confSyncer`,
	}
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "show config",
		Run:   ShowConfig,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run:   Version,
	}
	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "push",
		Run:   ConfigPush,
	}
	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "pull",
		Run:   ConfigPull,
	}
	deamonPullCmd = &cobra.Command{
		Use:   "deamon",
		Short: "deamon",
		Run:   DaemonPull,
	}
)

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	cfgFile = strings.Replace(cfgFile, "$HOME", u.HomeDir, 1)
	dirPath = strings.Replace(dirPath, "$HOME", u.HomeDir, 1)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// register commands
	rootCmd.AddCommand(
		configCmd,
		versionCmd,
		pushCmd,
		pullCmd,
		deamonPullCmd,
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Version
func Version(cmd *cobra.Command, args []string) {
	color.Set(color.Bold)
	color.HiWhite("confSyncer version: %s", version)
}
