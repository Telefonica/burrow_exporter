package burrow_exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	KafkaConsumerPartitionLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_lag",
			Help: "The lag of the consumer on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition", "status"},
	)
	KafkaConsumerPartitionLastOffsetLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_last_offset_lag",
			Help: "The lag of the latest offset commit on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition", "status"},
	)
	KafkaConsumerPartitionCurrentOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_current_offset",
			Help: "The latest offset commit on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition", "status"},
	)
	KafkaConsumerPartitionMaxOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_partition_max_offset",
			Help: "The log end offset on a partition as reported by burrow.",
		},
		[]string{"cluster", "group", "topic", "partition", "status"},
	)
	KafkaConsumerTotalLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_total_lag",
			Help: "The total amount of lag for the consumer group as reported by burrow",
		},
		[]string{"cluster", "group"},
	)
	KafkaTopicPartitionOffset = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_burrow_topic_partition_offset",
			Help: "The latest offset on a topic's partition as reported by burrow.",
		},
		[]string{"cluster", "topic", "partition"},
	)
)

func init() {
	prometheus.MustRegister(KafkaConsumerPartitionLag)
	prometheus.MustRegister(KafkaConsumerPartitionLastOffsetLag)
	prometheus.MustRegister(KafkaConsumerPartitionCurrentOffset)
	prometheus.MustRegister(KafkaConsumerPartitionMaxOffset)
	prometheus.MustRegister(KafkaConsumerTotalLag)
	prometheus.MustRegister(KafkaTopicPartitionOffset)
}
