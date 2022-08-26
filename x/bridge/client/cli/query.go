package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/ArableProtocol/acrechain/x/bridge/types"
)

// GetQueryCmd returns the parent command for all vesting CLI query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the bridge module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetAccountAllocationsCmd(),
	)
	return cmd
}

// GetAccountAllocationsCmd queries the minted tokens for a given account
func GetAccountAllocationsCmd() *cobra.Command {
	// cmd := &cobra.Command{
	// 	Use:   "balances [address]",
	// 	Short: "Gets locked, unvested and vested tokens for a vesting account",
	// 	Long:  "Gets locked, unvested and vested tokens for a vesting account",
	// 	Args:  cobra.ExactArgs(1),
	// 	RunE: func(cmd *cobra.Command, args []string) error {
	// 		clientCtx, err := client.GetClientQueryContext(cmd)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		queryClient := types.NewQueryClient(clientCtx)

	// 		req := &types.QueryBalancesRequest{
	// 			Address: args[0],
	// 		}

	// 		res, err := queryClient.Balances(context.Background(), req)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		return clientCtx.PrintString(
	// 			fmt.Sprintf("Locked: %s\nUnvested: %s\nVested: %s\n", res.Locked, res.Unvested, res.Vested))
	// 	},
	// }

	// flags.AddQueryFlagsToCmd(cmd)
	// return cmd
}
