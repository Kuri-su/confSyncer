package confSyncer

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  `show version`,
	Run: func(cmd *cobra.Command, args []string) {
		Version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
func Version() {
	color.Set(color.Bold)
	color.HiWhite("confSyncer version: %s", version)
}
