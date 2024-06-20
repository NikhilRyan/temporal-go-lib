package logging

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// InitLogger initializes the zap logger
func InitLogger() (*zap.Logger, error) {
    config := zap.Config{
        Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
        Development: false,
        Encoding:    "json",
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "time",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            StacktraceKey:  "stacktrace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.CapitalLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeDuration: zapcore.StringDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
    }
    return config.Build()
}

// InitDevelopmentLogger initializes the zap logger for development
func InitDevelopmentLogger() (*zap.Logger, error) {
    config := zap.Config{
        Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
        Development: true,
        Encoding:    "console",
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "time",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            StacktraceKey:  "stacktrace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.CapitalColorLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeDuration: zapcore.StringDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
    }
    return config.Build()
}
