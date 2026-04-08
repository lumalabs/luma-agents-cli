// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/luma-agents-cli/internal/mocktest"
	"github.com/stainless-sdks/luma-agents-cli/internal/requestflag"
)

func TestGenerationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"generations", "create",
			"--prompt", "A glass of iced coffee on a marble countertop, morning light streaming through a window",
			"--aspect-ratio", "3:1",
			"--image-ref", "{data: data, media_type: media_type, url: url}",
			"--model", "model",
			"--output-format", "png",
			"--source", "{data: data, media_type: media_type, url: url}",
			"--style", "auto",
			"--type", "image",
			"--web-search=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(generationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"generations", "create",
			"--prompt", "A glass of iced coffee on a marble countertop, morning light streaming through a window",
			"--aspect-ratio", "3:1",
			"--image-ref.data", "data",
			"--image-ref.media-type", "media_type",
			"--image-ref.url", "url",
			"--model", "model",
			"--output-format", "png",
			"--source.data", "data",
			"--source.media-type", "media_type",
			"--source.url", "url",
			"--style", "auto",
			"--type", "image",
			"--web-search=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prompt: >-\n" +
			"  A glass of iced coffee on a marble countertop, morning light streaming through\n" +
			"  a window\n" +
			"aspect_ratio: '3:1'\n" +
			"image_ref:\n" +
			"  - data: data\n" +
			"    media_type: media_type\n" +
			"    url: url\n" +
			"model: model\n" +
			"output_format: png\n" +
			"source:\n" +
			"  data: data\n" +
			"  media_type: media_type\n" +
			"  url: url\n" +
			"style: auto\n" +
			"type: image\n" +
			"web_search: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--auth-token", "string",
			"generations", "create",
		)
	})
}

func TestGenerationsGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"generations", "get",
			"--generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
