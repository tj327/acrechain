package types

import (
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgMintTokensToAccount{}
)

const (
	TypeMsgMintTokensToAccount = "mint_tokens_to_account"
)

// NewMsgMintTokensToAccount creates new instance of MsgMintTokensToAccount
func NewMsgMintTokensToAccount(
	fromAddr, toAddr sdk.AccAddress,
	amounts github_com_cosmos_cosmos_sdk_types.Coins,
) *MsgMintTokensToAccount {
	return &MsgMintTokensToAccount{
		Sender:    fromAddr.String(),
		ToAddress: toAddr.String(),
		Amounts:   amounts,
	}
}

// Route returns the name of the module
func (msg MsgMintTokensToAccount) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgMintTokensToAccount) Type() string { return TypeMsgMintTokensToAccount }

// GetSignBytes encodes the message for signing
func (msg *MsgMintTokensToAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// ValidateBasic runs stateless checks on the message
func (msg MsgMintTokensToAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(err, "invalid from_account address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.ToAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid to_account address")
	}

	return nil
}

// GetSigners defines whose signature is required
func (msg MsgMintTokensToAccount) GetSigners() []sdk.AccAddress {
	from := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{from}
}
