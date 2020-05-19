package nlog

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type CounterContainer struct {
	counter sync.Map
    lock sync.Mutex
 }


func InitCounterContainer() *CounterContainer{
	return &CounterContainer{}
}

func (c *CounterContainer) getCounter(metricName string)  *prometheus.CounterVec {
	value,ok := c.counter.Load(metricName)

	if !ok {
		return nil
	}

	return value.(*prometheus.CounterVec)
}

func (c *CounterContainer) addNew(metric *MetricName) *prometheus.CounterVec {
	c.lock.Lock()

	defer c.lock.Unlock()

	vec := c.getCounter(metric.Metric)

	if vec != nil {
		return vec
	}

	newCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metric.Metric,
			Help: "Counter Add By Newton",
		}, metric.GetTagsKey(),
	)
	prometheus.MustRegister(newCounter)
	c.counter.Store(metric.Metric, newCounter)
	return newCounter

}

func (c *CounterContainer) CounterValue(metric *MetricName, value float64)  {
	vec := c.getCounter(metric.Metric)
	if vec == nil {
		vec = c.addNew(metric)
	}
    vec.WithLabelValues(metric.GetTagsValues()...).Add(value)
}


type GaugeContainer struct {
	gauge sync.Map
	lock sync.Mutex
}

func InitGaugeContainer() *GaugeContainer{
	return &GaugeContainer{}
}

func (g *GaugeContainer) getGauge(metricName string)  *prometheus.GaugeVec {
	value,ok := g.gauge.Load(metricName)

	if !ok {
		return nil
	}

	return value.(*prometheus.GaugeVec)
}

func (g *GaugeContainer) addNew(metric *MetricName) *prometheus.GaugeVec {
	g.lock.Lock()

	defer g.lock.Unlock()

	vec := g.getGauge(metric.Metric)

	if vec != nil {
		return vec
	}

	newGauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metric.Metric,
			Help: "Gauge Add By Newton",
		},metric.GetTagsKey(),
	)
	prometheus.MustRegister(newGauge)
	g.gauge.Store(metric.Metric, newGauge)
	return newGauge

}

func (g *GaugeContainer) GaugeValue(metric *MetricName, value float64)  {
	vec := g.getGauge(metric.Metric)
	if vec == nil {
		vec = g.addNew(metric)
	}
	vec.WithLabelValues(metric.GetTagsValues()...).Add(value)
}

