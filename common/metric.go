package common

type Metric struct {
	Bucket   string
	Value    float64
	Modifier string
	Sampling float32
	//optional, only for "archive" metrics
	Time uint32
	// If the gauge is delta.
	IsDelta bool
}

// MetricAmount tracks amounts of metrics.
// implicitly received is 1 per instance, but amount of sent
// depends on the sampling rate value
type MetricAmount struct {
	Bucket   string
	Sampling float32
}
