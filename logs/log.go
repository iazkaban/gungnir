package logs

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := log.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05 0700",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	}

	log.SetFormatter(formatter)
}
