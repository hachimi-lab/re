package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Encoder struct {
	internal zapcore.Encoder
	cores    []zapcore.Core
	conf     *zap.Config
}

func New(opts ...Option) *Encoder {
	config := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "ts",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stack",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.DateTime))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	for _, opt := range opts {
		opt(config)
	}

	if len(config.Encoding) == 0 {
		if config.Development {
			config.Encoding = "console"
		} else {
			config.Encoding = "json"
		}
	}

	encoder := &Encoder{conf: config}
	switch config.Encoding {
	case "console":
		encoder.internal = zapcore.NewConsoleEncoder(config.EncoderConfig)
	case "json":
		encoder.internal = zapcore.NewJSONEncoder(config.EncoderConfig)
	default:
		panic("unknown encoding")
	}

	return encoder
}

func (slf *Encoder) Store(config *lumberjack.Logger, levelEnabler ...LevelEnabler) *Encoder {
	var enabler LevelEnabler
	enabler = slf.conf.Level
	if len(levelEnabler) > 0 {
		enabler = levelEnabler[0]
	}
	slf.cores = append(slf.cores,
		zapcore.NewCore(slf.internal, NewMultiWriteSyncer(AddSync(config)), enabler),
	)
	return slf
}

func (slf *Encoder) Extend(writeSyncer WriteSyncer, levelEnabler ...LevelEnabler) *Encoder {
	var enabler LevelEnabler
	enabler = slf.conf.Level
	if len(levelEnabler) > 0 {
		enabler = levelEnabler[0]
	}
	slf.cores = append(slf.cores, zapcore.NewCore(slf.internal, writeSyncer, enabler))
	return slf
}

func (slf *Encoder) Build(options ...Builder) *Log {
	ins, err := slf.conf.Build()
	if err != nil {
		panic(err)
	}
	options = append([]zap.Option{BuildWithAddCaller(), BuildWithAddCallerSkip(2)}, options...)
	options = append(options, BuildWithWrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(append(slf.cores, core)...)
	}))
	ins = ins.WithOptions(options...)
	return &Log{
		zap:   ins,
		sugar: ins.Sugar(),
	}
}
