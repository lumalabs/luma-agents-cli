// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/luma-agents-cli/internal/apiquery"
	"github.com/stainless-sdks/luma-agents-cli/internal/binaryparam"
	"github.com/stainless-sdks/luma-agents-cli/internal/requestflag"
	"github.com/stainless-sdks/luma-agents-go"
	"github.com/stainless-sdks/luma-agents-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var petsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add a new pet to the store",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "photo-url",
			Required: true,
			BodyPath: "photoUrls",
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "pet status in the store",
			BodyPath: "status",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handlePetsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[int64]{
			Name:       "category.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "category.name",
			InnerField: "name",
		},
	},
	"tag": {
		&requestflag.InnerFlag[int64]{
			Name:       "tag.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tag.name",
			InnerField: "name",
		},
	},
})

var petsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a single pet",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
	},
	Action:          handlePetsRetrieve,
	HideHelpCommand: true,
}

var petsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing pet by Id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "photo-url",
			Required: true,
			BodyPath: "photoUrls",
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "category",
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "pet status in the store",
			BodyPath: "status",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tag",
			BodyPath: "tags",
		},
	},
	Action:          handlePetsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"category": {
		&requestflag.InnerFlag[int64]{
			Name:       "category.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "category.name",
			InnerField: "name",
		},
	},
	"tag": {
		&requestflag.InnerFlag[int64]{
			Name:       "tag.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "tag.name",
			InnerField: "name",
		},
	},
})

var petsDelete = cli.Command{
	Name:    "delete",
	Usage:   "delete a pet",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
	},
	Action:          handlePetsDelete,
	HideHelpCommand: true,
}

var petsFindByStatus = cli.Command{
	Name:    "find-by-status",
	Usage:   "Multiple status values can be provided with comma separated strings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Status values that need to be considered for filter",
			Default:   "available",
			QueryPath: "status",
		},
	},
	Action:          handlePetsFindByStatus,
	HideHelpCommand: true,
}

var petsFindByTags = cli.Command{
	Name:    "find-by-tags",
	Usage:   "Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3\nfor testing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Tags to filter by",
			QueryPath: "tags",
		},
	},
	Action:          handlePetsFindByTags,
	HideHelpCommand: true,
}

var petsUpdateByID = cli.Command{
	Name:    "update-by-id",
	Usage:   "Updates a pet in the store with form data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Name of pet that needs to be updated",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Status of pet that needs to be updated",
			QueryPath: "status",
		},
	},
	Action:          handlePetsUpdateByID,
	HideHelpCommand: true,
}

var petsUploadImage = cli.Command{
	Name:    "upload-image",
	Usage:   "uploads an image",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "image",
			Required:  true,
			BodyRoot:  true,
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:      "additional-metadata",
			Usage:     "Additional Metadata",
			QueryPath: "additionalMetadata",
		},
	},
	Action:          handlePetsUploadImage,
	HideHelpCommand: true,
}

func handlePetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.PetNewParams{}

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
	_, err = client.Pets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets create", obj, format, transform)
}

func handlePetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
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
	_, err = client.Pets.Get(ctx, cmd.Value("pet-id").(int64), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets retrieve", obj, format, transform)
}

func handlePetsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.PetUpdateParams{}

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
	_, err = client.Pets.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets update", obj, format, transform)
}

func handlePetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
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

	return client.Pets.Delete(ctx, cmd.Value("pet-id").(int64), options...)
}

func handlePetsFindByStatus(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.PetFindByStatusParams{}

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
	_, err = client.Pets.FindByStatus(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets find-by-status", obj, format, transform)
}

func handlePetsFindByTags(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.PetFindByTagsParams{}

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
	_, err = client.Pets.FindByTags(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets find-by-tags", obj, format, transform)
}

func handlePetsUpdateByID(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.PetUpdateByIDParams{}

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

	return client.Pets.UpdateByID(
		ctx,
		cmd.Value("pet-id").(int64),
		params,
		options...,
	)
}

func handlePetsUploadImage(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("image") && len(unusedArgs) > 0 {
		cmd.Set("image", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	bodyReader, stdinInUse, err := binaryparam.FileOrStdin(os.Stdin, cmd.Value("image").(string))
	if err != nil {
		return fmt.Errorf("Failed on param '%s': %w", "body", err)
	}
	defer bodyReader.Close()

	params := lumaagents.PetUploadImageParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationOctetStream,
		stdinInUse,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pets.UploadImage(
		ctx,
		cmd.Value("pet-id").(int64),
		bodyReader,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "pets upload-image", obj, format, transform)
}
