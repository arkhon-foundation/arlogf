package arlogf

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func getdate() string {
	now := time.Now()

	hour, minute, second := now.Clock()
	year, month, day := now.Date()

	date := fmt.Sprintf("%d/%d/%d", month, day, year)
	clock := fmt.Sprintf("%d:%02d:%02d", hour, minute, second)

	return color.New(color.FgHiBlack).Sprintf("[%s | %s]", date, clock)
}

func filedate() string {
	now := time.Now()
	year, month, day := now.Date()

	return fmt.Sprintf("log-%d_%d_%d.txt", month, day, year)
}
