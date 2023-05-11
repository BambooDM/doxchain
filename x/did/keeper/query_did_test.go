package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/did/types"
)

func TestDidQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDid(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDidRequest
		response *types.QueryGetDidResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetDidRequest{FullyQualifiedDidIdentifier: string(msgs[0].Id)},
			response: &types.QueryGetDidResponse{Did: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetDidRequest{FullyQualifiedDidIdentifier: string(msgs[1].Id)},
			response: &types.QueryGetDidResponse{Did: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetDidRequest{FullyQualifiedDidIdentifier: string(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Did(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDidQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDid(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDidRequest {
		return &types.QueryAllDidRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DidAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Did), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Did),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DidAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Did), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Did),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DidAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Did),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DidAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
