package dna_test

import (
	"testing"

	keepertest "github.com/sherifmavericks/dna/testutil/keeper"
	"github.com/sherifmavericks/dna/testutil/nullify"
	"github.com/sherifmavericks/dna/x/dna"
	"github.com/sherifmavericks/dna/x/dna/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DnaKeeper(t)
	dna.InitGenesis(ctx, *k, genesisState)
	got := dna.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
