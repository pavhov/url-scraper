package log_lib

import (
	"fmt"
	"os"
	"service/src/internal/builtin_lib"
	"service/src/tools/config_lib"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var levels = map[string]zerolog.Level{
	"info":     zerolog.InfoLevel,
	"debug":    zerolog.DebugLevel,
	"error":    zerolog.ErrorLevel,
	"warn":     zerolog.WarnLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"no":       zerolog.NoLevel,
	"disabled": zerolog.Disabled,
	"trace":    zerolog.TraceLevel,
}

type Interface interface {
	Init() (err error)
	setLogLevel() (err error)
}

type Log struct {
}

func NewLog() *Log {
	var instance Log
	return &instance
}

func (l Log) Init() (err error) {
	defer builtin_lib.Recovery()
	if err = l.setLogLevel(); err != nil {
		return
	}
	return
}

func (l Log) setLogLevel() (err error) {

	var level zerolog.Level
	logLevel := config_lib.Config.Get("app_log").(string)
	level = levels[logLevel]
	wd, _ := os.Getwd()

	format := func(i interface{}) string { return fmt.Sprintf("[%s]", i) }
	emptyFormat := func(i interface{}) string { return "" }
	beforeFormat := func(i interface{}) string { return fmt.Sprintf("[%s: ", i) }
	afterFormat := func(i interface{}) string { return fmt.Sprintf("%s]", i) }
	callerFormat := func(i interface{}) string {
		return strings.Replace(fmt.Sprintf("[%s]", i), strings.Replace(wd, "\\", "/", -1), "", 1)
	}

	output := zerolog.NewConsoleWriter()
	output.TimeFormat = time.Stamp
	output.FormatTimestamp = emptyFormat
	output.FormatLevel = format
	output.FormatLevel = format
	output.FormatCaller = format
	output.FormatMessage = format
	output.FormatFieldName = beforeFormat
	output.FormatFieldValue = afterFormat
	output.FormatErrFieldName = beforeFormat
	output.FormatErrFieldValue = afterFormat
	output.FormatCaller = callerFormat
	output.NoColor = true

	zerolog.SetGlobalLevel(level)
	zerolog.DisableSampling(true)

	log.Logger = zerolog.New(output).With().Caller().Logger()
	return
}
