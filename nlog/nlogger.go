package nlog

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"github.com/bingzhilanmo-bowen/mylog-4go/flog"
	"github.com/bingzhilanmo-bowen/mylog-4go/logrus"
	"time"
)

var (
	PATH_SEP = os.PathSeparator
	Counter  *CounterContainer
	Gauges   *GaugeContainer
)


type NLogger struct {
	Log            *logrus.Logger
	BisLog         *logrus.Logger
	MetricLog      *logrus.Logger
	AuditLog       *logrus.Logger
	KratosLog      *logrus.Logger
	Hooks          map[HookeType]NlogHook
	OpenPrometheus bool
}

type NLoggerConfig struct {
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

func (c *NLoggerConfig) GetLogPath() string {

	path := "/var/logs/nlog/"

	if c.LogPath != "" {
		path = c.LogPath
		i := len(path) - 1
		if !os.IsPathSeparator(path[i]) {
			path = path + string(PATH_SEP)
		}
	}

	return path
}

func (c *NLoggerConfig) GetLogFileMaxSize() int {

	if c.LogFileMaxSize > 0 {
		return c.LogFileMaxSize
	}
	return 200
}

func (c *NLoggerConfig) GetLogFileMaxAge() int {

	if c.LogFileMaxAge > 0 {
		return c.LogFileMaxAge
	}
	return 7
}

func (c *NLoggerConfig) GetOpenPrometheus() bool {
	return c.OpenPrometheus
}



func NewDefault(logPath string)  (nlg *NLogger, closeFunc func() ) {

	config := &NLoggerConfig{
		LogPath: logPath,
	}

	return NewNLogger(config)
}

func NewDefault1(logPath string) (nlg *NLogger, logr *logrus.Logger, closeFunc func()) {

	config := &NLoggerConfig{
		LogPath: logPath,
	}

	nlogger, cf := NewNLogger(config)

	return nlogger, nlogger.Log, cf
}

func NewNLogger(config *NLoggerConfig) (nlg *NLogger, closeFunc func() ) {

	//custom && system  log config
	log := logrus.New()

	flogs := &flog.FileLogger{
		Filename:  config.GetLogPath() + "root.log",
		MaxSize:   config.GetLogFileMaxSize(),
		MaxAge:    config.GetLogFileMaxAge(), //days
		LocalTime: true,
	}

	var stdOut io.Writer
    if config.RootLogOutputModel == 0  {
    	//Output to the STDOUT and File at the same time
		log.SetOutput(os.Stdout)
	}else if config.RootLogOutputModel == 1 {
		log.SetOutput(flogs)
	} else if config.RootLogOutputModel == 2 {
		stdOut = io.MultiWriter(os.Stdout, flogs)
		log.SetOutput(stdOut)
	}



	switch config.LogLevel {
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetReportCaller(config.ReportCaller)

	if config.UseThreadLocal {
		InitCache()
		log.SetFormatter(&NJSONFormatter{})
	} else {
		log.SetFormatter(BaseFormatter)
	}



	// business log config

	blog := logrus.New()
	bflog := &flog.FileLogger{
		Filename:  config.GetLogPath() + "business.log",
		MaxSize:   config.GetLogFileMaxSize(),
		MaxAge:    config.GetLogFileMaxAge(), //days
		LocalTime: true,
	}
	blog.SetLevel(logrus.InfoLevel)
	blog.SetFormatter(&NFormatter{})
	blog.SetOutput(bflog)

	// business log config

	mlog := logrus.New()
	mflog := &flog.FileLogger{
		Filename:  config.GetLogPath() + "metric.log",
		MaxSize:   config.GetLogFileMaxSize(),
		MaxAge:    config.GetLogFileMaxAge(), //days
		LocalTime: true,
	}
	mlog.SetLevel(logrus.InfoLevel)
	mlog.SetFormatter(&NFormatter{})
	mlog.SetOutput(mflog)

	if config.OpenPrometheus {
		Counter = InitCounterContainer()
		Gauges = InitGaugeContainer()
	}

	// business log config

	alog := logrus.New()
	aflog := &flog.FileLogger{
		Filename:  config.GetLogPath() + "audit.log",
		MaxSize:   config.GetLogFileMaxSize(),
		MaxAge:    config.GetLogFileMaxAge(), //days
		LocalTime: true,
	}
	alog.SetLevel(logrus.InfoLevel)
	alog.SetFormatter(&NFormatter{})
	alog.SetOutput(aflog)

	newton_log := &NLogger{
		Log:            log,
		BisLog:         blog,
		MetricLog:      mlog,
		AuditLog:       alog,
		OpenPrometheus: config.OpenPrometheus,
	}

	return newton_log, func() {
		flogs.Close()
		bflog.Close()
		mflog.Close()
		aflog.Close()
	}
}

func (n *NLogger) AddNlogHook(t HookeType,h NlogHook)  {
	if n.Hooks == nil {
		n.Hooks = make(map[HookeType]NlogHook, 3)
	}
	n.Hooks[t] = h
}

func (n *NLogger) Debugf(format string, args ...interface{}) {
	n.Log.Debugf(format, args...)
}
func (n *NLogger) Infof(format string, args ...interface{}) {
	n.Log.Infof(format, args...)
}
func (n *NLogger) Printf(format string, args ...interface{}) {
	n.Log.Printf(format, args...)
}
func (n *NLogger) Warnf(format string, args ...interface{}) {
	n.Log.Warnf(format, args...)
}
func (n *NLogger) Warningf(format string, args ...interface{}) {
	n.Log.Warningf(format, args...)
}
func (n *NLogger) Errorf(format string, args ...interface{}) {
	n.Log.Errorf(format, args...)
}
func (n *NLogger) Fatalf(format string, args ...interface{}) {
	n.Log.Fatalf(format, args...)
}
func (n *NLogger) Panicf(format string, args ...interface{}) {
	n.Log.Panicf(format, args...)
}

func (n *NLogger) Debug(args ...interface{}) {
	n.Log.Debug(args...)
}
func (n *NLogger) Info(args ...interface{}) {
	n.Log.Info(args...)
}
func (n *NLogger) Print(args ...interface{}) {
	n.Log.Print(args...)
}
func (n *NLogger) Warn(args ...interface{}) {
	n.Log.Warn(args...)
}
func (n *NLogger) Warning(args ...interface{}) {
	n.Log.Warning(args...)
}
func (n *NLogger) Error(args ...interface{}) {
	n.Log.Error(args...)
}
func (n *NLogger) Fatal(args ...interface{}) {
	n.Log.Fatal(args...)
}
func (n *NLogger) Panic(args ...interface{}) {
	n.Log.Panic(args...)
}

func (n *NLogger) DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.DebugLevel) {
		n.nlogFields(ctx).Debugf(format, args...)
	}
}
func (n *NLogger) InfofCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.InfoLevel) {
		n.nlogFields(ctx).Infof(format, args...)
	}
}
func (n *NLogger) PrintfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.InfoLevel) {
		n.nlogFields(ctx).Printf(format, args...)
	}
}
func (n *NLogger) WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.WarnLevel) {
		n.nlogFields(ctx).Warnf(format, args...)
	}
}
func (n *NLogger) WarningfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.WarnLevel) {
		n.nlogFields(ctx).Warningf(format, args...)
	}
}
func (n *NLogger) ErrorfCtx(ctx context.Context, format string, args ...interface{}) {

	if n.Log.IsLevelEnabled(logrus.ErrorLevel) {
		n.nlogFields(ctx).Errorf(format, args...)
	}
}
func (n *NLogger) FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.FatalLevel) {
		n.nlogFields(ctx).Fatalf(format, args...)
	}
}
func (n *NLogger) PanicfCtx(ctx context.Context, format string, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.PanicLevel) {
		n.nlogFields(ctx).Panicf(format, args...)
	}
}

func (n *NLogger) DebugCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.DebugLevel) {
		n.nlogFields(ctx).Debug(args...)
	}
}

func (n *NLogger) InfoCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.InfoLevel) {
		n.nlogFields(ctx).Info(args...)
	}
}
func (n *NLogger) PrintCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.InfoLevel) {
		n.nlogFields(ctx).Print(args...)
	}
}

func (n *NLogger) WarnCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.WarnLevel) {
		n.nlogFields(ctx).Warn(args...)
	}

}

func (n *NLogger) WarningCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.WarnLevel) {
		n.nlogFields(ctx).Warning(args...)
	}

}
func (n *NLogger) ErrorCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.ErrorLevel) {
		n.nlogFields(ctx).Error(args...)
	}

}
func (n *NLogger) FatalCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.FatalLevel) {
		n.nlogFields(ctx).Fatal(args...)
	}

}
func (n *NLogger) PanicCtx(ctx context.Context, args ...interface{}) {
	if n.Log.IsLevelEnabled(logrus.PanicLevel) {
		n.nlogFields(ctx).Panic(args...)
	}
}

func (n *NLogger) nlogFields(ctx context.Context) *logrus.Entry {
	traceId := ctx.Value(ISTIO_TRACER_ID)
	if traceId == nil {
		traceId = DefaultTraceId()
	}
	requestId := ctx.Value(ISTIO_REQUEST_ID)
	if requestId == nil {
		requestId = DefaultRequestId()
	}

	return n.Log.WithFields(logrus.Fields{REQUEST_ID: requestId, TRACE_ID: traceId})
}

func (n *NLogger) Count(metricName *MetricName) {
	n.CountValueT(metricName, 1, time.Now())
}
func (n *NLogger) CountT(metricName *MetricName, times time.Time) {
	n.CountValueT(metricName, 1, times)
}
func (n *NLogger) CountValue(metricName *MetricName, value float64) {
	n.CountValueT(metricName, value, time.Now())
}
func (n *NLogger) CountValueT(metricName *MetricName, value float64, times time.Time) {
	if checkMetricName(metricName.Metric) {
		metricLog := Metric2Log(metricName, METRIC_TYPE_COUNTER, value, times.Format(time.RFC3339))
		n.MetricLog.Infoln(metricLog)
		if n.OpenPrometheus {
			Counter.CounterValue(metricName, value)
		}
	} else {
		n.Log.Warningf("metric can not be nil, and must Match [a-zA-Z0-9-./_]{1,128}")
	}
}

func (n *NLogger) Gauge(metricName *MetricName, value float64) {
	n.GaugeT(metricName, value, time.Now())
}

func (n *NLogger) GaugeT(metricName *MetricName, value float64, times time.Time) {
	if checkMetricName(metricName.Metric) {
		metricLog := Metric2Log(metricName, METRIC_TYPE_GAUGE, value, times.Format(time.RFC3339))
		n.MetricLog.Infoln(metricLog)
		if n.OpenPrometheus {
			Gauges.GaugeValue(metricName, value)
		}
	} else {
		n.Log.Warningf("metric can not be nil, and must Match ^[a-zA-Z0-9_.]{1,128}$")
	}
}

func (n *NLogger) BusinessLog(bis interface{}) {
	json_bytes, _ := json.Marshal(bis)
	n.BisLog.Infoln(string(json_bytes))
}

func (n *NLogger) Audit(auditLog *AuditLog) {
	n.AuditT(auditLog, time.Now())
}

func (n *NLogger) AuditT(auditLog *AuditLog, times time.Time) {
	auditLog.DateTime = times.Format(time.RFC3339)
	alog := Audit2Log(auditLog)
	n.AuditLog.Infoln(alog)

	if h := n.getHooke(AUDIT);  h != nil {
       		ha := h.(AuditHook)
       		ha.StorageAudit(auditLog)
	}

}

func (n *NLogger) getHooke(t HookeType) NlogHook {
	if n != nil && n.Hooks != nil {
		return n.Hooks[t]
	}
	return nil
}

func checkMetricName(name string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9_.]{1,128}$", name)
	return ok
}
