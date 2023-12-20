package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// Options for the CLI.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

// GreetingInput represents the greeting operation request.
type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// ReviewInput represents the review operation request.
type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
	}
}

// Review represents a review.
type Review struct {
	Author  string `json:"author"`
	Rating  int    `json:"rating"`
	Message string `json:"message,omitempty"`
}

// Reviews represents the response of the "get all reviews" operation.
type ReviewsOutput struct {
	Body struct {
		Reviews []Review `json:"reviews"`
	}
}

// Health represents the response of the "get health" operation.
type HealthOutput struct {
	Body struct {
		Status string `json:"status"`
	}
}

func addRoutes(api huma.API) {
	// Add GET / for health checks
	huma.Register(api, huma.Operation{
		OperationID:   "get-health",
		Summary:       "Get health",
		Method:        http.MethodGet,
		Path:          "/",
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *struct{}) (*HealthOutput, error) {
		resp := &HealthOutput{}
		resp.Body.Status = "ok"
		return resp, nil
	})

	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Summary:     "Get a greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
	}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	reviews := make([]Review, 0)

	// Add a POST /reviews
	huma.Register(api, huma.Operation{
		OperationID:   "post-review",
		Summary:       "Post a review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *ReviewInput) (*struct{}, error) {
		review := Review{
			Author:  i.Body.Author,
			Rating:  i.Body.Rating,
			Message: i.Body.Message,
		}
		reviews = append(reviews, review)
		return nil, nil
	})

	// Add a GET /reviews
	huma.Register(api, huma.Operation{
		OperationID: "get-all-reviews",
		Summary:     "Get all reviews",
		Method:      http.MethodGet,
		Path:        "/reviews",
	}, func(ctx context.Context, i *struct{}) (*ReviewsOutput, error) {
		resp := &ReviewsOutput{}
		resp.Body.Reviews = reviews
		return resp, nil
	})

}

func main() {
	// Create a CLI app which takes a port option.
	cli := huma.NewCLI(func(hooks huma.Hooks, options *Options) {
		// Create a new router & API
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My Simple API", "1.0.0"))

		addRoutes(api)

		// Start the server!
		addr := "http://127.0.0.1:" + fmt.Sprintf("%d", options.Port)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			log.Printf("Go to %s/docs for documentation", addr)
			err := http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
