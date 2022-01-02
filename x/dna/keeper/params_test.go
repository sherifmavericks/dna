package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/sherifmavericks/dna/testutil/keeper"
	"github.com/sherifmavericks/dna/x/dna/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DnaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
