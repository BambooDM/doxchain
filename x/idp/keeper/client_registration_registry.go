package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetClientRegistrationRegistry(ctx sdk.Context, clientRegistrationRegistry types.ClientRegistrationRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&clientRegistrationRegistry)

	store.Set(types.ClientRegistrationRegistryKey(
		clientRegistrationRegistry.Owner.GetW3CIdentifier(),
	), b)
}

func (k Keeper) GetClientRegistrationRegistry(ctx sdk.Context, clientRegistryW3CIdentifier string) (val types.ClientRegistrationRegistry, found bool) {
	//TODO: Validate did string. Should be => did:sdk.AccAddress:sdk.AccAddress
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	b := store.Get(types.ClientRegistrationRegistryKey(
		clientRegistryW3CIdentifier,
	))
	
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveClientRegistrationRegistry(ctx sdk.Context, clientRegistryW3CIdentifier string) {
	//TODO: Validate did string. Should be => did:sdk.AccAddress:sdk.AccAddress
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	
	store.Delete(types.ClientRegistrationRegistryKey(
		clientRegistryW3CIdentifier,
	))
}

func (k Keeper) GetAllClientRegistrationRegistry(ctx sdk.Context) (list []types.ClientRegistrationRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientRegistrationRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
