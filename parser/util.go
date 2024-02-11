package parser

import (
	"fmt"
	"math"
	"time"
)

type Format string

const (
	Default       Format = ""
	Short         Format = "t"
	Long          Format = "T"
	ShortDate     Format = "d"
	LongDate      Format = "D"
	ShortDateTime Format = "f"
	LongDateTime  Format = "F"
	Relative      Format = "R"
)

func Convert(date int64, format Format) string {
	switch format {
	case Default:
		return fmt.Sprintf("<t:%d>", date)
	default:
		return fmt.Sprintf("<t:%d:%s>", date, format)
	}
}

func ConvertOffset(offset int) string {
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

func ParseDate(args []string) (int64, error) {
	date := ""
	if _, err := time.Parse(time.RFC3339, args[0]+"T00:00:00Z"); err == nil {
		date += args[0]
	} else {
		fmt.Println(err.Error())
		fmt.Println("[ERROR]: wrong time arg format. Need 'YYYY-MM-DD'")
		return 0, err
	}

	_, offset := time.Now().Zone()
	off := ConvertOffset(offset)
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
		return 0, err
	}

	fmt.Println("Time:", date)
	return ticks.Unix(), nil
}
