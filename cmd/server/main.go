package main

import (
  "os"
  "os/signal"
  "github.com/joerodriguez/cf-application-metrics/pkg/metrics"
  "github.com/joerodriguez/cf-application-metrics/pkg/metron"
  "github.com/joerodriguez/cf-application-metrics/pkg/processor"
  "github.com/joerodriguez/cf-application-metrics/pkg/appcontainerinfo"
  "github.com/joerodriguez/cf-application-metrics/pkg/opentsdb"
  "github.com/joerodriguez/cf-application-metrics/pkg/server"
)

func main() {
  metrics := make(chan *metrics.Metric)


  m := metron.New()
  p := processor.New(metrics, m)
  drainEgress := p.Start()

  a := appcontainerinfo.New()
  s := server.New(metrics, a, opentsdb.NewFactory())
  stopIngress := s.Start()

  defer func() {
    stopIngress()
    drainEgress()
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill)
  <-c
}

