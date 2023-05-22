package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetKYCRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetKYCRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendKYCRegistration(ctx sdk.Context, kycRequest types.KYCRegistration) string {
	count := k.GetKYCRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	w3cIdentifier := kycRequest.Owner.GetW3CIdentifier()

	store.Set(GetKYCRegistrationIDBytes(w3cIdentifier), k.cdc.MustMarshal(&kycRequest))

	k.SetKYCRegistrationCount(ctx, count+1)

	return w3cIdentifier
}

func (k Keeper) SetKYCRegistration(ctx sdk.Context, request types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := k.cdc.MustMarshal(&request)

	store.Set(GetKYCRegistrationIDBytes(request.Owner.GetW3CIdentifier()), b)
}

func (k Keeper) GetKYCRegistration(ctx sdk.Context, registrationW3CIdentifier string) (result types.KYCRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := store.Get(GetKYCRegistrationIDBytes(registrationW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveKYCRegistration(ctx sdk.Context, registrationW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))

	store.Delete(GetKYCRegistrationIDBytes(registrationW3CIdentifier))
}

func (k Keeper) GetAllKYCRegistration(ctx sdk.Context) (list []types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KYCRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ApproveKYCRegistration(ctx sdk.Context, registrationW3CIdentifier string) {
	request, found := k.GetKYCRegistration(ctx, registrationW3CIdentifier)

	if found {
		request.Approved = true

		k.SetKYCRegistration(ctx, request)
	}
}

func GetKYCRegistrationIDBytes(did string) []byte {
	return []byte(did)
}

func GetKYCRegistrationIDFromBytes(bz []byte) string {
	return string(bz)
}
