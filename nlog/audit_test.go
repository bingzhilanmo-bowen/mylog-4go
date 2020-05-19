package nlog

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAudit2Log(t *testing.T) {

	RAW := make(map[string]interface{})
	RAW["ID"] = 1
	RAW["NAME"] = "BOWEN"

	auditLog := &AuditLog{}

	config := &NLoggerConfig{
		LogPath: "log",
		OpenPrometheus: true,
	}

	l, cl := NewNLogger(config)

	l.AddNlogHook(AUDIT, &TestHook{})

	l.Audit(auditLog)


	cl()
}

type TestHook struct {

}

func (th *TestHook) HookType() HookeType  {
	return AUDIT
}

func (th *TestHook) Close()  {
	fmt.Println("hook close")
}

func (th *TestHook) StorageAudit(audit *AuditLog) error {

	bts , _ := json.Marshal(audit)

	fmt.Printf("hooke fmt %s", string(bts))

	return nil
}

