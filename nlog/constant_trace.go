package nlog

import (
	"net"
	"os"
)

const (
	ISTIO_TRACER_ID = "x-b3-traceid"
	ISTIO_SPAN_ID = "x-b3-spanid"
	ISTIO_PARENT_SPAN_ID = "x-b3-parentspanid"
	ISTIO_SAMPLED = "x-b3-sampled"
	ISTIO_FLAGS = "x-b3-flags"
	ISTIO_REQUEST_ID = "x-request-id"
	ISTIO_SPAN_CONTEXT = "x-ot-span-context"
)

var LOCAL_IP  = ""

func DefaultTraceId() string {

	hostname,err  :=  os.Hostname()

	if err != nil {
		hostname = "NA"
	}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "T_" + hostname + "_" + "255.255.255.255"
	}

	ip := getFirstIp(addrs)

	return "T_" + hostname + "_" + ip
}

func DefaultRequestId() string {

	hostname,err  :=  os.Hostname()

	if err != nil {
		hostname = "NA"
	}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "R_" + hostname + "_" + "255.255.255.255"
	}

	ip := getFirstIp(addrs)

	return "R_" + hostname + "_" + ip
}

func getFirstIp(addrs []net.Addr) string {
	ip := LOCAL_IP

	if ip != "" {
		return ip
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				LOCAL_IP = ip
				break
			}

		}
	}

	return ip
}
