// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/luma-agents-cli/internal/apiquery"
	"github.com/stainless-sdks/luma-agents-cli/internal/requestflag"
	"github.com/stainless-sdks/luma-agents-go"
	"github.com/stainless-sdks/luma-agents-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var generationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Submit an image generation or edit job. Returns immediately with an opaque job\nID to poll via GET /generations/{id}.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Text prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "aspect-ratio",
			Usage:    "Output aspect ratio",
			BodyPath: "aspect_ratio",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "image-ref",
			Usage:    "Reference images for style/content guidance. Up to 8 reference images.",
			BodyPath: "image_ref",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model to use",
			Default:  "uni-1",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "output-format",
			Usage:    "Output image format",
			BodyPath: "output_format",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "source",
			Usage:    "Reference image for guided generation. Provide either url or inline base64 data (not both).",
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
		&requestflag.InnerFlag[any]{
			Name:       "image-ref.data",
			Usage:      "Base64-encoded image data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[any]{
			Name:       "image-ref.media-type",
			Usage:      "MIME type (e.g. image/jpeg). Required with data.",
			InnerField: "media_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "image-ref.url",
			Usage:      "Publicly accessible image URL",
			InnerField: "url",
		},
	},
	"source": {
		&requestflag.InnerFlag[any]{
			Name:       "source.data",
			Usage:      "Base64-encoded image data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[any]{
			Name:       "source.media-type",
			Usage:      "MIME type (e.g. image/jpeg). Required with data.",
			InnerField: "media_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "source.url",
			Usage:      "Publicly accessible image URL",
			InnerField: "url",
		},
	},
})

var generationsGet = cli.Command{
	Name:    "get",
	Usage:   "Poll for generation status and output. On completion, the response includes\npresigned URLs to download the generated images.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "generation-id",
			Required: true,
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

	params := lumaagents.GenerationNewParams{}

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
