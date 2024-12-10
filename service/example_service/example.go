package exampleservice

import (
	"cron-job/repository"

	"go.uber.org/zap"
)

type ExampleService interface {
}

type exampleService struct {
	Repo *repository.AllRepository
	Log  *zap.Logger
}

func NewExampleService(Repo *repository.AllRepository, Log *zap.Logger) ExampleService {
	return &exampleService{Repo, Log}
}
