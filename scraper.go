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
	unknown    = "-1"
	standalone = "0"
	leader     = "1"
	follower   = "2"
)

func getState(value string) string {
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
func getMetrics() (map[string]string, map[string]string) {
	metrics := map[string]string{}
	labels := map[string]string{}

	metrics["zk_serving_requests"] = "0"

	// open connection
	timeout := time.Duration(*zkTimeout) * time.Second
	dialer := net.Dialer{Timeout: timeout}
	connection, err := dialer.Dial("tcp", *zkAddress)
	defer connection.Close()

	if err != nil {
		log.Printf("warning: cannot connect to %s: %v", *zkAddress, err)
		return metrics, labels
	}

	_, err = connection.Write([]byte("mntr"))
	errFatal(err)

	// read response
	res, err := ioutil.ReadAll(connection)
	errFatal(err)

	// get slice of strings from response, like 'zk_avg_latency 0'
	lines := strings.Split(string(res), "\n")

	// skip instance if it in a leader only state and doesnt serving client requets
	if lines[0] == "This ZooKeeper instance is not currently serving requests" {
		return metrics, labels
	}

	metrics["zk_serving_requests"] = "1"

	re := regexp.MustCompile(`(.+)\t(.+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) != 3 {
			continue
		}

		key := match[1]
		value := match[2]
		metrics[key] = value
	}

	labels["zk_version"] = strings.Split(metrics["zk_version"], "-")[0]
	labels["zk_instance"] = *zkAddress
	delete(metrics, "zk_version")

	metrics["zk_server_state"] = getState(metrics["zk_server_state"])

	

	return metrics, labels
}

func scrapMetrics() {
	for {
		//log.Print("start scraping")

		metrics, labels := getMetrics()

		for metricName, metric := range promMetrics {
			rawValue := metrics[metricName]

			value, err := strconv.ParseFloat(rawValue, 64)
			if err == nil {
				metric.With(prometheus.Labels(labels)).Set(value)
			} else {
				log.Printf("metric=%+v, value=%+v, error=%+v\n", metricName, rawValue, err)
			}
		}

		time.Sleep(time.Duration(*scrapeInterval) * time.Second)
	}
}
