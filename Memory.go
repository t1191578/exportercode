package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type MemoryMetrics struct {
	alloc float64
	//idle  float64
	//other float64
	//total float64
}

func MemoryGetMetrics() *MemoryMetrics {
	return ParseMemoryMetrics(MemoryData())
}

 

func ParseMemoryMetrics(input []byte) *MemoryMetrics {
	var cm MemoryMetrics
            //if strings.Contains(string(input), "/") {
	    splitted := strings.Split(strings.TrimSpace(string(input)), "/")
           

for i := 0; i < len(splitted); i++ {
val, _ := strconv.ParseFloat(strings.TrimSpace(splitted[i]), 64)	  

		cm.alloc += val
	}
	
	  //  cm.idle, _  = strconv.ParseFloat(splitted[1], 64)
	   // cm.other, _ = strconv.ParseFloat(splitted[2], 64)
	    //cm.total, _ = strconv.ParseFloat(splitted[3], 64)
	//}
	return &cm
}
//ecute the sinfo command and return its output
func MemoryData() []byte {
	cmd := exec.Command("sinfo", "-h", "-o %m/", "-N", "|xargs")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	out, _ := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return out
}

/*
 * Implement the Prometheus Collector interface and feed the
 * Slurm scheduler metrics into it.
 * https://godoc.org/github.com/prometheus/client_golang/prometheus#Collector
 */

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{
		alloc: prometheus.NewDesc("slurm_Memory_alloc", "Allocated Memory", nil, nil),
	//	idle:  prometheus.NewDesc("slurm_Memory_idle", "Idle Memory", nil, nil),
	//	other: prometheus.NewDesc("slurm_Memory_other", "Mix Memory", nil, nil),
	///	total: prometheus.NewDesc("slurm_Memory_total", "Total Memory", nil, nil),
	}
}

type MemoryCollector struct {
	alloc *prometheus.Desc
	//idle  *prometheus.Desc
//	other *prometheus.Desc
//	total *prometheus.Desc
}

// Send all metric descriptions
func (cc *MemoryCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cc.alloc
//	ch <- cc.idle
//	ch <- cc.other
//	ch <- cc.total
}
func (cc *MemoryCollector) Collect(ch chan<- prometheus.Metric) {
	cm := MemoryGetMetrics()
	ch <- prometheus.MustNewConstMetric(cc.alloc, prometheus.GaugeValue, cm.alloc)
//	ch <- prometheus.MustNewConstMetric(cc.idle,  prometheus.GaugeValue, cm.idle)
//	ch <- prometheus.MustNewConstMetric(cc.other, prometheus.GaugeValue, cm.other)
//	ch <- prometheus.MustNewConstMetric(cc.total, prometheus.GaugeValue, cm.total)
}
