package test

import (
	"context"
	log "github.com/bingzhilanmo-bowen/mylog-4go"
	"testing"
)

func TestDefaultLog(t *testing.T) {

	log.Infof("default log : %s", "data default log")
	log.Info("test infoinfo")

	log.InfofCtx(context.Background(),"default ctx log : %s", "data ctx default log")

	//BusinessLog(&nlog.BisTest{nlog.Name: "HydxinDefault", Age: 39, Six: 1})

	log.Count(log.NewMetricBuilder().Metric("b@owenD").Tags("t1","v1").TagsOnFly("tf1", "vfr1").Build())
	log.Gauge(log.NewMetricBuilder().Metric("hydxinD").Tags("t1","v1").TagsOnFly("tf1", "vfr1").Build(), 700)

}


