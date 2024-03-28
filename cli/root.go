package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"joern-go/core/parser"
	"os"
	"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "parse",
	Short: "Go parser is a parsing tool for go code",
	Long:  "Go parser is a parsing tool for go code",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if Output == "" {
			Output = currentDirectory()
		}
		Ignore, _ := cmd.Flags().GetStringSlice("ignore")
		parser.Parse(args[1], Output, Ignore)
	},
}

type stringArray []string

func (i *stringArray) Type() string {
	return "stringArray"
}

func (i *stringArray) String() string {

	return strings.Join(*i, ", ")
}

func (i *stringArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var Output string
var Ignore []string

func init() {
	Ignore = []string{}
	rootCmd.Flags().StringVarP(&Output, "output", "o", "", "Output")
	rootCmd.Flags().StringSliceP("ignore", "i", []string{}, "Ignore file or directory")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func currentDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
