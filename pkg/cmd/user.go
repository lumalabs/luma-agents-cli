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

var usersCreate = cli.Command{
	Name:    "create",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			BodyPath: "username",
		},
		&requestflag.Flag[int64]{
			Name:     "user-status",
			Usage:    "User Status",
			BodyPath: "userStatus",
		},
	},
	Action:          handleUsersCreate,
	HideHelpCommand: true,
}

var usersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get user by user name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUsersRetrieve,
	HideHelpCommand: true,
}

var usersUpdate = cli.Command{
	Name:    "update",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "existing-username",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "first-name",
			BodyPath: "firstName",
		},
		&requestflag.Flag[string]{
			Name:     "last-name",
			BodyPath: "lastName",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			BodyPath: "username",
		},
		&requestflag.Flag[int64]{
			Name:     "user-status",
			Usage:    "User Status",
			BodyPath: "userStatus",
		},
	},
	Action:          handleUsersUpdate,
	HideHelpCommand: true,
}

var usersDelete = cli.Command{
	Name:    "delete",
	Usage:   "This can only be done by the logged in user.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "username",
			Required: true,
		},
	},
	Action:          handleUsersDelete,
	HideHelpCommand: true,
}

var usersCreateWithList = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-with-list",
	Usage:   "Creates list of users with given input array",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			BodyRoot: true,
		},
	},
	Action:          handleUsersCreateWithList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"item": {
		&requestflag.InnerFlag[int64]{
			Name:       "item.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.first-name",
			InnerField: "firstName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.last-name",
			InnerField: "lastName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.password",
			InnerField: "password",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.phone",
			InnerField: "phone",
		},
		&requestflag.InnerFlag[string]{
			Name:       "item.username",
			InnerField: "username",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "item.user-status",
			Usage:      "User Status",
			InnerField: "userStatus",
		},
	},
})

var usersLogin = cli.Command{
	Name:    "login",
	Usage:   "Logs user into the system",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "password",
			Usage:     "The password for login in clear text",
			QueryPath: "password",
		},
		&requestflag.Flag[string]{
			Name:      "username",
			Usage:     "The user name for login",
			QueryPath: "username",
		},
	},
	Action:          handleUsersLogin,
	HideHelpCommand: true,
}

var usersLogout = cli.Command{
	Name:            "logout",
	Usage:           "Logs out current logged in user session",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsersLogout,
	HideHelpCommand: true,
}

func handleUsersCreate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.UserNewParams{}

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
	_, err = client.Users.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users create", obj, format, transform)
}

func handleUsersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
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
	_, err = client.Users.Get(ctx, cmd.Value("username").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users retrieve", obj, format, transform)
}

func handleUsersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("existing-username") && len(unusedArgs) > 0 {
		cmd.Set("existing-username", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.UserUpdateParams{}

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

	return client.Users.Update(
		ctx,
		cmd.Value("existing-username").(string),
		params,
		options...,
	)
}

func handleUsersDelete(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("username") && len(unusedArgs) > 0 {
		cmd.Set("username", unusedArgs[0])
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

	return client.Users.Delete(ctx, cmd.Value("username").(string), options...)
}

func handleUsersCreateWithList(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.UserNewWithListParams{}

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
	_, err = client.Users.NewWithList(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users create-with-list", obj, format, transform)
}

func handleUsersLogin(ctx context.Context, cmd *cli.Command) error {
	client := lumaagents.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lumaagents.UserLoginParams{}

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
	_, err = client.Users.Login(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users login", obj, format, transform)
}

func handleUsersLogout(ctx context.Context, cmd *cli.Command) error {
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

	return client.Users.Logout(ctx, options...)
}
