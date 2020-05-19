package test

import (
	log "github.com/bingzhilanmo-bowen/mylog-4go"
	"github.com/bingzhilanmo-bowen/mylog-4go/flog"
	"github.com/bingzhilanmo-bowen/mylog-4go/logrus"
	"github.com/bingzhilanmo-bowen/mylog-4go/nlog"
	"testing"
)

func BenchmarkNlog(b *testing.B) {

	log.ConfigLog(&log.LoggerConfig{LogPath: "log"})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log.Infof("hello bench mark %s", "bowen")
		log.Debugf("hello bench mark %s", "bowen")
		log.Warningf("hello bench mark %s", "bowen")
		log.Warnf("hello bench mark %s", "bowen")
		log.Errorf("hello bench mark %s", "bowen")
		log.Printf("hello bench mark %s", "bowen")
	}

}

func BenchmarkLogrus(b *testing.B) {

	logrus.SetLevel(logrus.InfoLevel)

	flogs := &flog.FileLogger{
		Filename: "log/root2.log",
	}

	logrus.SetOutput(flogs)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(nlog.BaseFormatter)

	b.ResetTimer()


	for i := 0; i < b.N; i++ {
		logrus.Infof("hello bench mark %s", "bowen")
		logrus.Debugf("hello bench mark %s", "bowen")
		logrus.Warningf("hello bench mark %s", "bowen")
		logrus.Warnf("hello bench mark %s", "bowen")
		logrus.Errorf("hello bench mark %s", "bowen")
		logrus.Printf("hello bench mark %s", "bowen")
	}
}
