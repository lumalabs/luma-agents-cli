// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/lumalabs/luma-agents-cli/internal/mocktest"
	"github.com/lumalabs/luma-agents-cli/internal/requestflag"
)

func TestGenerationsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"generations", "create",
			"--prompt", "A glass of iced coffee on a marble countertop, morning light streaming through a window",
			"--aspect-ratio", "3:1",
			"--image-ref", "{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}",
			"--model", "uni-1",
			"--output-format", "png",
			"--source", "{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}",
			"--style", "auto",
			"--type", "image",
			"--user-id", "user_id",
			"--video", "{duration: 5s, edit: {auto_controls: true, controls: {depth: {blur: 0, enabled: true}, face: {enabled: true}, normals: {augmentation: 0, enabled: true}, pose: {enabled: true, strength: precise}, trajectory: {enabled: true, sparsity: 0}}, keyframe_indexes: [0], keyframes: [{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}], strength: adhere_1}, end_frame: {data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}, exr_export: true, hdr: true, keyframe_indexes: [0], keyframes: [{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}], loop: true, resolution: 360p, source_position: {h_norm: 1, w_norm: 1, x_norm: -2, y_norm: -2}, start_frame: {data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}}",
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
			"--image-ref.generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--image-ref.media-type", "media_type",
			"--image-ref.url", "url",
			"--model", "uni-1",
			"--output-format", "png",
			"--source.data", "data",
			"--source.generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--source.media-type", "media_type",
			"--source.url", "url",
			"--style", "auto",
			"--type", "image",
			"--user-id", "user_id",
			"--video.duration", "5s",
			"--video.edit", "{auto_controls: true, controls: {depth: {blur: 0, enabled: true}, face: {enabled: true}, normals: {augmentation: 0, enabled: true}, pose: {enabled: true, strength: precise}, trajectory: {enabled: true, sparsity: 0}}, keyframe_indexes: [0], keyframes: [{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}], strength: adhere_1}",
			"--video.end-frame", "{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}",
			"--video.exr-export=true",
			"--video.hdr=true",
			"--video.keyframe-indexes", "[0]",
			"--video.keyframes", "[{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}]",
			"--video.loop=true",
			"--video.resolution", "360p",
			"--video.source-position", "{h_norm: 1, w_norm: 1, x_norm: -2, y_norm: -2}",
			"--video.start-frame", "{data: data, generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, media_type: media_type, url: url}",
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
			"    generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    media_type: media_type\n" +
			"    url: url\n" +
			"model: uni-1\n" +
			"output_format: png\n" +
			"source:\n" +
			"  data: data\n" +
			"  generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  media_type: media_type\n" +
			"  url: url\n" +
			"style: auto\n" +
			"type: image\n" +
			"user_id: user_id\n" +
			"video:\n" +
			"  duration: 5s\n" +
			"  edit:\n" +
			"    auto_controls: true\n" +
			"    controls:\n" +
			"      depth:\n" +
			"        blur: 0\n" +
			"        enabled: true\n" +
			"      face:\n" +
			"        enabled: true\n" +
			"      normals:\n" +
			"        augmentation: 0\n" +
			"        enabled: true\n" +
			"      pose:\n" +
			"        enabled: true\n" +
			"        strength: precise\n" +
			"      trajectory:\n" +
			"        enabled: true\n" +
			"        sparsity: 0\n" +
			"    keyframe_indexes:\n" +
			"      - 0\n" +
			"    keyframes:\n" +
			"      - data: data\n" +
			"        generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"        media_type: media_type\n" +
			"        url: url\n" +
			"    strength: adhere_1\n" +
			"  end_frame:\n" +
			"    data: data\n" +
			"    generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    media_type: media_type\n" +
			"    url: url\n" +
			"  exr_export: true\n" +
			"  hdr: true\n" +
			"  keyframe_indexes:\n" +
			"    - 0\n" +
			"  keyframes:\n" +
			"    - data: data\n" +
			"      generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"      media_type: media_type\n" +
			"      url: url\n" +
			"  loop: true\n" +
			"  resolution: 360p\n" +
			"  source_position:\n" +
			"    h_norm: 1\n" +
			"    w_norm: 1\n" +
			"    x_norm: -2\n" +
			"    y_norm: -2\n" +
			"  start_frame:\n" +
			"    data: data\n" +
			"    generation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    media_type: media_type\n" +
			"    url: url\n" +
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
