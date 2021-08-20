package api

import (
	"DI/controllers"
	"fmt"
)

type HttpAPI struct {
	controller controllers.Controller
}

func (api HttpAPI) Start() {
	fmt.Println("API Start")

}

func NewHttpAPI(controller controllers.Controller) HttpAPI {
	return HttpAPI{controller: controller}
}
