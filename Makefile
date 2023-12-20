.PHONY: dev
dev:
	@echo "Running dev..."
	@go run .

# Update project dependencies
.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-coverage
test-coverage:
	@go test -v -cover ./...

.PHONY: test-report
test-report:
	@echo "Cleaning up old report..."
	@rm -rf report
	@echo "Running e2e tests and generating report..."
	@echo "Make sure to run 'make dev' in a separate terminal window"
	@hurl --test tests/*.hurl --report-html report

.PHONY: view-test-report
view-test-report:
	@echo "View test report..."
	@bunx serve report -p 3456

.PHONY: generate-openapi
generate-openapi:
	@echo "Generating OpenAPI..."
	@restish :8888/openapi.yaml > openapi.yaml
	@echo "Adding OpenAPI to Fern..."
	@rm -rf .fern
	@fern init --openapi openapi.yaml
	@fern add fernapi/fern-typescript-browser-sdk
	@fern add fernapi/fern-python-sdk
	@fern add fernapi/fern-go-sdk

.PHONY: generate-sdk-docker
generate-sdk-docker:
	@echo "Generating SDK..."
	@rm -rf generated
	@fern generate --local

.PHONY: generate-sdk
generate-sdk:
	@echo "Generating SDK..."
	@rm -rf generated
	@fern generate

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  dev                Run the application in development mode"
	@echo "  tidy               Update project dependencies"
	@echo "  test               Run unit tests"
	@echo "  test-coverage      Run unit tests with coverage"
	@echo "  test-report        Run end-to-end tests and generate a report"
	@echo "  view-test-report   View the generated test report"
	@echo "  generate-openapi   Generate OpenAPI specification and add it to Fern"
	@echo "  generate-sdk-docker Generate SDK using Docker"
	@echo "  generate-sdk       Generate SDK"
	@echo "  help               Show this help message"
