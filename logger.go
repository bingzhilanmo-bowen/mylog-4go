package nlog_4go

import (
	"context"
	"github.com/bingzhilanmo-bowen/mylog-4go/logrus"
	"github.com/bingzhilanmo-bowen/mylog-4go/nlog"
	"time"
)

var (
	nl, cf = nlog.NewDefault("/var/logs/nlog")
)

const (
	AUDIT = "AUDIT"
	METRIC = "METRIC"
	BIS = "BUSINESS"
)

type LoggerConfig struct {
	LogLevel           string
	OpenPrometheus     bool
	UseThreadLocal     bool
	RootLogOutputModel int
	LogPath            string
	LogFileMaxSize     int
	LogFileMaxAge      int
	LocalTime          bool
	Compress           bool
	ReportCaller       bool
}

//自定义配置Log
func ConfigLog(config *LoggerConfig) {
	conf := &nlog.NLoggerConfig{
		LogLevel:           config.LogLevel,
		OpenPrometheus:     config.OpenPrometheus,
		UseThreadLocal:     config.UseThreadLocal,
		RootLogOutputModel: config.RootLogOutputModel,
		LogPath:            config.LogPath,
		LogFileMaxSize:     config.LogFileMaxSize,
		LogFileMaxAge:      config.LogFileMaxAge,
		LocalTime:          config.LocalTime,
		Compress:           config.Compress,
		ReportCaller:       config.ReportCaller,
	}
	nl, cf = nlog.NewNLogger(conf)
}

func Debugf(format string, args ...interface{}) {
	nl.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	nl.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	nl.Printf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	nl.Warnf(format, args...)
}
func Warningf(format string, args ...interface{}) {
	nl.Warningf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	nl.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	nl.Fatalf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	nl.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	nl.Debug(args...)
}
func Info(args ...interface{}) {
	nl.Info(args...)
}
func Print(args ...interface{}) {
	nl.Print(args...)
}
func Warn(args ...interface{}) {
	nl.Warn(args...)
}
func Warning(args ...interface{}) {
	nl.Warning(args...)
}
func Error(args ...interface{}) {
	nl.Error(args...)
}
func Fatal(args ...interface{}) {
	nl.Fatal(args...)
}
func Panic(args ...interface{}) {
	nl.Panic(args...)
}

// with context
func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.DebugfCtx(ctx, format, args...)
}
func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	nl.InfofCtx(ctx, format, args...)
}
func PrintfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.PrintfCtx(ctx, format, args...)
}
func WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.WarnfCtx(ctx, format, args...)
}
func WarningfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.WarningfCtx(ctx, format, args...)
}
func ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.ErrorfCtx(ctx, format, args...)
}
func FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.FatalfCtx(ctx, format, args...)
}
func PanicfCtx(ctx context.Context, format string, args ...interface{}) {
	nl.PanicfCtx(ctx, format, args...)
}

func DebugCtx(ctx context.Context, args ...interface{}) {
	nl.DebugCtx(ctx, args...)
}
func InfoCtx(ctx context.Context, args ...interface{}) {
	nl.InfoCtx(ctx, args...)
}
func PrintCtx(ctx context.Context, args ...interface{}) {
	nl.PrintCtx(ctx, args...)
}
func WarnCtx(ctx context.Context, args ...interface{}) {
	nl.WarnCtx(ctx, args...)
}
func WarningCtx(ctx context.Context, args ...interface{}) {
	nl.WarningCtx(ctx, args...)
}
func ErrorCtx(ctx context.Context, args ...interface{}) {
	nl.ErrorCtx(ctx, args...)
}
func FatalCtx(ctx context.Context, args ...interface{}) {
	nl.FatalCtx(ctx, args...)
}
func PanicCtx(ctx context.Context, args ...interface{}) {
	nl.PanicCtx(ctx, args...)
}

func Count(metricName *nlog.MetricName) {
	nl.Count(metricName)
}
func CountT(metricName *nlog.MetricName, times time.Time) {
	nl.CountT(metricName, times)
}
func CountValue(metricName *nlog.MetricName, value float64) {
	nl.CountValue(metricName, value)
}
func CountValueT(metricName *nlog.MetricName, value float64, times time.Time) {
	nl.CountValueT(metricName, value, times)
}

func Gauge(metricName *nlog.MetricName, value float64) {
	nl.Gauge(metricName, value)
}

func GaugeT(metricName *nlog.MetricName, value float64, times time.Time) {
	nl.GaugeT(metricName, value, times)
}

func BusinessLog(bis interface{}) {
	nl.BusinessLog(bis)
}

func Audit(auditLog *nlog.AuditLog) {
	nl.Audit(auditLog)
}

func AuditT(auditLog *nlog.AuditLog, times time.Time) {
	nl.AuditT(auditLog, times)
}

func AddNlogHook(t string, h nlog.NlogHook){
	if t == AUDIT {
		nl.AddNlogHook(nlog.AUDIT, h)
	}else if t == METRIC {
		nl.AddNlogHook(nlog.METRIC, h)
	}else {
		nl.AddNlogHook(nlog.BIS, h)
	}
}

func KratosLog()  *logrus.Logger {
	return nl.KratosLog
}

func Close()  {
	cf()
}
