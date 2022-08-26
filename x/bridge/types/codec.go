package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// ModuleCdc references the global erc20 module codec. Note, the codec should
// ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to modules/erc20 and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	// NOTE: BaseVestingAccount is still supported to as it's the underlying embedded

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}