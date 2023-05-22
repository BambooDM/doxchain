package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAMLRegistration(goCtx context.Context, msg *types.MsgCreateAMLRegistrationRequest) (result *types.MsgCreateAMLRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.Keeper.GetAMLRegistration(ctx, msg.Owner.GetW3CIdentifier())

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	if msg.Creator != msg.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	k.Keeper.SetAMLRegistration(
		ctx,
		types.AMLRegistration{
			Owner:    msg.Owner,
			Approved: false,
		},
	)

	return result, nil
}

func (k msgServer) DeleteAMLRegistration(goCtx context.Context, msg *types.MsgDeleteAMLRegistrationRequest) (result *types.MsgDeleteAMLRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userDid := didUtils.NewDidTokenFactory().Create(msg.Creator, "")
	match, found := k.Keeper.GetAMLRegistration(ctx, userDid.GetW3CIdentifier())

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	if msg.Creator != match.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation is not supported")
	}

	k.Keeper.RemoveAMLRegistration(ctx, userDid.GetW3CIdentifier())

	return result, nil
}
