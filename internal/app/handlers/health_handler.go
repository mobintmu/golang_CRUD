package handlers

import (
	"encoding/json"
	"net/http"

	"golang_crud/internal/pkg/health/service"
)

type healthHandler struct {
	healthService service.HealthService
}

func NewHealthHandler(healthService service.HealthService) *healthHandler {
	return &healthHandler{
		healthService: healthService,
	}
}

func (h *healthHandler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.healthCheck)
}

func (h *healthHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status, err := h.healthService.CheckHealth(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
