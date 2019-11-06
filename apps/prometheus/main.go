package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
		},
	})
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "manny_http_request_count",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)


func rootHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from example application."))
}


func main(){
	var port string
	flag.StringVar(&port, "port", "9090", "")

	flag.Parse()

	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)

	http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestsTotal, http.HandlerFunc(rootHandler)))
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	log.Println("starting web server")

	log.Println("serving on ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


