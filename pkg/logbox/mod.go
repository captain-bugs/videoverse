package logbox

import (
	"flag"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"videoverse/pkg/config"
)

var instance *zerolog.Logger = nil
var once sync.Once

func NewLogBox() *zerolog.Logger {
	once.Do(func() {
		debug := flag.Bool("debug", false, "sets log level to debug")
		flag.Parse()

		logfile, err := os.OpenFile(config.LOGGING_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimestampFieldName = "timestamp"
		if *debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			return filepath.Base(file) + ":" + strconv.Itoa(line)
		}
		logger := zerolog.New(io.MultiWriter(consoleWriter, logfile)).With().Caller().Timestamp().Logger()
		if err != nil {
			logger.Error().Err(err).Msg("error craeting a temp file four our log")
		}
		instance = &logger
	})
	return instance
}
