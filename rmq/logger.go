package rmq

import (
	"base-frame/log"

	"github.com/gowins/dionysus/rmq"
)

type rlogger struct {
}

func (r rlogger) Debug(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Debug(msg)
}

func (r rlogger) Info(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Info(msg)
}

func (r rlogger) Warning(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Warn(msg)
}

func (r rlogger) Error(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Error(msg)
}

func (r rlogger) Fatal(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Fatal(msg)
}

func (r rlogger) Level(level string) {
}

func (r rlogger) OutputPath(path string) (err error) {
	return nil
}

func SetLogger() {
	rmq.SetLogger(&rlogger{})
}
