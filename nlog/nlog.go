package nlog

import (
	"context"
	"time"
)

type HookeType string

const (
	AUDIT HookeType = "AUDIT"
	METRIC HookeType = "METRIC"
	BIS HookeType = "BUSINESS"
)

type NlogHook interface {
	HookType() HookeType
	Close()
}


type Nlogger interface {
	// Use Thread Local or Default
	AddNlogHook(t HookeType,h *NlogHook)
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type MetricsLogger interface {

	Count(metricName *MetricName)
	CountT(metricName *MetricName, times time.Time)
	CountValue(metricName *MetricName, value float64)
	CountValueT(metricName *MetricName, value float64, times time.Time)

	Gauge(metricName *MetricName, value float64)
	GaugeT(metricName *MetricName, value float64, times time.Time)

}

type BusinessLogger interface {
	BusinessLog(bis interface{})
}

type AuditLogger interface {
	Audit(auditLog *AuditLog)
	AuditT(auditLog *AuditLog, times time.Time)
}


type NloggerExt interface {
	Nlogger
	MetricsLogger
	BusinessLogger
	AuditLogger

	// with context
	DebugfCtx(ctx context.Context, format string, args ...interface{})
	InfofCtx(ctx context.Context, format string, args ...interface{})
	PrintfCtx(ctx context.Context, format string, args ...interface{})
	WarnfCtx(ctx context.Context, format string, args ...interface{})
	WarningfCtx(ctx context.Context, format string, args ...interface{})
	ErrorfCtx(ctx context.Context, format string, args ...interface{})
	FatalfCtx(ctx context.Context, format string, args ...interface{})
	PanicfCtx(ctx context.Context, format string, args ...interface{})

	DebugCtx(ctx context.Context, args ...interface{})
	InfoCtx(ctx context.Context, args ...interface{})
	PrintCtx(ctx context.Context, args ...interface{})
	WarnCtx(ctx context.Context, args ...interface{})
	WarningCtx(ctx context.Context, args ...interface{})
	ErrorCtx(ctx context.Context, args ...interface{})
	FatalCtx(ctx context.Context, args ...interface{})
	PanicCtx(ctx context.Context, args ...interface{})
}
