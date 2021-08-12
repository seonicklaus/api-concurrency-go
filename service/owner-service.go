package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the URL %s", ownerServiceURL)

	// Get the external API
	resp, _ := client.Get(ownerServiceURL)

	// Write response to channel (TODO)
	ownerDataChannel <- resp
}
