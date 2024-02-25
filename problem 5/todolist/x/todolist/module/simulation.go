package todolist

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"todolist/testutil/sample"
	todolistsimulation "todolist/x/todolist/simulation"
	"todolist/x/todolist/types"
)

// avoid unused import issue
var (
	_ = todolistsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateItem = "op_weight_msg_create_item"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateItem int = 100

	opWeightMsgDeleteItem = "op_weight_msg_delete_item"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteItem int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	todolistGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&todolistGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateItem int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateItem, &weightMsgCreateItem, nil,
		func(_ *rand.Rand) {
			weightMsgCreateItem = defaultWeightMsgCreateItem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateItem,
		todolistsimulation.SimulateMsgCreateItem(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteItem int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteItem, &weightMsgDeleteItem, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteItem = defaultWeightMsgDeleteItem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteItem,
		todolistsimulation.SimulateMsgDeleteItem(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateItem,
			defaultWeightMsgCreateItem,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				todolistsimulation.SimulateMsgCreateItem(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteItem,
			defaultWeightMsgDeleteItem,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				todolistsimulation.SimulateMsgDeleteItem(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
