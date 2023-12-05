package log

import (
	"github.com/fatih/color"
	"go.uber.org/zap/zapcore"
)

var (
	_levelToColor = map[zapcore.Level]*color.Color{
		zapcore.DebugLevel:  color.New(color.FgMagenta),
		zapcore.InfoLevel:   color.New(color.FgBlue),
		zapcore.WarnLevel:   color.New(color.FgYellow),
		zapcore.ErrorLevel:  color.New(color.FgRed),
		zapcore.DPanicLevel: color.New(color.FgRed),
		zapcore.PanicLevel:  color.New(color.FgRed),
		zapcore.FatalLevel:  color.New(color.FgRed),
	}
	_unknownLevelColor = color.New(color.FgRed)

	_levelToLowercaseColorString = make(map[zapcore.Level]*color.Color, len(_levelToColor))
	_levelToCapitalColorString   = make(map[zapcore.Level]*color.Color, len(_levelToColor))
)

func init() {
	for level, color := range _levelToColor {
		_levelToLowercaseColorString[level] = color
		_levelToCapitalColorString[level] = color
	}
}

func lowercaseColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	c, ok := _levelToLowercaseColorString[l]
	if !ok {
		c = _unknownLevelColor
	}
	s := c.Sprint(l.String())
	enc.AppendString(s)
}

func capitalColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	c, ok := _levelToCapitalColorString[l]
	if !ok {
		c = _unknownLevelColor
	}
	s := c.Sprint(l.CapitalString())
	enc.AppendString(s)
}
