/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ashsobeck/tico/parser"
	"github.com/spf13/cobra"
)

// longdatetimeCmd represents the longdatetime command
var longdatetimeCmd = &cobra.Command{
	Use:   "longdatetime",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("longdatetime called")
		ticks, err := parser.ParseDate(args)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dateStr := parser.Convert(ticks, parser.LongDateTime)
		fmt.Println(dateStr)
	},
}

func init() {
	rootCmd.AddCommand(longdatetimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// longdatetimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// longdatetimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
