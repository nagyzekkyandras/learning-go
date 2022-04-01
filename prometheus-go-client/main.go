package main

import(
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

var criticismCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "aaa_criticism",
	Help: "Total number of criticism.",

})

func main(){
	criticismCounter.Inc()
	fmt.Println("running")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":5000", nil))
}