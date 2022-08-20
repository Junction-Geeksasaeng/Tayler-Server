package paper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "paper/testutil/keeper"
	"paper/testutil/nullify"
	"paper/x/paper"
	"paper/x/paper/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PaperKeeper(t)
	paper.InitGenesis(ctx, *k, genesisState)
	got := paper.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
