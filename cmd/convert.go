/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert called")
		date := ""
		if _, err := time.Parse(time.RFC3339, args[0]+"T00:00:00Z"); err == nil {
			date += args[0]
		} else {
			fmt.Println(err.Error())
			fmt.Println("[ERROR]: wrong time arg format. Need 'YYYY-MM-DD'")
			return
		}

		_, offset := time.Now().Zone()
		off := convertOffset(offset)
		if len(args) >= 2 && args[1] != "" {
			date += "T" + args[1] + ":00" + off
		} else {
			fmt.Println(offset)
			date += "T00:00:00" + off
		}
		ticks, err := time.Parse(time.RFC3339, date)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("[ERROR]: wrong time arg format. Need 'HH:MM' (24 hour)")
			return
		}
		fmt.Println("Time:", date)
		str := Convert(ticks.Unix(), Default)
		fmt.Println(str)
	},
}

func convertOffset(offset int) string {
	var symbol string
	if offset >= 0 {
		symbol = "+"
	} else {
		symbol = "-"
	}
	offset /= 3600
	offset = int(math.Abs(float64(offset)))

	if offset < 10 {
		symbol += fmt.Sprintf("0%d:00", offset)
	} else {
		symbol += fmt.Sprintf("%d:00", offset)
	}

	return symbol
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
