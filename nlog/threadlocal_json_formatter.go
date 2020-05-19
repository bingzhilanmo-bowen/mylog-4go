package nlog

import (
	"github.com/bingzhilanmo-bowen/mylog-4go/logrus"
	"time"
)

type NJSONFormatter struct {
}

var BaseFormatter = &logrus.JSONFormatter{
	TimestampFormat: time.RFC3339,
	FieldMap: logrus.FieldMap{
		logrus.FieldKeyMsg:         "message",
		logrus.FieldKeyTime:        "dateTime",
		logrus.FieldKeyLogrusError: "errors",
	},}

// Format building log message.
func (f *NJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	 if cache, ok := GetThreadCache(); ok {
		 data := make(logrus.Fields , len(entry.Data) + 2)
		 data[REQUEST_ID] = cache.RequestId
		 data[TRACE_ID] = cache.TraceId
		 entry.Data = data
	 }

	return BaseFormatter.Format(entry)
}
