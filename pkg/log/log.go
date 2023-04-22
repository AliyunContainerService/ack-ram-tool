package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"
)

var logLevelEncoders = map[string]zapcore.LevelEncoder{
	"lower":        zapcore.LowercaseLevelEncoder,
	"capital":      zapcore.CapitalLevelEncoder,
	"color":        zapcore.LowercaseColorLevelEncoder,
	"capitalcolor": zapcore.CapitalColorLevelEncoder,
}

var Logger *zap.SugaredLogger
var (
	DefaultLogLevel        = "INFO"
	DefaultLogLevelKey     = "level"
	DefaultLogLevelEncoder = "capital"
)

const (
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelDebug = "debug"
)

func init() {
	Logger, _ = newLogger(DefaultLogLevel, DefaultLogLevelKey, DefaultLogLevelEncoder)
}

func SetupLogger(logLevel string, logLevelKey string, logLevelEncoder string) error {
	logger, err := newLogger(logLevel, logLevel, logLevelEncoder)
	if err != nil {
		return err
	}
	Logger = logger
	return nil
}

func newLogger(logLevel string, logLevelKey string, logLevelEncoder string) (*zap.SugaredLogger, error) {
	encoder, ok := logLevelEncoders[logLevelEncoder]
	if !ok {
		return nil, fmt.Errorf("invalid log level encoder: %v", logLevelEncoder)
	}
	var zlog *zap.Logger

	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		sink := zapcore.AddSync(os.Stderr)
		lvl := zap.NewAtomicLevelAt(zap.DebugLevel)

		eCfg := zap.NewDevelopmentEncoderConfig()
		eCfg.LevelKey = logLevelKey
		eCfg.EncodeLevel = encoder
		eCfg.EncodeTime = zapcore.RFC3339TimeEncoder
		eCfg.ConsoleSeparator = " "
		enc := zapcore.NewConsoleEncoder(eCfg)

		zlog = zap.New(zapcore.NewCore(enc, sink, lvl))
		zlog = zlog.WithOptions(zap.AddCaller())
		logger := zapr.NewLogger(zlog)
		klog.SetLogger(logger)
	case "WARN", "WARNING", "ERROR":
		zlog = setLoggerForProduction(logLevelKey, encoder)
	case "INFO":
		fallthrough
	default:
		sink := zapcore.AddSync(os.Stderr)
		lvl := zap.NewAtomicLevelAt(zap.InfoLevel)

		eCfg := zap.NewProductionEncoderConfig()
		eCfg.LevelKey = logLevelKey
		eCfg.EncodeLevel = encoder
		eCfg.EncodeTime = zapcore.RFC3339TimeEncoder
		eCfg.ConsoleSeparator = " "
		enc := zapcore.NewConsoleEncoder(eCfg)

		zlog = zap.New(zapcore.NewCore(enc, sink, lvl))
		zlog = zlog.WithOptions(zap.AddCaller())
		logger := zapr.NewLogger(zlog)
		klog.SetLogger(logger)
	}

	return zlog.Sugar(), nil
}

func setLoggerForProduction(logLevelKey string, encoder zapcore.LevelEncoder) *zap.Logger {
	sink := zapcore.AddSync(os.Stderr)
	var opts []zap.Option
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.LevelKey = logLevelKey
	encCfg.EncodeLevel = encoder
	encCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encCfg.ConsoleSeparator = " "
	enc := zapcore.NewConsoleEncoder(encCfg)
	lvl := zap.NewAtomicLevelAt(zap.WarnLevel)

	opts = append(opts, zap.AddStacktrace(zap.ErrorLevel),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSamplerWithOptions(core, time.Second, 100, 100)
		}),
		zap.AddCallerSkip(1), zap.ErrorOutput(sink),
		zap.AddCaller(),
	)

	zlog := zap.New(zapcore.NewCore(enc, sink, lvl))
	zlog = zlog.WithOptions(opts...)
	newlogger := zapr.NewLogger(zlog)
	klog.SetLogger(newlogger)

	return zlog
}
