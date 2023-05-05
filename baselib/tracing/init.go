

package tracing

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// TracingMe type;
type TracingMe struct {
	opentracing.Tracer
	io.Closer
}

var tracerInstance *TracingMe

// InitTracing func;
func InitTracing(serviceName string) *TracingMe {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		err := fmt.Errorf("ERROR: cannot init Jaeger: %v", err)
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)

	tracerInstance = &TracingMe{tracer, closer}
	return tracerInstance
}

// GetTracer func;
func GetTracer() *TracingMe {
	if tracerInstance == nil {
		err := fmt.Errorf("Tracing was not initialized")
		panic(err)
	}

	return tracerInstance
}

// LetClose func;
func (t *TracingMe) LetClose() error {
	return t.Closer.Close()
}
