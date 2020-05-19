package nlog

import "github.com/bingzhilanmo-bowen/mylog-4go/logrus"

type NFormatter struct {

}


// Format building log message.
func (f *NFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(entry.Message ), '\n'), nil
}
