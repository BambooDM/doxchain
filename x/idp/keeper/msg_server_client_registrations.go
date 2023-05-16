package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateClientRegistry(goCtx context.Context, msg *types.MsgCreateClientRegistry) (*types.MsgCreateClientRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetClientRegistry(
		ctx,
		msg.Creator,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var ClientRegistry = types.ClientRegistry{
		Creator: msg.Creator,
	}

	k.SetClientRegistry(
		ctx,
		ClientRegistry,
	)
	return &types.MsgCreateClientRegistryResponse{}, nil
}

func (k msgServer) UpdateClientRegistry(goCtx context.Context, msg *types.MsgUpdateClientRegistry) (*types.MsgUpdateClientRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClientRegistry(
		ctx,
		msg.Creator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "creator not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var ClientRegistry = types.ClientRegistry{
		Creator: msg.Creator,
	}

	k.SetClientRegistry(ctx, ClientRegistry)

	return &types.MsgUpdateClientRegistryResponse{}, nil
}

func (k msgServer) DeleteClientRegistry(goCtx context.Context, msg *types.MsgDeleteClientRegistry) (*types.MsgDeleteClientRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClientRegistry(
		ctx,
		msg.Creator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveClientRegistry(
		ctx,
		msg.Creator,
	)

	return &types.MsgDeleteClientRegistryResponse{}, nil
}