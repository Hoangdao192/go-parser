package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Get author of go parser",
	Long:  "Get author of go parser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name: Nguyen Dang Hoang Dao")
		fmt.Println("Email: 200202390@vnu.edu.vn")
	},
}

func init() {
	rootCmd.AddCommand(authorCmd)
}
