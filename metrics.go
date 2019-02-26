package main

import "github.com/prometheus/client_golang/prometheus"

var promMetrics = map[string]prometheus.GaugeVec{
	"zk_approximate_data_size": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_approximate_data_size",
		Help: "zk_approximate_data_size.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_avg_latency": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_avg_latency",
		Help: "zk_avg_latency.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_ephemerals_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_ephemerals_count",
		Help: "zk_ephemerals_count.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_fsync_threshold_exceed_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_fsync_threshold_exceed_count",
		Help: "zk_fsync_threshold_exceed_count.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_max_file_descriptor_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_max_file_descriptor_count",
		Help: "zk_max_file_descriptor_count.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_max_latency": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_max_latency",
		Help: "zk_max_latency.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_min_latency": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_min_latency",
		Help: "zk_min_latency.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_num_alive_connections": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_num_alive_connections",
		Help: "zk_num_alive_connections.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_open_file_descriptor_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_open_file_descriptor_count",
		Help: "zk_open_file_descriptor_count.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_outstanding_requests": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_outstanding_requests",
		Help: "zk_outstanding_requests.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_packets_received": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_packets_received",
		Help: "zk_packets_received.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_packets_sent": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_packets_sent",
		Help: "zk_packets_sent.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_server_state": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_server_state",
		Help: "zk_server_state.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_serving_requests": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_serving_requests",
		Help: "zk_serving_requests.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_watch_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_watch_count",
		Help: "zk_watch_count.",
	}, []string{"zk_version", "zk_instance"}),
	"zk_znode_count": *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "zk_znode_count",
		Help: "zk_znode_count.",
	}, []string{"zk_version", "zk_instance"}),
}
