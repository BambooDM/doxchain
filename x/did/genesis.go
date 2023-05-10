package did

import (
	"doxchain/x/did/keeper"
	"doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the did
	for _, elem := range genState.DidList {
		k.SetDid(ctx, elem)
	}

	// Set did count
	k.SetDidCount(ctx, genState.DidCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DidList = k.GetAllDid(ctx)
	genesis.DidCount = k.GetDidCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
