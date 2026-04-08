// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/luma-agents-cli/internal/apiquery"
	"github.com/stainless-sdks/luma-agents-cli/internal/requestflag"
	"github.com/stainless-sdks/luma-agents-go"
	"github.com/stainless-sdks/luma-agents-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var storeOrdersCreate = cli.Command{
	Name:    "create",
	Usage:   "Place a new order in the store",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[bool]{
			Name:     "complete",
			BodyPath: "complete",
		},
		&requestflag.Flag[int64]{
			Name:     "pet-id",
			BodyPath: "petId",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			BodyPath: "quantity",
		},
		&requestflag.Flag[any]{
			Name:     "ship-date",
			BodyPath: "shipDate",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Order Status",
			BodyPath: "status",
		},
	},
	Action:          handleStoreOrdersCreate,
	HideHelpCommand: true,
}

var storeOrdersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "For valid response try integer IDs with value <= 5 or > 10. Other values will\ngenerate exceptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleStoreOrdersRetrieve,
	HideHelpCommand: true,
}

var storeOrdersDelete = cli.Command{
	Name:    "delete",
	Usage:   "For valid response try integer IDs with value < 1000. Anything above 1000 or\nnonintegers will generate API errors",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "order-id",
			Required: true,
		},
	},
	Action:          handleStoreOrdersDelete,
	HideHelpCommand: true,
}

func handleStoreOrdersCreate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.StoreOrderNewParams{}

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
	_, err = client.Store.Orders.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "store:orders create", obj, format, transform)
}

func handleStoreOrdersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
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
	_, err = client.Store.Orders.Get(ctx, cmd.Value("order-id").(int64), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "store:orders retrieve", obj, format, transform)
}

func handleStoreOrdersDelete(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("order-id") && len(unusedArgs) > 0 {
		cmd.Set("order-id", unusedArgs[0])
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

	return client.Store.Orders.Delete(ctx, cmd.Value("order-id").(int64), options...)
}
