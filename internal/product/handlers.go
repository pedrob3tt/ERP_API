package product

import (
	"net/http"

	"github.com/pedrob3tt/GoApi/internal/json"
)

type handler struct {
	service service
}

func NewHandler(service service) *handler {
	return &handler{
		service: service,
		//Handler constructor
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. call the service to get the products
	// 2. write the JSON response
	err := h.service.ListProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.Write(w, http.StatusRequestTimeout, map[string]string{"error": "request timed out"})
		return
	}

	products := h.service.ListProducts(r.Context())
	json.Write(w, http.StatusOK, products)

}
