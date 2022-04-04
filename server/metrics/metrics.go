package metrics

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

const (
	metricsRestartDelay = time.Second * 30
	metricsPort         = 9090
	metricsNamespace    = "kiss_server"
)

var (
	metricsRequestTimings = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricsNamespace,
		Name:      "request_timings",
		Help:      "Handling timings of gRPC requests received by the server",
		Buckets:   []float64{0, 0.05, 0.1, 0.5, 1, 2, 10},
	}, []string{"method"})

	metricsAuthTimings = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metricsNamespace,
		Name:      "request_auth_timings",
		Help:      "Runtime timings of verifying request payload signatures",
		Buckets:   []float64{0, 0.01, 0.05, 0.1, 0.5, 1},
	}, []string{"method"})

	metricsResponseCodes = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Name:      "response_codes",
		Help:      "Response codes of gRPC responses made by the server",
	}, []string{"method", "success", "code"})
)

func RegisterRequestTiming(method string, timing float64) {
	metricsRequestTimings.WithLabelValues(method).Observe(timing)
}

func RegisterAuthTiming(method string, timing float64) {
	metricsAuthTimings.WithLabelValues(method).Observe(timing)
}

func RegisterResponseCode(method string, success bool, code string) {
	metricsResponseCodes.WithLabelValues(
		method,
		strconv.FormatBool(success),
		code,
	).Add(1)
}

// InitMetricsServer initialises a Prometheus metrics server in the background
func InitMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", metricsPort),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// We delay for a fixed amount of time if metrics server has crashed before restarting.
	go func() {
		for {
			err := srv.ListenAndServe()
			if err != nil {
				log.Error().Msgf(
					"Metrics server encountered error, waiting %v before restarting: %v",
					err,
					metricsRestartDelay,
				)
				time.Sleep(metricsRestartDelay)
				log.Info().Msgf("Restaring metrics server...")
			}
		}
	}()
}
