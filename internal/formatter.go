package internal

import (
	"fmt"
	"math"
	"time"
)

func FormatTotal(time int) string {
	hours := time / 3600
	minutes := (time % 3600) / 60
	seconds := time % 60
	totalString := ""
	if hours > 0 {
		totalString += fmt.Sprintf("%dh", hours)
		if minutes > 0 {
			totalString += " "
		}
	}

	if minutes > 0 {
		totalString += fmt.Sprintf("%dm", minutes)
		if minutes > 0 {
			totalString += " "
		}
	}

	if seconds > 0 {
		totalString += fmt.Sprintf("%ds", seconds)
	}

	return totalString
}

func FormatTime(timeStamp int64, format string) string {
	return time.Unix(timeStamp, 0).UTC().Format(format)
}

func Round(time, roundBy int) int {
	if time < roundBy/2 {
		return roundBy
	}

	return int(float64(roundBy) * math.Round(float64(time)/float64(roundBy)))
}
