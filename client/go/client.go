package main

import (
	"context"
	"fmt"

	"github.com/jellydn/huma-demo/sdk"
)

func main() {
	ctx := context.Background()

	// Initialize an SDK client.
	client, _ := sdk.NewClientWithResponses("http://localhost:8888")

	// Make the greeting request.
	greeting, err := client.GetGreetingWithResponse(ctx, "world")
	if err != nil {
		panic(err)
	}

	if greeting.StatusCode() > 200 {
		panic(greeting.ApplicationproblemJSONDefault)
	}

	// Everything was successful, so print the message.
	fmt.Println(greeting.JSON200.Message)
}
