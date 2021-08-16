/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/11/6 16:14
*/

package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	zcommon "rosefintech-rosl2/common"
)

var Logger *zap.Logger

func LogInit(config *zcommon.Config) {
	hook := lumberjack.Logger{
		Filename:   "./logs/tesrasch.zlog",
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// set log level
	atomicLevel := zap.NewAtomicLevel()
	if config.Debug == true {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.ErrorLevel)
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)


	caller := zap.AddCaller()

	development := zap.Development()

	Logger = zap.New(core, caller, development)

	Logger.Info("logger init is succeed")
}
