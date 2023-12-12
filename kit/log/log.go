package log

import (
	"go.uber.org/zap"
)

type Log struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func (slf *Log) Named(name string) *Log {
	slf.zap = slf.zap.Named(name)
	slf.sugar = slf.sugar.Named(name)
	return slf
}

func (slf *Log) debug(msg string, fields ...Field) {
	slf.zap.Debug(msg, fields...)
}

func (slf *Log) info(msg string, fields ...Field) {
	slf.zap.Info(msg, fields...)
}

func (slf *Log) warn(msg string, fields ...Field) {
	slf.zap.Warn(msg, fields...)
}

func (slf *Log) error(msg string, fields ...Field) {
	slf.zap.Error(msg, fields...)
}

func (slf *Log) dPanic(msg string, fields ...Field) {
	slf.zap.DPanic(msg, fields...)
}

func (slf *Log) panic(msg string, fields ...Field) {
	slf.zap.Panic(msg, fields...)
}

func (slf *Log) fatal(msg string, fields ...Field) {
	slf.zap.Fatal(msg, fields...)
}

func (slf *Log) debugf(format string, args ...any) {
	slf.sugar.Debugf(format, args...)
}

func (slf *Log) infof(format string, args ...any) {
	slf.sugar.Infof(format, args...)
}

func (slf *Log) warnf(format string, args ...any) {
	slf.sugar.Warnf(format, args...)
}

func (slf *Log) errorf(format string, args ...any) {
	slf.sugar.Errorf(format, args...)
}

func (slf *Log) dPanicf(format string, args ...any) {
	slf.sugar.DPanicf(format, args...)
}

func (slf *Log) panicf(format string, args ...any) {
	slf.sugar.Panicf(format, args...)
}

func (slf *Log) fatalf(format string, args ...any) {
	slf.sugar.Fatalf(format, args...)
}

func (slf *Log) Debug(msg string, fields ...Field) {
	slf.debug(msg, fields...)
}

func (slf *Log) Info(msg string, fields ...Field) {

	slf.info(msg, fields...)
}

func (slf *Log) Warn(msg string, fields ...Field) {
	slf.warn(msg, fields...)
}

func (slf *Log) Error(msg string, fields ...Field) {
	slf.error(msg, fields...)
}

// DPanic 仅在开发模式下Panic
func (slf *Log) DPanic(msg string, fields ...Field) {
	slf.dPanic(msg, fields...)
}

func (slf *Log) Panic(msg string, fields ...Field) {
	slf.panic(msg, fields...)
}

func (slf *Log) Fatal(msg string, fields ...Field) {
	slf.fatal(msg, fields...)
}

func (slf *Log) Debugf(format string, args ...any) {
	slf.debugf(format, args...)
}

func (slf *Log) Infof(format string, args ...any) {
	slf.infof(format, args...)
}

func (slf *Log) Warnf(format string, args ...any) {
	slf.warnf(format, args...)
}

func (slf *Log) Errorf(format string, args ...any) {
	slf.errorf(format, args...)
}

func (slf *Log) DPanicf(format string, args ...any) {
	slf.dPanicf(format, args...)
}

func (slf *Log) Panicf(format string, args ...any) {
	slf.panicf(format, args...)
}

func (slf *Log) Fatalf(format string, args ...any) {
	slf.fatalf(format, args...)
}
