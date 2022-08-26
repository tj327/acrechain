package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ArableProtocol/acrechain/x/bridge/types"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) MintTokensToAccount(
	goCtx context.Context,
	msg *types.MsgMintTokensToAccount,
) (*types.MsgMintTokensToAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	mk := k.mintKeeper
	bk := k.bankKeeper

	// from := sdk.MustAccAddressFromBech32(msg.Sender)
	to := sdk.MustAccAddressFromBech32(msg.ToAddress)

	coinAmounts := msg.Amounts

	if coinAmounts.Empty() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins,
			"invalid coins, can not mint to ", msg.ToAddress,
		)
	}

	mintError := mk.MintCoins(ctx, msg.Amounts)
	if mintError != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins,
			"failed to mint to coins to ", types.ModuleName,
		)
	}

	sendError := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, msg.Amounts)
	if sendError != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrPanic,
			"failed to mint to coins to ", msg.ToAddress,
		)
	}

	store := ctx.KVStore(k.storeKey)

	res := store.Get(to.Bytes())
	if res == nil {
		store.Set(to.Bytes())
	}

	store.Set(to.Bytes())
	return &types.MsgMintTokensToAccountResponse{}, nil
}
