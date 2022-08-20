package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "paper/testutil/keeper"
	"paper/x/paper/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PaperKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
