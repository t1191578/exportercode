

package main

import (
  "flag"
  "net/http"
  "github.com/prometheus/common/log"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
  // Metrics have to be registered to be exposed
  prometheus.MustRegister(NewSchedulerCollector()) // from scheduler.go
  prometheus.MustRegister(NewQueueCollector())     // from queue.go
  prometheus.MustRegister(NewNodesCollector())     // from nodes.go
  prometheus.MustRegister(NewCPUsCollector())      // from cpus.go
  prometheus.MustRegister(NewMemoryCollector())      // from Memory.go
}

var listenAddress = flag.String(
  "listen-address",
  ":8080",
  "The address to listen on for HTTP requests.")

func main() {
  flag.Parse()
  // The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
  log.Infof("Starting Server: %s", *listenAddress)
  http.Handle("/metrics", promhttp.Handler())
  log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
