package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	//设置日志级别
	level := zapcore.DebugLevel

	//编码器设置：决定日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,//彩色输出
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	//文件写入
	fileWriter :=zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/video-platform.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     30,   // days
		Compress:   true, // disabled by default
	})

	//控制台输出
	consoleEncoder :=zapcore.NewConsoleEncoder(encoderConfig)
	consoleWriter :=zapcore.AddSync(os.Stdout)

	//组合core
	core := zapcore.NewCore(consoleEncoder,consoleWriter,level)
	core = zapcore.NewTee(core,zapcore.NewCore(consoleEncoder,fileWriter,level))
	
	//构建日志
	logger :=zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1))
	Logger =logger.Sugar()
}