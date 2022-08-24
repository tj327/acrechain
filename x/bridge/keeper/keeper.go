package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ArableProtocol/acrechain/x/bridge/types"
)

// Keeper of this module maintains collections of vesting.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	mintKeeper types.MintKeeper
	bankKeeper types.BankKeeper
}

// NewKeeper creates new instances of the vesting Keeper
func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	mk types.MintKeeper,
	bk types.BankKeeper,
) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		mintKeeper: mk,
		bankKeeper: bk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
