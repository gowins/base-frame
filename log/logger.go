package log

import (
	"github.com/gowins/dionysus/log"
)

var (
	logger      log.Logger
	projectName string
	Info        func(args ...interface{})
	Warn        func(args ...interface{})
	Error       func(args ...interface{})
	Fatal       func(args ...interface{})
	Debug       func(args ...interface{})
	Infof       func(format string, args ...interface{})
	Warnf       func(format string, args ...interface{})
	Errorf      func(format string, args ...interface{})
	Fatalf      func(format string, args ...interface{})
	Debugf      func(format string, args ...interface{})

	WithField  func(key string, value interface{}) log.Logger
	WithFields func(fields map[string]interface{}) log.Logger
)

type Option log.Option

func Setup(opts ...log.Option) {
	os := []log.Option{ // 根据实际需求添加option
		log.WithLevelEnabler(log.DebugLevel),
		log.WithEncoderCfg(log.NewEncoderConfig()),
		log.AddCallerSkip(1),
		log.AddCaller(),
	}

	//todo: opts隐藏与否
	logger, _ = log.New(log.ZapLogger, append(os, opts...)...)
	if projectName != "" {
		logger = logger.WithField("app", projectName)
	}
	Debug = logger.Debug
	Debugf = logger.Debugf
	Info = logger.Info
	Infof = logger.Infof
	Warn = logger.Warn
	Warnf = logger.Warnf
	Error = logger.Error
	Errorf = logger.Errorf
	Fatal = logger.Fatal
	Fatalf = logger.Fatalf
	WithField = logger.WithField
	WithFields = logger.WithFields

}

func SetProjectName(pname string) Option {
	return log.Fields(map[string]interface{}{
		"app": pname,
	})
}
