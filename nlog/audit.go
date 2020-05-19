package nlog

import "encoding/json"

type AuditHook interface {
	NlogHook
	StorageAudit(audit *AuditLog) error
}

type AuditLog struct {
	Operator   *OperatorInfo     `json:"operator"`
	Action     string            `json:"action"`
	OptObject  string            `json:"optObject"`
	Tags       map[string]string `json:"tags"`
	RawData    interface{}       `json:"rawData"`
	UpdateData interface{}       `json:"updateData"`
	DateTime   string            `json:"dateTime"`
}

type OperatorInfo struct {
	Name   string
	Email  string
	Mobile string
}

func Audit2Log(audit *AuditLog) string {
	if audit == nil {
		return ""
	}

	json_bytes, _ := json.Marshal(audit)
	return string(json_bytes)
}