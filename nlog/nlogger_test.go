package nlog

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"runtime"
	"github.com/bingzhilanmo-bowen/mylog-4go/logrus"
	"testing"
)

type BisTest struct {
	Name string
	Age int8
	Six int
}

func TestLogrus(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.Infof("test %s", "3423423")
}

func TestNewNlogger(t *testing.T) {

	config := &NLoggerConfig{
		LogPath: "log",
		OpenPrometheus: true,
	}

	l,cl := NewNLogger(config)

	l.Infof("newton log info %s", "bowen")
	l.Info("just info")
	l.Debugf("hi data %d", 100)
	l.Warnf("ni hao warn %v", &BisTest{Name: "Hydxin", Age: 29, Six: 1})
	l.Errorf("error log %v", errors.New("test error log"))


	l.BusinessLog(&BisTest{Name: "Hydxin", Age: 29, Six: 1})

	ctx := context.WithValue(context.Background(), ISTIO_TRACER_ID,"newtontraceid2")

	ctx2:= context.WithValue(ctx, ISTIO_REQUEST_ID,"newtonrequestid2")


	l.InfofCtx(ctx2, " test ctx f %s", "cloudminds")
	l.InfofCtx(context.Background(), " test ctx is null %s", "cloudminds")

	l.InfofCtx(context.Background(), " test ip has get ctx is null %s", "cloudminds")

	cl()

}


func TestNewNloggerThread(t *testing.T) {

	config := &NLoggerConfig{
		LogPath: "log",
		OpenPrometheus: true,
		UseThreadLocal: true,
	}

	ctx := context.WithValue(context.Background(), ISTIO_TRACER_ID,"newtontraceid2")

	ctx2:= context.WithValue(ctx, ISTIO_REQUEST_ID,"newtonrequestid2")

	InitCache()

	SetTraceInfoFromContext(ctx2)


	l, cl := NewNLogger(config)

	l.Infof("newton log info %s", "bowen")
	l.Info("just info")
	l.Debugf("hi data %d", 100)
	l.Warnf("ni hao warn %v", &BisTest{Name: "Hydxin", Age: 29, Six: 1})
	l.Errorf("error log %v", errors.New("test error log"))


	l.BusinessLog(&BisTest{Name: "Hydxin", Age: 29, Six: 1})


    cl()

}

func TestFileLog(t *testing.T)  {

	config := &NLoggerConfig{
		LogPath: "log",
		LogFileMaxSize: 2,
		BridgeKratos:true,
	}

	l,cl := NewNLogger(config)


	fmt.Println("start add log")

	i := 1

	for i < 50000 {

		l.Infof("newton log info %s", "bowen")
		l.Info("just info")
		l.Debugf("hi data %d", 100)
		l.Warnf("ni hao warn %v", &BisTest{Name: "Hydxin", Age: 29, Six: 1})
		l.Errorf("error log %v", errors.New("test error log"))


		l.BusinessLog(&BisTest{Name: "Hydxin", Age: 29, Six: 1})
		i++
		//fmt.Println(i)
	}
	l.KratosLog.Info("test kratos log")

    cl()

}

func TestHostName(t *testing.T){
	fmt.Println(runtime.GOOS)
	name,_  :=  os.Hostname()

	fmt.Println(name)


	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}

}

func TestRegx(t *testing.T) {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9_.]{1,128}$", "5")

	fmt.Println(ok)
}

func TestRange(t *testing.T)  {

	var default_file_keys = [3]string{"root.log", "business.log", "metrics.log"}

	for _, fk := range default_file_keys {
		fmt.Println(fk)
	}

}
