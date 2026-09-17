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

var filesCreate = cli.Command{
	Name:    "create",
	Usage:   "Upload a file to your namespace, then reference it from a generation via\nImageRef.file_id (as source, image_ref[], video.start_frame, keyframes, and so\non). Two upload modes share this endpoint, selected by Content-Type:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "mime-type",
			Usage:    "MIME type of the bytes you will upload.",
			Required: true,
			BodyPath: "mime_type",
		},
		&requestflag.Flag[int64]{
			Name:     "size-bytes",
			Usage:    "Exact size in bytes of the object you will PUT. Up to 5 GiB (the S3 single-PUT ceiling).",
			Required: true,
			BodyPath: "size_bytes",
		},
		&requestflag.Flag[any]{
			Name:     "expires-at",
			Usage:    "Optional TTL. After this time Luma may automatically delete the file and reclaim its bytes.",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[*string]{
			Name:     "filename",
			Usage:    "Optional original filename to record.",
			BodyPath: "filename",
		},
		&requestflag.Flag[string]{
			Name:     "purpose",
			Usage:    "How the file is intended to be used in a generation. `input` is the primary subject (e.g. the source image for an edit); `reference` is style/content guidance.",
			BodyPath: "purpose",
		},
		&requestflag.Flag[*string]{
			Name:     "user-id",
			Usage:    "Optional opaque end-user tag for abuse attribution. Mirrors the user_id field on POST /generations.",
			BodyPath: "user_id",
		},
	},
	Action:          handleFilesCreate,
	HideHelpCommand: true,
}

var filesList = cli.Command{
	Name:    "list",
	Usage:   "List the files in your namespace, newest first. Keyset-paginated: when has_more\nis true, pass next_cursor back as cursor.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a prior response's next_cursor.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum files to return (1–100). Defaults to 25.",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "purpose",
			Usage:     "How the file is intended to be used in a generation. `input` is the primary subject (e.g. the source image for an edit); `reference` is style/content guidance.",
			QueryPath: "purpose",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Lifecycle state of an uploaded file. `pending` until bytes are received and the ingest pipeline runs; `ready` once it can be referenced from a generation; `failed` if ingest/moderation rejected it; `deleted` after a soft-delete.",
			QueryPath: "state",
		},
	},
	Action:          handleFilesList,
	HideHelpCommand: true,
}

var filesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Soft-delete a file. It can no longer be referenced from new generations. Returns\n204 with no body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
	},
	Action:          handleFilesDelete,
	HideHelpCommand: true,
}

var filesComplete = cli.Command{
	Name:    "complete",
	Usage:   "Finalize a presigned upload after you have PUT the bytes to the upload URL.\nKicks off ingest/moderation and returns the file, which transitions to `ready`\n(or `failed`) asynchronously — poll GET /files/{file_id} to observe the terminal\nstate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
	},
	Action:          handleFilesComplete,
	HideHelpCommand: true,
}

var filesGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve metadata for a single file in your namespace.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
	},
	Action:          handleFilesGet,
	HideHelpCommand: true,
}

func handleFilesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := lumaagents.FileNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.New(ctx, params, options...)
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
		Title:          "files create",
		Transform:      transform,
	})
}

func handleFilesList(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := lumaagents.FileListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.List(ctx, params, options...)
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
		Title:          "files list",
		Transform:      transform,
	})
}

func handleFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	return client.Files.Delete(ctx, cmd.Value("file-id").(string), options...)
}

func handleFilesComplete(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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
	_, err = client.Files.Complete(ctx, cmd.Value("file-id").(string), options...)
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
		Title:          "files complete",
		Transform:      transform,
	})
}

func handleFilesGet(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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
	_, err = client.Files.Get(ctx, cmd.Value("file-id").(string), options...)
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
		Title:          "files get",
		Transform:      transform,
	})
}
