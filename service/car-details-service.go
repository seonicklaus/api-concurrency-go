package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seonicklaus/api-concurrency-go/entity"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type service struct{}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.CarDetails {
	// goroutine call endpoint 1: https://myfakeapi.com/api/cars/1
	go carService.FetchData()

	// goroutine call endpoint 2: https://myfakeapi.com/api/users/1
	go ownerService.FetchData()

	// create Car channel to get the data from endpoint 1
	car, _ := getCarData()
	// create User channel to get the data from endpoint 2
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		log.Println(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		log.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
