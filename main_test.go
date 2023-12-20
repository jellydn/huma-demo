package main

import (
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestGetHealth(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Get("/")
	if !strings.Contains(resp.Body.String(), `"status":"ok"`) {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}

func TestGetGreeting(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Get("/greeting/itman")
	if !strings.Contains(resp.Body.String(), "Hello, itman!") {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}

func TestPutReview(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"author": "dunghd",
		"rating": 5,
	})

	if resp.Code != 201 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestPutReviewError(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"rating": 10,
	})

	if resp.Code != 422 {
		t.Fatalf("Unexpected status code: %d", resp.Code)
	}
}

func TestGetReviews(t *testing.T) {
	_, api := humatest.New(t)
	addRoutes(api)

	// Setup a review
	api.Post("/reviews", map[string]any{
		"author": "dunghd",
		"rating": 5,
	})

	resp := api.Get("/reviews")
	if !strings.Contains(resp.Body.String(), `"author":"dunghd"`) {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}
