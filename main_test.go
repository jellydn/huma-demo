package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

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

func TestMain(t *testing.T) {
    // Capture stdout to a buffer
    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
    
    // Call the main function in a goroutine so it doesn't block
    go main()

    // Wait a short time for server to start
    time.Sleep(100 * time.Millisecond)
    
    // Make a request to the health check endpoint
    resp, err := http.Get("http://localhost:8888/")
    if err != nil {
        t.Fatalf("Health check request failed: %v", err)
    }
    defer resp.Body.Close()

    // Verify the response
    if resp.StatusCode != 200 {
        t.Errorf("Health check returned status %d, expected 200", resp.StatusCode)
    }
    
    // Check that expected output was written to stdout 
    w.Close()
    os.Stdout = old
    var buf bytes.Buffer
    io.Copy(&buf, r)
    output := buf.String()
    if !strings.Contains(output, "Starting server on port") {
        t.Errorf("Unexpected stdout output:\n%s", output)
    }
}
