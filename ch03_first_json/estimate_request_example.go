package main

import (
	"time"
)

var EstimateRequestExample1 = func() *EstimateRequest {
	clientId := int64(1)
	productId := int64(2)
	return &EstimateRequest{
		ClientID:    &clientId,
		ClientName:  "John Doe",
		RequestedAt: time.Now(),
		Items: EstimateRequestItems{
			{
				ProductID:   &productId,
				ProductName: "Apple",
				Amount:      7,
			},
			{
				ProductName: "Orange",
				Amount:      11,
			},
		},
	}
}()
