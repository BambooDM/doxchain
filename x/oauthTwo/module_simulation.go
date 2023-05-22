package oauthtwo

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	oauthtwosimulation "github.com/be-heroes/doxchain/x/oauthtwo/simulation"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oauthtwosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))

	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	oauthGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AccessTokenRegistries: []types.AccessTokenRegistry{
			{
				Owner: *didUtils.NewDidTokenFactory().Create("1", "did:methodname:methodid"),
			},
			{
				Owner: *didUtils.NewDidTokenFactory().Create("2", "did:methodname:methodid"),
			},
		},
		AuthorizationCodeRegistries: []types.AuthorizationCodeRegistry{
			{
				Owner: *didUtils.NewDidTokenFactory().Create("1", "did:methodname:methodid"),
			},
			{
				Owner: *didUtils.NewDidTokenFactory().Create("2", "did:methodname:methodid"),
			},
		},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oauthGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	return operations
}
