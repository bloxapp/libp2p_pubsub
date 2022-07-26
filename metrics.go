package pubsub

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricsPubsubIncoming = promauto.NewCounter(prometheus.CounterOpts{
		Name: "libp2p:pubsub:incoming",
		Help: "Count incoming messages",
	})
	metricsPubsubIncomingFull = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "libp2p:pubsub:incoming:full",
		Help: "Indicates that incoming channel is full",
	})
)

func init() {
	_ = prometheus.Register(metricsPubsubIncoming)
}
