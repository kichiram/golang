package main

import (
  "fmt"
  "net/http"
  "time"
  "log"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
  testMetrics = promauto.NewCounterVec(
    prometheus.CounterOpts{
      Name: "a_http_request_count_total",
      Help: "Test Counter",
    },
    []string{"testlabel"},
  )
)

func newServer(addr string) (*http.ServeMux, *http.Server) {
  mux := http.NewServeMux()
  return mux, &http.Server{
    Addr:              addr,
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
  }
}

func handler1(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
  log.Print("Hello!\n")
  testMetrics.With(prometheus.Labels{"testlabel": "Hello"}).Inc()
}

func handler2(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "World")
  log.Print("World\n")
  testMetrics.With(prometheus.Labels{"testlabel": "World"}).Inc()
}

func main() {
  mux1, srv1 := newServer(":8080")
  mux2, srv2 := newServer(":8081")

  go func() {
     mux1.HandleFunc("/hello", handler1)
     mux1.HandleFunc("/world", handler2)
     srv1.ListenAndServe()
  }()

  mux2.Handle("/metrics", promhttp.Handler())
  srv2.ListenAndServe()
}
