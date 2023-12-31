package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AMLRegistration(goCtx context.Context, req *types.QueryGetAMLRegistrationRequest) (result *types.QueryGetAMLRegistrationResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.GetAMLRegistration(ctx, req.AmlRegistrationW3CIdentifier)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	result.Registration = match

	return result, nil
}
