package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceURL = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the URL %s", carServiceURL)

	// Get the external API
	resp, _ := client.Get(carServiceURL)

	// Write response to channel (TODO)
	carDataChannel <- resp
}
