package paper

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"paper/testutil/sample"
	papersimulation "paper/x/paper/simulation"
	"paper/x/paper/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = papersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePaper = "op_weight_msg_create_paper"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePaper int = 100

	opWeightMsgUpdatePaper = "op_weight_msg_update_paper"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePaper int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	paperGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&paperGenesis)
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

	var weightMsgCreatePaper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePaper, &weightMsgCreatePaper, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePaper = defaultWeightMsgCreatePaper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePaper,
		papersimulation.SimulateMsgCreatePaper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePaper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePaper, &weightMsgUpdatePaper, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePaper = defaultWeightMsgUpdatePaper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePaper,
		papersimulation.SimulateMsgUpdatePaper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
