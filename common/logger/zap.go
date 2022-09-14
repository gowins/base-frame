package logger

import (
	"os"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultTimeKey       = "time"
	DefaultLevelKey      = "level"
	DefaultNameKey       = "logger"
	DefaultCallerKey     = "caller"
	DefaultMessageKey    = "msg"
	DefaultStacktraceKey = "stacktrace"
)

var (
	once               sync.Once
	coustomTimeEncoder = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
)

// 初始化设置
func Setup(opts ...zap.Option) {
	if len(opts) == 0 {
		opts = defaultOpts()
	}
	once.Do(func() {
		writer, err := zapx.NewZapWriter(opts...)
		logx.Must(err)
		logx.SetWriter(writer)
		logx.Info("log setup success!!!")
	})

}

// zap logger
// type Logger struct {
// 	core zapcore.Core
// 	development bool
// 	addCaller   bool
// 	onFatal     zapcore.CheckWriteAction // default is WriteThenFatal
// 	name        string
// 	errorOutput zapcore.WriteSyncer
// 	addStack zapcore.LevelEnabler
// 	callerSkip int
// 	clock zapcore.Clock
// }

func defaultOpts() (opts []zap.Option) {
	devOpt := zap.Development()
	callerOpt := zap.AddCaller()
	fatalOpt := zap.OnFatal(zapcore.WriteThenNoop)
	levelOpt := zap.IncreaseLevel(zap.InfoLevel)
	initialFields := zap.Fields(zap.String("serviceName", "base-frame"))
	opts = append(opts, devOpt, callerOpt, fatalOpt, levelOpt, initialFields)
	// warpCore
	cfg := defaultConfig()
	cfg.EncoderConfig.EncodeTime = coustomTimeEncoder
	wrapConsoleCore := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		// return newConsoleCore(*cfg, []zapcore.WriteSyncer{})
		return newJsonCore(*cfg, []zapcore.WriteSyncer{})
	})
	opts = append(opts, wrapConsoleCore)
	return opts
}

func newConsoleCore(cfg zap.Config, opts []zapcore.WriteSyncer) zapcore.Core {
	opts = append(opts, zapcore.AddSync(os.Stdout))
	syncWriter := zapcore.NewMultiWriteSyncer(opts...)
	return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), syncWriter, cfg.Level)
}

func newJsonCore(cfg zap.Config, opts []zapcore.WriteSyncer) zapcore.Core {
	opts = append(opts, zapcore.AddSync(os.Stdout))
	syncWriter := zapcore.NewMultiWriteSyncer(opts...)
	return zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), syncWriter, cfg.Level)
}

func defaultEncoderConfig() (enc zapcore.EncoderConfig) {
	return zapcore.EncoderConfig{
		TimeKey:        DefaultTimeKey,
		LevelKey:       DefaultLevelKey,
		NameKey:        DefaultNameKey,
		CallerKey:      DefaultCallerKey,
		MessageKey:     DefaultMessageKey,
		StacktraceKey:  DefaultStacktraceKey,
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func defaultConfig() (cfg *zap.Config) {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "json",                                              // 输出格式 console 或 json
		EncoderConfig:    defaultEncoderConfig(),                              // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "base-frame"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"},                                  // 输出到指定文件 stdout（标准输出，正常颜色）[]string{"stdout", "./logs/base.log"}
		ErrorOutputPaths: []string{"stderr"},                                  // 	stderr（错误输出，红色）
	}
}

func warpZapCore(cfg *zap.Config) []zap.Option {
	if cfg == nil {
		cfg = defaultConfig()
	}
	sink, closeOut, _ := zap.Open(cfg.OutputPaths...)
	errSink, _, err := zap.Open(cfg.ErrorOutputPaths...)
	if err != nil {
		closeOut()
	}
	newcore := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), sink, cfg.Level)
	opts := []zap.Option{zap.ErrorOutput(errSink)}
	opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return newcore
	}))
	return opts
}
