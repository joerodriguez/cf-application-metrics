package processor

import "github.com/joerodriguez/cf-application-metrics/pkg/metrics"

type Processor struct {
  metrics chan *metrics.Metric
  metron metron
}

type metron interface {

}

func New(metrics chan *metrics.Metric, metron metron) *Processor {
  return &Processor{
    metrics: metrics,
    metron: metron,
  }
}

func (p *Processor) Start() func() {

}