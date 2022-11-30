package metrics

// Counter is metrics counter.
// Counter 只增不减的计数器， 代表指标累计值
type Counter interface {
	With(lvs ...string) Counter
	Inc()
	Add(delta float64)
}

// Gauge is metrics gauge.
// Gauge 可增可减的仪表盘， 反应系统的当前状态，
type Gauge interface {
	With(lvs ...string) Gauge
	Set(value float64)
	Add(delta float64)
	Sub(delta float64)
}

// Observer is metrics observer.
// Observer 指标统计数据
type Observer interface {
	With(lvs ...string) Observer
	Observe(float64)
}
