// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/lumalabs/luma-agents-cli/internal/apiquery"
	"github.com/lumalabs/luma-agents-cli/internal/requestflag"
	"github.com/lumalabs/luma-agents-go"
	"github.com/lumalabs/luma-agents-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var generationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Submit an image or video generation job. Returns immediately with an opaque job\nID to poll via GET /generations/{id}.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Text prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[*string]{
			Name:     "aspect-ratio",
			Usage:    "Output aspect ratio. Valid values depend on the selected model and generation type; the server validates the final model-specific set.",
			BodyPath: "aspect_ratio",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "image-ref",
			Usage:    "Reference images for style/content guidance. Up to 9 for type 'image', up to 8 for type 'image_edit'.",
			BodyPath: "image_ref",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model identifier. `uni-1` is the default image tier; `uni-1-max` produces higher-quality output than `uni-1` at a higher per-image price. `ray-3.2` is the public video model for text-to-video, image-to-video, and video-to-video editing.",
			BodyPath: "model",
		},
		&requestflag.Flag[*string]{
			Name:     "output-format",
			Usage:    "Output image format",
			BodyPath: "output_format",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "source",
			Usage:    "Media reference for guided generation. Provide exactly one of url, inline base64 data, or generation_id. URL/data references accept image media at image positions; video_edit and video_reframe sources also accept source.url or source.data when source.media_type is a video/* MIME. generation_id chains image_edit off a prior image output, video_edit/video_reframe off a prior video output, and video.start_frame/end_frame for extension.",
			BodyPath: "source",
		},
		&requestflag.Flag[string]{
			Name:     "style",
			Usage:    "Style preset (auto, manga)",
			Default:  "auto",
			BodyPath: "style",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "The kind of generation to perform",
			Default:  "image",
			BodyPath: "type",
		},
		&requestflag.Flag[*string]{
			Name:     "user-id",
			Usage:    "Your end-user's stable opaque identifier (no PII). Forwarded to upstream model providers as their per-user tagging field so trust & safety violations can be attributed to a specific end-user rather than the whole API account. Also used for per-end-user usage breakdowns in /v1/usage. Strongly recommended for partner integrations.",
			BodyPath: "user_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "video",
			Usage:    "Ray 3.2 video request options. Common output settings live at the top level for `type=video`, `type=video_edit`, and `type=video_reframe`; video-to-video conditioning lives under `edit`.",
			BodyPath: "video",
		},
		&requestflag.Flag[bool]{
			Name:     "web-search",
			Usage:    "Enable web search grounding — the agent can search the web and download reference images before generating.",
			Default:  false,
			BodyPath: "web_search",
		},
	},
	Action:          handleGenerationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"image-ref": {
		&requestflag.InnerFlag[*string]{
			Name:       "image-ref.data",
			Usage:      "Base64-encoded image or video data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "image-ref.generation-id",
			Usage:      "UUID of a prior generation owned by the same caller. Used on source for image_edit, video_edit, and video_reframe chaining and on video.start_frame / video.end_frame for video extension.",
			InnerField: "generation_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "image-ref.media-type",
			Usage:      "MIME type (for example, image/jpeg or video/mp4). Required with data. Required with source.url on video_edit/video_reframe so the route can dispatch video ingest before fetching bytes; optional for image URLs.",
			InnerField: "media_type",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "image-ref.url",
			Usage:      "Publicly accessible image URL, or a video URL when used as source for video_edit/video_reframe with media_type=video/*.",
			InnerField: "url",
		},
	},
	"source": {
		&requestflag.InnerFlag[*string]{
			Name:       "source.data",
			Usage:      "Base64-encoded image or video data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "source.generation-id",
			Usage:      "UUID of a prior generation owned by the same caller. Used on source for image_edit, video_edit, and video_reframe chaining and on video.start_frame / video.end_frame for video extension.",
			InnerField: "generation_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "source.media-type",
			Usage:      "MIME type (for example, image/jpeg or video/mp4). Required with data. Required with source.url on video_edit/video_reframe so the route can dispatch video ingest before fetching bytes; optional for image URLs.",
			InnerField: "media_type",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "source.url",
			Usage:      "Publicly accessible image URL, or a video URL when used as source for video_edit/video_reframe with media_type=video/*.",
			InnerField: "url",
		},
	},
	"video": {
		&requestflag.InnerFlag[*string]{
			Name:       "video.duration",
			Usage:      "Video duration",
			InnerField: "duration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "video.edit",
			Usage:      "Ray 3.2 video-to-video edit controls. Only valid under `video.edit` when `type` is `video_edit`. The source video must be 18 seconds or shorter; output duration matches the source.",
			InnerField: "edit",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "video.end-frame",
			Usage:      "Media reference for guided generation. Provide exactly one of url, inline base64 data, or generation_id. URL/data references accept image media at image positions; video_edit and video_reframe sources also accept source.url or source.data when source.media_type is a video/* MIME. generation_id chains image_edit off a prior image output, video_edit/video_reframe off a prior video output, and video.start_frame/end_frame for extension.",
			InnerField: "end_frame",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "video.exr-export",
			Usage:      "Export EXR alongside the MP4. Requires hdr=true.",
			InnerField: "exr_export",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "video.hdr",
			Usage:      "Generate HDR video. Requires HDR access. Not supported for video_reframe.",
			InnerField: "hdr",
		},
		&requestflag.InnerFlag[any]{
			Name:       "video.keyframe-indexes",
			Usage:      "Parallel list of non-negative, unique output-frame positions where each keyframes[i] is anchored, in the duration x 24fps grid (5s -> 0..120, 10s -> 0..240). Must match keyframes in length.",
			InnerField: "keyframe_indexes",
		},
		&requestflag.InnerFlag[any]{
			Name:       "video.keyframes",
			Usage:      "Image-to-video guide frames (type=video only), each pinned to an output-frame position via the parallel keyframe_indexes. 1-64 anchors: a single anchor is a valid start-pinned i2v (an alternate to start_frame), and any count up to 64 places guide frames at arbitrary positions. Unlike start_frame/end_frame (the legacy 2-frame surface), this supports arbitrary positions, 10s durations, and HDR. Mutually exclusive with start_frame / end_frame / loop. Only supported on model ray-3.2. For video-to-video keyframes use video.edit.keyframes on type=video_edit instead.",
			InnerField: "keyframes",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "video.loop",
			Usage:      "Generate a seamlessly looping video. Only valid for type=video; not supported with duration=10s or hdr=true.",
			InnerField: "loop",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "video.resolution",
			Usage:      "Ray 3.2 video output resolution. 360p is the draft tier (fast, low-cost previews), accepted on type=video, video_edit, and video_reframe; on type=video it is SDR-only (not valid with hdr=true). 1080p is public for video generation; video_reframe 1080p is still rolling out and may return a coming-soon validation error until enabled for the caller.",
			InnerField: "resolution",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "video.source-position",
			Usage:      "Normalized source rectangle inside the output canvas for video_reframe. Omit to let the model choose the default centered-fit crop.",
			InnerField: "source_position",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "video.start-frame",
			Usage:      "Media reference for guided generation. Provide exactly one of url, inline base64 data, or generation_id. URL/data references accept image media at image positions; video_edit and video_reframe sources also accept source.url or source.data when source.media_type is a video/* MIME. generation_id chains image_edit off a prior image output, video_edit/video_reframe off a prior video output, and video.start_frame/end_frame for extension.",
			InnerField: "start_frame",
		},
	},
})

var generationsGet = cli.Command{
	Name:    "get",
	Usage:   "Poll for generation status and output. On completion, the response includes\npresigned URLs to download the generated images.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "generation-id",
			Required:  true,
			PathParam: "generation_id",
		},
	},
	Action:          handleGenerationsGet,
	HideHelpCommand: true,
}

func handleGenerationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := lumaagents.GenerationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Generations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "generations create",
		Transform:      transform,
	})
}

func handleGenerationsGet(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("generation-id") && len(unusedArgs) > 0 {
		cmd.Set("generation-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Generations.Get(ctx, cmd.Value("generation-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "generations get",
		Transform:      transform,
	})
}
