package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDidCount fetches the Did counter from the KVStore
func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefix(types.DidCountKey))

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetDidCount updates the Did counter in the KVStore
func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.KeyPrefix(types.DidCountKey), bz)
}

// SetDid adds a Did to the KVStore and updates the Did counter
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&did)

	store.Set(GetDidIDBytes(did.GetFullyQualifiedDidIdentifier()), b)
	
	k.SetDidCount(ctx, k.GetDidCount(ctx)+1)
}

// GetDid returns a Did from its FullyQualifiedDidIdentifier (did:MethodName:MethodId)
func (k Keeper) GetDid(ctx sdk.Context, fullyQualifiedDidIdentifier string) (val types.Did, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := store.Get(GetDidIDBytes(fullyQualifiedDidIdentifier))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveDid removes a Did from the KVStore
func (k Keeper) RemoveDid(ctx sdk.Context, fullyQualifiedDidIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	store.Delete(GetDidIDBytes(fullyQualifiedDidIdentifier))
}

// GetAllDid returns all Dids in the KVStore
func (k Keeper) GetAllDid(ctx sdk.Context) (list []types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Did
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDidIDBytes returns the byte representation of the Did
func GetDidIDBytes(did string) []byte {
	return []byte(did)
}

// GetDidIDFromBytes returns ID in uint64 format from a byte array
func GetDidIDFromBytes(bz []byte) string {
	return string(bz)
}
