package cmd

import (
	"fmt"
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
