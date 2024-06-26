package arlogf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/fatih/color"
)

var LastLog string
var DumpLogsFolder string = ""

type Logger struct{}

type LogOptions struct {
	id              string
	message         string
	stackTrace      string
	isFatal         bool
	prefix          string
	prefixFormatter []color.Attribute
}

func NewLogger(disableColors bool) *Logger {
	color.NoColor = disableColors

	return &Logger{}
}

func PrintWithOptions(options *LogOptions) {
	var out []string

	out = append(out,
		getdate(),
		fmt.Sprintf("%s (%s): %s",
			color.New(options.prefixFormatter...).Sprint(options.prefix),
			color.HiCyanString(options.id),
			options.message,
		))

	if options.stackTrace != "" {
		out = append(out,
			color.New(color.FgBlack, color.Italic).Sprintf("-- Begin Trace for %s", options.id),
			color.RedString(options.stackTrace),
			color.New(color.FgBlack, color.Italic).Sprintf("-- End Trace for %s", options.id),
		)
	}

	log := strings.Join(out, "\n")
	LastLog = stripansi.Strip(log)

	if DumpLogsFolder != "" {
		dumplastlog(filepath.Join(DumpLogsFolder, filedate()))
	}

	fmt.Println(log)
}

func (l *Logger) Builder(id string) *LogOptions {
	lo := &LogOptions{id: id}

	*&lo.prefix = "TYPE "
	*&lo.prefixFormatter = []color.Attribute{color.FgHiBlack, color.Bold}

	return lo
}

func (lo *LogOptions) Log() *LogOptions {
	*&lo.prefix = "LOG  "
	*&lo.prefixFormatter = []color.Attribute{color.FgHiGreen, color.Bold}

	return lo
}

func (lo *LogOptions) Warn() *LogOptions {
	*&lo.prefix = "WARN "
	*&lo.prefixFormatter = []color.Attribute{color.FgHiYellow, color.Bold}

	return lo
}

func (lo *LogOptions) Error(stackTrace string) *LogOptions {
	*&lo.prefix = "ERROR"
	*&lo.prefixFormatter = []color.Attribute{color.FgHiRed, color.Bold}
	*&lo.stackTrace = stackTrace

	return lo
}

func (lo *LogOptions) Fatal(stackTrace string) *LogOptions {
	*&lo.prefix = "FATAL"
	*&lo.prefixFormatter = []color.Attribute{color.FgBlack, color.Bold}
	*&lo.stackTrace = stackTrace
	*&lo.isFatal = true

	return lo
}

func (lo *LogOptions) Print(message string) {
	*&lo.message = message

	PrintWithOptions(lo)

	if lo.isFatal {
		os.Exit(1)
	}
}

func (lo *LogOptions) Printf(format string, a ...any) {
	*&lo.message = fmt.Sprintf(format, a...)

	PrintWithOptions(lo)

	if lo.isFatal {
		os.Exit(1)
	}
}

func (l *Logger) Space() {
	fmt.Println()
}

func dumplastlog(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(LastLog + "\n"); err != nil {
		return err
	}

	return nil
}
