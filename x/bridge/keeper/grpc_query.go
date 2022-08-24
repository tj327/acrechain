package keeper

import (
	"context"

	"github.com/ArableProtocol/acrechain/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) AccountAllocation(
	goCtx context.Context,
	req *types.QueryAccountAllocationRequest,
) (*types.QueryAccountAllocationResponse, error) {
	return nil, status.Error(codes.InvalidArgument, "empty request")

}

func (k Keeper) AllAllocations(
	goCtx context.Context,
	req *types.QueryAllAllocationsRequest,
) (*types.QueryAllAllocationsResponse, error) {
	return nil, status.Error(codes.InvalidArgument, "empty request")

}

func (k Keeper) Params(
	goCtx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	return nil, status.Error(codes.InvalidArgument, "empty request")

}
