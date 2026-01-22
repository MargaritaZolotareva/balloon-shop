package metrics

import prometheus "github.com/prometheus/client_golang/prometheus"

var (
	SuccessCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_success_count",
			Help: "Total number of successful API requests",
		},
		[]string{"status"},
	)

	Error400Counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_error_400_count",
			Help: "Total number of 400 errors (client errors)",
		},
		[]string{"status"},
	)

	Error404Counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_error_404_count",
			Help: "Total number of 404 errors (not found)",
		},
		[]string{"status"},
	)

	Error500Counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_error_500_count",
			Help: "Total number of 500 errors (server errors)",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(SuccessCounter)
	prometheus.MustRegister(Error400Counter)
	prometheus.MustRegister(Error404Counter)
	prometheus.MustRegister(Error500Counter)
}
