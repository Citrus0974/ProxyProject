package logger

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

type Logger interface {
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatalf(msg string, args ...any)
}

type ZerologLogger struct {
	logger zerolog.Logger
}

func NewZerologLogger() Logger {
	asyncWriter := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		log.Printf("Logger dropped %d messages", missed)
	})

	return &ZerologLogger{
		logger: zerolog.New(asyncWriter).With().Timestamp().Logger(),
	}
}

func (l *ZerologLogger) Debugf(msg string, args ...any) {
	l.logger.Debug().Msgf(msg, args...)
}

func (l *ZerologLogger) Infof(msg string, args ...any) {
	l.logger.Info().Msgf(msg, args...)
}

func (l *ZerologLogger) Warnf(msg string, args ...any) {
	l.logger.Warn().Msgf(msg, args...)
}

func (l *ZerologLogger) Errorf(msg string, args ...any) {
	l.logger.Error().Msgf(msg, args...)
}

func (l *ZerologLogger) Fatalf(msg string, args ...any) {
	l.logger.Fatal().Msgf(msg, args...)
}
