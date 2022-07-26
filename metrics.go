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
	metricsPubsubIncomingFull = promauto.NewCounter(prometheus.CounterOpts{
		Name: "libp2p:pubsub:incoming:full",
		Help: "Count the amount of times the incoming queue is full",
	})
)

func init() {
	_ = prometheus.Register(metricsPubsubIncoming)
}
