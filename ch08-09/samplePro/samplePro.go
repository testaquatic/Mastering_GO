package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PORT = ":1234"

var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "testaquatic",
		Name:      "my_counter",
		Help:      "This is my counter",
	},
)

var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "testaquatic",
		Name:      "my_gauge",
		Help:      "This is my gauge",
	},
)

var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "testaquatic",
		Name:      "my_histogram",
		Help:      "This is my histogram",
	},
)

var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "testaquatic",
		Name:      "my_summary",
		Help:      "This is my summary",
	},
)

func main() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on port", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
