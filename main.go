package main

import (
	"cron-job/infra"
	"cron-job/router"
	orderservice "cron-job/service/order_service"
	"log"
)

func main() {

	ctx, err := infra.NewIntegrateContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := router.NewRoutes(*ctx)

	orderservice.StartCronJob(ctx.DB, ctx.Log)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
