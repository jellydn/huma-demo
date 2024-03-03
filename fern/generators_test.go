package fern

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/your-username/repo-name/sdk"
)

func TestGenerators(t *testing.T) {
	// Test the updated versions of the generators in generators.yml

	// Test fernapi/fern-typescript-node-sdk
	t.Run("TestTypescriptNodeSDK", func(t *testing.T) {
		// Create a new instance of the generator
		generator := sdk.NewTypescriptNodeSDKGenerator()

		// Test the version
		assert.Equal(t, "0.9.5", generator.GetVersion())

		// Test the output location and path
		output := generator.GetOutput()
		assert.Equal(t, "local-file-system", output.Location)
		assert.Equal(t, "../generated/typescript", output.Path)
	})

	t.Run("TestTypescriptBrowserSDKV1_0_0", func(t *testing.T) {
		// Create a new instance of the generator
		generator := sdk.NewTypescriptBrowserSDKGenerator()

		// Test the version
		assert.Equal(t, "1.0.0", generator.GetVersion())

		// Test the output location and path
		output := generator.GetOutput()
		assert.Equal(t, "local-file-system", output.Location)
		assert.Equal(t, "../generated/typescript-browser", output.Path)
	})
	t.Run("TestTypescriptBrowserSDK", func(t *testing.T) {
		// Create a new instance of the generator
		generator := sdk.NewTypescriptBrowserSDKGenerator()

		// Test the version
		assert.Equal(t, "0.9.5", generator.GetVersion())

		// Test the output location and path
		output := generator.GetOutput()
		assert.Equal(t, "local-file-system", output.Location)
		assert.Equal(t, "../generated/typescript-browser", output.Path)
	})

	t.Run("TestPythonSDKV1_0_0", func(t *testing.T) {
		// Create a new instance of the generator
		generator := sdk.NewPythonSDKGenerator()

		// Test the version
		assert.Equal(t, "1.0.0", generator.GetVersion())

		// Test the output location and path
		output := generator.GetOutput()
		assert.Equal(t, "local-file-system", output.Location)
		assert.Equal(t, "../generated/python", output.Path)
	})
	t.Run("TestPythonSDK", func(t *testing.T) {
		// Create a new instance of the generator
		generator := sdk.NewPythonSDKGenerator()

		// Test the version
		assert.Equal(t, "0.8.1", generator.GetVersion())

		// Test the output location and path
		output := generator.GetOutput()
		assert.Equal(t, "local-file-system", output.Location)
		assert.Equal(t, "../generated/python", output.Path)
	})
}
