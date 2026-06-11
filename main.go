package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var requestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisicoes recebidas",
	},
)

var serviceUp = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "service_up",
		Help: "Disponibilidade do servico",
	},
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(serviceUp)

	serviceUp.Set(1)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
}

func handler(w http.ResponseWriter, r *http.Request) {

	requestsTotal.Inc()

	resp := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {

	http.HandleFunc("/projeto-korp", handler)

	http.HandleFunc("/health", healthHandler)
	// endpoint prometheus
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
