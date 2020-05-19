package nlog_4go

import "github.com/bingzhilanmo-bowen/mylog-4go/nlog"

type Builder interface {
	Metric(metricName string) Builder
	Tags(key, value string) Builder
	TagsOnFly(key, value string) Builder
	Build() *nlog.MetricName
}

type MetricNameBuilder struct {
	metric *nlog.MetricName
}

func NewMetricBuilder() *MetricNameBuilder  {
	return &MetricNameBuilder{metric:&nlog.MetricName{}}
}

func (m *MetricNameBuilder) Metric(metricName string) Builder {

	if m.metric == nil {
		m.metric = &nlog.MetricName{}
	}
	m.metric.Metric = metricName

	return m
}

func (m *MetricNameBuilder) Tags(key, value string) Builder {

	if m.metric == nil {
		m.metric = &nlog.MetricName{}
	}

	if m.metric.Tags == nil {
		m.metric.Tags = make(map[string]string, 8)
	}

	m.metric.Tags[key] = value
	return m
}

func (m *MetricNameBuilder) TagsOnFly(key, value string) Builder {

	if m.metric == nil {
		m.metric = &nlog.MetricName{}
	}

	if m.metric.TagsOnFly == nil {
		m.metric.TagsOnFly = make(map[string]string, 8)
	}

	m.metric.TagsOnFly[key] = value

	return m
}

func (m *MetricNameBuilder) Build() *nlog.MetricName {
	return m.metric
}
