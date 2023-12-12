package log

import (
	"io"
	"os"

	"go.uber.org/zap/zapcore"
)

type WriteSyncer = zapcore.WriteSyncer

func NewMultiWriteSyncer(ws ...WriteSyncer) WriteSyncer {
	return zapcore.NewMultiWriteSyncer(ws...)
}

func AddSync(w io.Writer) WriteSyncer {
	return zapcore.AddSync(w)
}

func StdoutWriteSyncer() WriteSyncer {
	return zapcore.AddSync(io.Writer(os.Stdout))
}
