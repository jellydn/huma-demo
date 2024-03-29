<h1 align="center">Welcome to huma-demo 👋</h1>
<p>
  Welcome to huma-demo, a modern, simple, fast, and flexible micro-framework for building HTTP REST/RPC APIs in Golang. It is backed by OpenAPI 3 and JSON Schema, providing a robust and standardized approach to API development.
</p>

[![Exploring Huma with IT Man: A Journey Through Modern Micro-Frameworks](https://i.ytimg.com/vi/1iReyppMxXk/hqdefault.jpg)](https://www.youtube.com/watch?v=1iReyppMxXk)

## 🏠 [Homepage](https://huma.rocks/)

### ✨ [Demo](https://huma-demo.fly.dev/docs)

## Prerequisites

Before getting started with huma-demo, make sure you have the following tools installed:

- [The Go Programming Language](https://go.dev/) (version 1.20 or higher)
- [Bun](https://bun.sh/) - A fast all-in-one JavaScript runtime
- [OrbStack](https://orbstack.dev/) - Fast, light, simple Docker & Linux on macOS
- [Hurl](https://hurl.dev/) - Run and Test HTTP Requests

Make sure these tools are installed and properly configured on your system before proceeding.

## Usage

To get started with huma-demo, simply run the following command:

```sh
make dev
```

[![Run dev](https://i.gyazo.com/e79c264de5566902a5d8e7ff7619d59c.gif)](https://gyazo.com/e79c264de5566902a5d8e7ff7619d59c)

This will start the development server and allow you to test your API locally.

## Test

To run the tests for huma-demo, you can use the following commands:

```sh
make test
make test-report # Run e2e tests with hurl and generate a report
make test-coverage
```

[![Run test](https://i.gyazo.com/2af6fbe3c5afb6cc93725dcaa5b8267f.gif)](https://gyazo.com/2af6fbe3c5afb6cc93725dcaa5b8267f)

These commands will run the unit tests and generate a coverage report, respectively.

## More usages

To view the available commands and their usage, you can use the following command:

```sh
make help
```

This will display a list of targets and their descriptions, helping you navigate the project.

```sh
Usage: make [target]

Targets:
  dev                Run the application in development mode
  tidy               Update project dependencies
  test               Run unit tests
  test-coverage      Run unit tests with coverage
  test-report        Run end-to-end tests and generate a report
  view-test-report   View the generated test report
  generate-openapi   Generate OpenAPI specification and add it to Fern
  generate-fern      Generate Fern
  generate-go-sdk    Generate Go SDK
  help               Show this help message
```

## Resources

- [Installation - Huma](https://huma.rocks/tutorial/installation/)
- [APIs in Go with Huma 2.0](https://dgt.hashnode.dev/apis-in-go-with-huma-20)
- [fern-api/fern: SDKs and Docs for your API](https://github.com/fern-api/fern)
- [A CLI for REST APIs - DEV Community](https://dev.to/danielgtaylor/a-cli-for-rest-apis-part-1-104b)

## Author

👤 **Huynh Duc Dung**

- Website: https://productsway.com/
- Twitter: [@jellydn](https://twitter.com/jellydn)
- Github: [@jellydn](https://github.com/jellydn)

## Show your support

Give a ⭐️ if this project helped you!

[![kofi](https://img.shields.io/badge/Ko--fi-F16061?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/dunghd)
[![paypal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/dunghd)
[![buymeacoffee](https://img.shields.io/badge/Buy_Me_A_Coffee-FFDD00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/dunghd)
