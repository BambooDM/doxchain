package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/be-heroes/doxchain/x/idp/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		oauthTwoKeeper types.OAuthTwoKeeper
		authzKeeper    types.AuthzKeeper
		evidenceKeeper types.EvidenceKeeper
	}
)

//TODO: Finish ClientRegistration concept
//TODO: Implement ClientRegistrationGraph concept to infer audience relationships
//TODO: Implement IDP metadata logic (.well-known, etc)
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	oauthTwoKeeper types.OAuthTwoKeeper,
	authzKeeper types.AuthzKeeper,
	evidenceKeeper types.EvidenceKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		oauthTwoKeeper: oauthTwoKeeper,
		authzKeeper:    authzKeeper,
		evidenceKeeper: evidenceKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
