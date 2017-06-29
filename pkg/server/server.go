package server

import "github.com/joerodriguez/cf-application-metrics/pkg/metrics"

type appContainerInfo interface {

}

type implementationFactory interface {

}

type Server struct {
  metrics chan *metrics.Metric
  appContainerInfo appContainerInfo
  factories []implementationFactory
}

func New (
  metrics chan *metrics.Metric,
  appContainerInfo appContainerInfo,
  factories... implementationFactory,
) *Server {
  return &Server{
    metrics: metrics,
    appContainerInfo: appContainerInfo,
    factories: factories,
  }
}

func (s *Server) Start() func() {

}