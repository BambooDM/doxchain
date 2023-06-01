package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDidRequest) (result *types.MsgCreateDidResponse, err error) {
	if msg.Creator != msg.Did.Creator || msg.Did.IsModuleIdentifier() {
		return nil, types.ErrImpersonation
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	w3cIdentifier := msg.Did.GetW3CIdentifier()
	_, found := k.Keeper.GetDid(ctx, w3cIdentifier)

	if found {
		return nil, types.ErrDidExists
	}

	k.Keeper.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, false)

	result = &types.MsgCreateDidResponse{
		DidW3CIdentifier: w3cIdentifier,
	}

	return result, nil
}

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDidRequest) (result *types.MsgUpdateDidResponse, err error) {
	if msg.Creator != msg.Did.Creator || msg.Did.IsModuleIdentifier() {
		return nil, types.ErrImpersonation
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	w3cIdentifier := msg.Did.GetW3CIdentifier()
	match, found := k.Keeper.GetDid(ctx, w3cIdentifier)

	if found && msg.Creator != match.Creator {
		return nil, types.ErrImpersonation
	}

	k.Keeper.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, true)

	result.DidW3CIdentifier = w3cIdentifier

	return result, nil
}

func (k msgServer) DeleteDid(goCtx context.Context, msg *types.MsgDeleteDidRequest) (result *types.MsgDeleteDidResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDid(ctx, msg.DidW3CIdentifier)

	if !found {
		return nil, types.ErrDidNotFound
	}

	if msg.Creator != match.Creator {
		return nil, types.ErrImpersonation
	}

	k.Keeper.RemoveDid(ctx, msg.DidW3CIdentifier)

	return result, nil
}
