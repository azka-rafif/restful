package log

import log "github.com/sirupsen/logrus"

func GetLogger() *log.Entry {
	apiLogger := log.WithFields(log.Fields{"Server": "Info"})
	return apiLogger
}
