package nlog_4go

import (
	"context"
	"testing"
)


type BisTest struct {
	Name string
	Age int8
	Six int
}

func TestLogger(t *testing.T) {

	ConfigLog(&LoggerConfig{
		LogPath: "log",
		LogLevel: "DEBUG",
		RootLogOutputModel: 2,
	})

	Infof("test info %s, %d", "hydxin", 55667788)
	Debugf("test debug %s, %d", "hydxin", 55667788)
	InfofCtx(context.Background(), "test info %s, %d", "hydxin", 55667788)
	Count(NewMetricBuilder().Metric("test").Tags("k1", "v1").Build())
	Gauge(NewMetricBuilder().Metric("testGauge").Tags("k1", "v1").Build(), 100)
	BusinessLog(&BisTest{Name: "Hydxin", Age: 29, Six: 1})

	RAW := make(map[string]interface{})
	RAW["ID"] = 1
	RAW["NAME"] = "BOWEN"

	UP := make(map[string]interface{})
	UP["ID"] = 1
	UP["NAME"] = "HYDXIN"

	auditLog := NewBuilderAudit().Operator("bowen", "email", "xxxx").
		Action("INSERT").OptObject("ORG INFO").Tags("k", "v").RawData(RAW).UpdateData(UP).Build()

	Audit(auditLog)

}
