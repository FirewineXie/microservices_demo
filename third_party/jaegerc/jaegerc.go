package jaegerc

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
	"io"

	"time"
)

type TraceConf struct {
	ServerName string `json:"server_name"`
	CollectorEndpoint  string `json:"collector_endpoint"`
	LocalAgentHostPort string `json:"local_agent_host_port"`
}

func InitGlobalTracerProd(conf *TraceConf, logger *zap.Logger) (io.Closer, error) {

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint: "http://192.168.20.190:14268/api/traces",
			LocalAgentHostPort: "192.168.20.190:6831",
		},
	}


	jLogger := zapLogger{
		logger: logger,
	}

	jMetricsFactory := prometheus.New()

	// Initialize tracer with a logger and a metrics factory
	return cfg.InitGlobalTracer(
		conf.ServerName,
		jaegercfg.Logger(&jLogger),
		//jaegercfg.Logger(&jLogger),
		jaegercfg.Metrics(jMetricsFactory))

}
