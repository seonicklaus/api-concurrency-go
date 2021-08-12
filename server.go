package main

import (
	"github.com/seonicklaus/api-concurrency-go/controller"
	router "github.com/seonicklaus/api-concurrency-go/http"
	"github.com/seonicklaus/api-concurrency-go/service"
)

var (
	carDetailsService    = service.NewCarDetailsService()
	carDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           = router.NewChiRouter()
)

func main() {
	const port = "8000"
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
