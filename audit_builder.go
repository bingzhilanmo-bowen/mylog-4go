package nlog_4go

import "github.com/bingzhilanmo-bowen/mylog-4go/nlog"

type AuditBuilder interface {
	OperatorName(name string) AuditBuilder
	OperatorEmail(email string) AuditBuilder
	Operator(name, email, mobile string) AuditBuilder
	Action(action string) AuditBuilder
	OptObject(optObject string) AuditBuilder
	Tags(key, values string) AuditBuilder
	MarkerError(err error) AuditBuilder
	RawData(rawData interface{}) AuditBuilder
	UpdateData(updateData interface{}) AuditBuilder
	Build() *nlog.AuditLog
}

type BuilderAudit struct {
	audit *nlog.AuditLog
}

func NewBuilderAudit() *BuilderAudit {
	return &BuilderAudit{audit:&nlog.AuditLog{}}
}

func (b *BuilderAudit) OperatorName(name string) AuditBuilder {
	if b.audit.Operator == nil {
		opt := &nlog.OperatorInfo{Name: name}
		b.audit.Operator = opt
		return b
	}else {
		b.audit.Operator.Name = name
		return b
	}
}

func (b *BuilderAudit) OperatorEmail(email string) AuditBuilder {
	if b.audit.Operator == nil {
		opt := &nlog.OperatorInfo{Email: email}
		b.audit.Operator = opt
		return b
	}else {
		b.audit.Operator.Email = email
		return b
	}
}

func (b *BuilderAudit) Operator(name, email, mobile string) AuditBuilder {
	opt := &nlog.OperatorInfo{Name: name, Email: email, Mobile: mobile}
	b.audit.Operator = opt
	return b
}

func (b *BuilderAudit) Action(action string) AuditBuilder {
	b.audit.Action = action
	return b
}

func (b *BuilderAudit) OptObject(optObject string) AuditBuilder {
	b.audit.OptObject = optObject
	return b
}

func (b *BuilderAudit) Tags(key, values string) AuditBuilder {
	if b.audit.Tags == nil {
		b.audit.Tags = make(map[string]string, 8)
	}
	b.audit.Tags[key] = values
	return b
}

func (b *BuilderAudit) MarkerError(err error) AuditBuilder {
	if b.audit.Tags == nil {
		b.audit.Tags = make(map[string]string, 8)
	}
	if err != nil {
		b.audit.Tags["error"] = "true"
		b.audit.Tags["err_msg"] = err.Error()
	}
	return b
}

func (b *BuilderAudit) RawData(rawData interface{}) AuditBuilder {
	b.audit.RawData = rawData
	return b
}

func (b *BuilderAudit) UpdateData(updateData interface{}) AuditBuilder {
	b.audit.UpdateData = updateData
	return b
}

func (b *BuilderAudit) Build() *nlog.AuditLog {
	return b.audit
}
