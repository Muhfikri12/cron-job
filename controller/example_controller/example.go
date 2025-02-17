package examplecontroller

import (
	"cron-job/service"

	"go.uber.org/zap"
)

type ExampleController interface {
}

type exampleController struct {
	service *service.AllService
	log     *zap.Logger
}

func NewExampleController(service *service.AllService, log *zap.Logger) ExampleController {
	return &exampleController{service, log}
}
