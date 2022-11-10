package main

import (
	"context"
	"food-delivery-service/component/asyncjob"
	"log"
	"time"
)

func main() {
	j1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Job 1")
		return nil
	}, asyncjob.WithName("J1"), asyncjob.WithRetriesDuration([]time.Duration{time.Second * 5}))

	j2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("Job 2")
		return nil
	}, asyncjob.WithName("J2"), asyncjob.WithRetriesDuration([]time.Duration{time.Second * 5}))

	jm := asyncjob.NewGroup(true, j1, j2)

	if err := jm.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
