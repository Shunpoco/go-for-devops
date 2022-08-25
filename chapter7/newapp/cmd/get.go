/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := cmd.Flags()
		fmt.Println(fs)
		addr := mustString(fs, "addr")
		fmt.Println(addr)
	},
}

func mustString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}

	return v
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("dev", "d", false, "Uses the dev server instead of prod")
	getCmd.Flags().String("addr", "127.0.0.1:80", "Set the server to use, defaults to production")
	getCmd.Flags().StringP("author", "a", "", "specify the author to get a quote for")
	getCmd.Flags().Bool("json", false, "Output is in JSON format")
}
