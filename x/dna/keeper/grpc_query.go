package keeper

import (
	"github.com/mavericks/dna/x/dna/types"
)

var _ types.QueryServer = Keeper{}
