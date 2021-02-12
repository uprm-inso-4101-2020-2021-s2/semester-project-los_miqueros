package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group Themis queries under a subcommand
	ThemisQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ThemisQueryCmd.AddCommand(
		flags.GetCommands(
			// this line is used by starport scaffolding # 1
			GetCmdListPoll(queryRoute, cdc),
			GetCmdListGroupPolls(queryRoute, cdc),
			GetCmdGetPoll(queryRoute, cdc),
			GetCmdListAccount(queryRoute, cdc),
			GetCmdListUserAccounts(queryRoute, cdc),
			GetCmdListGroupAccounts(queryRoute, cdc),
			GetCmdGetAccount(queryRoute, cdc),
			GetCmdListGroup(queryRoute, cdc),
			GetCmdGetGroup(queryRoute, cdc),
		)...,
	)

	return ThemisQueryCmd
}
