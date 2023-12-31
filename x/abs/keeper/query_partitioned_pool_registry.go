package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PartitionedPoolRegistryAll(goCtx context.Context, req *types.QueryAllPartitionedPoolRegistriesRequest) (result *types.QueryAllPartitionedPoolRegistriesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var partitionedPoolRegistryList []types.PartitionedPoolRegistry

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	partitionedPoolRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))

	pageRes, err := query.Paginate(partitionedPoolRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var partitionedPoolRegistry types.PartitionedPoolRegistry

		if err := k.cdc.Unmarshal(value, &partitionedPoolRegistry); err != nil {
			return err
		}

		partitionedPoolRegistryList = append(partitionedPoolRegistryList, partitionedPoolRegistry)
		
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result.PartitionedPoolRegistryList = partitionedPoolRegistryList
	result.Pagination = pageRes

	return result, nil
}

func (k Keeper) PartitionedPoolRegistry(goCtx context.Context, req *types.QueryGetPartitionedPoolRegistryRequest) (result *types.QueryGetPartitionedPoolRegistryResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	match, found := k.GetPartitionedPoolRegistry(
		ctx,
		req.PartitionedPoolRegistryW3CIdentifier,
	)

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	result.PartitionedPoolRegistry = match

	return result, nil
}
