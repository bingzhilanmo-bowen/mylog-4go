package nlog

import (
	"encoding/json"
	"sort"
)

var (
	METRIC_TYPE_COUNTER = "Counter"
	METRIC_TYPE_GAUGE = "Gauge"
)

type MetricName struct {
	Metric    string
	Tags      map[string]string
	TagsOnFly map[string]string
}

func (metricName *MetricName) GetTagsKey() []string  {
	i := 0
	keys := make([]string, len(metricName.Tags))
	for k := range metricName.Tags {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

func (metricName *MetricName) GetTagsValues() []string  {
	i := 0
	values := make([]string, len(metricName.Tags))
	for _,k := range metricName.GetTagsKey() {
		values[i] = metricName.Tags[k]
		i++
	}

	return values
}


type MetricLog struct {
	Metric    string `json:"metric"`
	Tags      map[string]string `json:"tags"`
	TagsOnFly map[string]string `json:"tagsOnFly"`
	MetricType string `json:"metricType"`
	Value interface{} `json:"value"`
	DateTime interface{} `json:"dateTime"`
}

func Metric2Log(metricName *MetricName, metricType  string, value, dateTime interface{}) string{

	newLog := &MetricLog{
		Metric: metricName.Metric,
		Tags: metricName.Tags,
		TagsOnFly: metricName.TagsOnFly,
		MetricType: metricType,
		Value: value,
		DateTime: dateTime,
	}
	json_bytes, _ := json.Marshal(newLog)

	return string(json_bytes)
}




