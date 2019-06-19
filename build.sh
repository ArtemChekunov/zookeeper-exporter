#!/bin/bash -e
docker build -t ${1:-'artemchekunov/zookeeper-exporter:latest'} .
