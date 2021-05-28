package main

import (
	"context"
	"github.com/mercari/go-circuitbreaker"
	"log"
)

type patient struct {
	name string
	age  int
}

func fetchPatientInfo(ctx context.Context, name string) (*patient, error) {
	return &patient{name: name, age: 30}, nil
}

func main() {
	cb := circuitbreaker.New(nil)
	ctx := context.Background()

	data, err := cb.Do(context.Background(), func() (interface{}, error) {
		user, err := fetchPatientInfo(ctx, "Trump")
		if err != nil && err.Error() == "UserNoFound" {
			return nil, circuitbreaker.Ignore(err)
		}
		return user, err
	})

	if err != nil {
		log.Fatalf("failed to fetch patient:%s\n", err.Error())
	}
	log.Printf("fetched patient:%+v\n", data.(*patient))
}
