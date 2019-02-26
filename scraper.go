package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	unknown    float64 = -1
	standalone float64 = 0
	leader     float64 = 1
	follower   float64 = 2
)

func getState(value string) float64 {
	switch value {
	case "standalone":
		return standalone
	case "leader":
		return leader
	case "follower":
		return follower
	default:
		return unknown
	}
}

// open tcp connections to zk nodes, send 'mntr' and return result as a map
func getRawMetrics() (list []string, err error) {
	// open connection
	timeout := time.Duration(*zkTimeout) * time.Second
	dialer := net.Dialer{Timeout: timeout}
	connection, err := dialer.Dial("tcp", *zkAddress)
	defer connection.Close()

	if err != nil {
		log.Printf("warning: cannot connect to %s: %v", *zkAddress, err)
	}

	_, err = connection.Write([]byte("mntr"))
	errFatal(err)

	// read response
	res, err := ioutil.ReadAll(connection)
	errFatal(err)

	// get slice of strings from response, like 'zk_avg_latency 0'
	lines := strings.Split(string(res), "\n")

	return lines, nil
}

func parseMetrics(lines []string) (map[string]float64, map[string]string) {
	metrics := map[string]float64{}
	labels := map[string]string{}

	metrics["zk_serving_requests"] = 0

	// skip instance if it in a leader only state and doesnt serving client requets
	if lines[0] == "This ZooKeeper instance is not currently serving requests" {
		return metrics, labels
	}

	metrics["zk_serving_requests"] = 1

	rawMetrics := map[string]string{}
	re := regexp.MustCompile(`(.+)\t(.+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}

		key := match[1]
		value := match[2]
		rawMetrics[key] = value
	}

	labels["zk_version"] = strings.Split(rawMetrics["zk_version"], "-")[0]
	labels["zk_instance"] = *zkAddress
	delete(rawMetrics, "zk_version")

	metrics["zk_server_state"] = getState(rawMetrics["zk_server_state"])
	delete(rawMetrics, "zk_server_state")

	for metricName, metric := range rawMetrics {
		rawValue := metrics[metricName]

		value, err := strconv.ParseFloat(metric, 64)
		if err == nil {
			metrics[metricName] = value
		} else {
			log.Printf("metric=%+v, value=%+v, error=%+v\n", metricName, rawValue, err)
		}
	}

	return metrics, labels
}

func scrapMetrics() {
	for {
		//log.Print("start scraping")

		rawMetrics, err := getRawMetrics()
		errFatal(err)
		metrics, labels := parseMetrics(rawMetrics)

		for metricName, metric := range promMetrics {
			value := metrics[metricName]
			metric.With(prometheus.Labels(labels)).Set(value)
		}

		time.Sleep(time.Duration(*scrapeInterval) * time.Second)
	}
}
