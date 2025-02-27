package kafka

import "github.com/prometheus/client_golang/prometheus"

const (
	LblType               = "type"
	LblRule               = "rule"
	LblOp                 = "op"
	LblStatus             = "status"
	LblMessageTooLargeErr = "messageTooLarge"
	LblReq                = "req"
	LblRetry              = "retry"
	LblSuccess            = "ok"
	LblErr                = "err"
	LblBackOff            = "backoff"
	LblMsg                = "msg"
	LblBytes              = "bytes"
)

var (
	KafkaWriterErrCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka_client",
		Subsystem: "writer",
		Name:      "err_counter",
		Help:      "counter of Kafka Client Writer IO",
	}, []string{LblType, LblRule, LblOp})

	KafkaWriterBatchCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka_client",
		Subsystem: "writer_batch",
		Name:      "counter",
		Help:      "counter of Kafka Client Writer IO",
	}, []string{LblType, LblStatus, LblRule, LblOp})

	KafkaWriterBatchTotalBytes = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka_client",
		Subsystem: "writer",
		Name:      "total_bytes",
		Help:      "Total Bytes Kafka Client Writer IO",
	}, []string{LblType, LblRule, LblOp})

	KafkaWriterBatchDurationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "kafka_client",
		Subsystem: "writer",
		Name:      "duration_microseconds",
		Help:      "Histogram of Kafka Client Writer IO",
		Buckets:   prometheus.ExponentialBuckets(10, 2, 20), // 10us ~ 5s
	}, []string{LblType, LblRule, LblOp})

	KafkaWriterBatchGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "kafka_client",
		Subsystem: "writer_batch",
		Name:      "gauge",
		Help:      "Gauge of Kafka Client Writer IO",
	}, []string{LblType, LblRule, LblOp})
)

func init() {
	prometheus.MustRegister(KafkaWriterBatchCounter)
	prometheus.MustRegister(KafkaWriterErrCounter)
	prometheus.MustRegister(KafkaWriterBatchDurationHist)
	prometheus.MustRegister(KafkaWriterBatchTotalBytes)
	prometheus.MustRegister(KafkaWriterBatchGauge)
}
