package keeper_test

import (
	"testing"

	testkeeper "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AmlKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
