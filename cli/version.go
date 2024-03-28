package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

const VERSION = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of go parser",
	Long:  "Show version of go parser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go parser version %s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
