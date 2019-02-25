package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	location       = flag.String("location", "/metrics", "metrics location")
	listen         = flag.String("listen", "0.0.0.0:8080", "address to listen on")
	zkAddress      = flag.String("zk-address", "127.0.0.1:2181", "zookeeper address")
	zkTimeout      = flag.Int64("zk-timeout", 120, "timeout for connection to zk servers, in seconds")
	scrapeInterval = flag.Int64("scrape-interval", 1, "scrape interval, in seconds")
)

func errFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func init() {
	flag.Parse()

	for _, metric := range promMetrics {
		prometheus.MustRegister(metric)
	}

}

func main() {
	go scrapMetrics()

	log.Printf("starting serving metrics at %s%s", *listen, *location)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listen, nil))

}
