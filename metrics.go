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
	LblRange              = "range"
)

const (
	LblRange1to10    = "<=10"
	LblRange10to20   = "10-20"
	LblRange20to40   = "20-40"
	LblRange40to80   = "40-80"
	LblRange80to160  = "80-160"
	LblRange160to300 = "160-300"
	LblRange300to500 = "300-500"
	LblRange500toInf = ">500"
)

const (
	Kb              = 1024
	Mb              = 1024 * Kb
	LblRange1Kb     = "<=1kb"
	LblRange5Kb     = "1-5kb"
	LblRange10Kb    = "5-10kb"
	LblRange20Kb    = "10-20kb"
	LblRange50Kb    = "20-50kb"
	LblRange100Kb   = "50-100kb"
	LblRange256Kb   = "100-256kb"
	LblRange512Kb   = "256-512kb"
	LblRangeMb      = "512kb-1Mb"
	LblRangeMbToInf = ">1Mb"
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

	KafkaWriterBatchRangeCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "kafka_client",
		Subsystem: "writer_batch",
		Name:      "range_counter",
		Help:      "range counter of Kafka Client Batch Writer IO",
	}, []string{LblType, LblRange, LblRule, LblOp})
)

func init() {
	prometheus.MustRegister(KafkaWriterBatchCounter)
	prometheus.MustRegister(KafkaWriterErrCounter)
	prometheus.MustRegister(KafkaWriterBatchDurationHist)
	prometheus.MustRegister(KafkaWriterBatchTotalBytes)
	prometheus.MustRegister(KafkaWriterBatchRangeCounter)
}
