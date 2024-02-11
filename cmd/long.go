/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ashsobeck/tico/parser"
	"github.com/spf13/cobra"
)

// longCmd represents the long command
var longCmd = &cobra.Command{
	Use:   "long",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("long called")
		ticks, err := parser.ParseDate(args)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dateStr := parser.Convert(ticks, parser.Long)

		fmt.Println(dateStr)
	},
}

func init() {
	rootCmd.AddCommand(longCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// longCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// longCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
