package bar

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/paullockaby/test-go/internal/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

type Config struct {
	Verbose              bool
	ListenerHost         string
	ListenerPort         int
	HideHealthAccessLogs bool
	HideNormalAccessLogs bool
	Options              *viper.Viper
}

var (
	logger = logging.Log

	healthCheckTimeout time.Duration

	requestProcessingTimeHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "foo_processing_time",
			Help: "Histogram of processing time for request handling",
		},
		[]string{"endpoint", "status_code"},
	)
)

func Run(config Config) error {
	var err error

	prometheus.MustRegister(requestProcessingTimeHistogram)

	// this is where healthcheck measurements will be sent
	healthCheckTimeout = config.Options.GetDuration("listener.health.timeout")

	// make sure the timeouts are greater than zero
	if healthCheckTimeout <= 0 {
		return fmt.Errorf("listener.health.timeout must be greater than zero")
	}

	// set up an access logger and wrap our server with it
	accessLogger := log.New(os.Stderr, "", log.LstdFlags)
	mux := http.NewServeMux()

	// internal monitoring endpoint
	mux.Handle("/_/health", loggingMiddleware(healthHandler, accessLogger, config))

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.ListenerHost, config.ListenerPort),
		ReadTimeout:  config.Options.GetDuration("listener.http.read_timeout"),
		WriteTimeout: config.Options.GetDuration("listener.http.write_timeout"),
		IdleTimeout:  config.Options.GetDuration("listener.http.idle_timeout"),
		Handler:      mux,
	}

	logger.Info(fmt.Sprintf("listening on %s:%d", config.ListenerHost, config.ListenerPort))
	if err = server.ListenAndServe(); err != nil {
		logger.Error(fmt.Sprintf("failed to start server: %s", err))
		return err
	}
	return nil
}

/*
###
### add a logging middleware handler
###
*/

type responseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWithStatus) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.HandlerFunc, logger *log.Logger, config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// use a response writer wrapper to capture the status code
		wrappedWriter := &responseWriterWithStatus{ResponseWriter: w, statusCode: http.StatusOK}

		// call the next handler in the chain (the actual request handler)
		next.ServeHTTP(wrappedWriter, r)

		showLogs := true
		successRequests := map[int]bool{
			http.StatusOK:        true,
			http.StatusCreated:   true,
			http.StatusAccepted:  true,
			http.StatusNoContent: true,
		}

		if config.HideNormalAccessLogs {
			if successRequests[wrappedWriter.statusCode] {
				showLogs = false
			}
		} else {
			if config.HideHealthAccessLogs && (r.URL.Path == "/_/health" || r.URL.Path == "/ping" || r.URL.Path == "/health") {
				if successRequests[wrappedWriter.statusCode] {
					showLogs = false
				}
			}
		}

		requestLength := time.Since(start)
		if showLogs {
			// after the handler completes, log the request details to stderr
			logger.Printf("%s - %s %s %s %d %s",
				r.RemoteAddr,
				r.Method,
				r.URL.Path,
				r.Proto,
				wrappedWriter.statusCode,
				requestLength,
			)
		}

		// then record how long it took
		requestProcessingTimeHistogram.WithLabelValues(r.URL.Path, fmt.Sprintf("%d", wrappedWriter.statusCode)).Observe(requestLength.Seconds())
	}
}
