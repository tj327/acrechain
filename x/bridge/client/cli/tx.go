package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ArableProtocol/acrechain/x/bridge/types"
)

const (
	FlagFrom = "from"
)

// NewTxCmd returns a root CLI command handler for certain modules/bridge
// transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bridge transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewMsgMintTokensToAccountCmd(),
	)

	return txCmd
}

// NewMsgMintTokensToAccount transaction.
func NewMsgMintTokensToAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [to_address] [coin] [,[coin]]",
		Short: "Mint tokens.",
		Long:  `Mint tokens to account.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			msg := types.NewMsgMintTokensToAccount(clientCtx.GetFromAddress(), toAddr, coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
