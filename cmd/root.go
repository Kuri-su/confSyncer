package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

type Config struct {
	GitRepo string `yaml:"gitRepo"`
	GitPull struct {
		TimeInternal int `yaml:"timeInternal"`
	} `yaml:"gitPull"`
	Configs []Path `yaml:"configs"`
}

type Path struct {
	Src  string `yaml:"src"`
	Dist string `yaml:"dist"`
}

var (
	dirPath              = "$HOME/.confSyncer"
	cfgFile              = "$HOME/.confSyncer/config.yaml"
	TmpDirPath           = "/tmp/confSyncer-" + fmt.Sprint(time.Now().Format("20060102"))
	version              = "0.0.1"
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
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "confSyncer",
	Short: "confSyncer",
	Long:  `confSyncer`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
