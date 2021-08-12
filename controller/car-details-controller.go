package controller

import (
	"encoding/json"
	"net/http"

	"github.com/seonicklaus/api-concurrency-go/service"
)

var (
	carDetailsService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
