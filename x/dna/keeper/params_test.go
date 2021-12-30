package keeper_test

import (
	"testing"

	testkeeper "github.com/mavericks/dna/testutil/keeper"
	"github.com/mavericks/dna/x/dna/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DnaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
