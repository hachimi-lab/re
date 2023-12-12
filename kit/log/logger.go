package log

var logger *Log

func init() {
	logger = New().Build()
}

func SetLogger(log *Log) {
	if log == nil {
		return
	}
	_ = log.zap.Sync()
	_ = log.sugar.Sync()
	logger = log
}

func GetLogger() *Log {
	return logger
}

func Debug(msg string, fields ...Field) {
	logger.debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.error(msg, fields...)
}

func DPanic(msg string, fields ...Field) {
	logger.dPanic(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	logger.panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger.fatal(msg, fields...)
}

func Debugf(format string, args ...any) {
	logger.debugf(format, args...)
}

func Infof(format string, args ...any) {
	logger.infof(format, args...)
}

func Warnf(format string, args ...any) {
	logger.warnf(format, args...)
}

func Errorf(format string, args ...any) {
	logger.errorf(format, args...)
}

func DPanicf(format string, args ...any) {
	logger.dPanicf(format, args...)
}

func Panicf(format string, args ...any) {
	logger.panicf(format, args...)
}

func Fatalf(format string, args ...any) {
	logger.fatalf(format, args...)
}
