package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidate(t *testing.T) {
	testCases := []struct {
		name         string
		genesisState GenesisState
		expErr       bool
	}{
		{
			"valid genesisState",
			GenesisState{
				Params: DefaultParams(),
			},
			false,
		},
		{"empty genesisState", GenesisState{}, true},
		{
			"invalid params ",
			GenesisState{
				Params: Params{
					MaxTxSize: 0,
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateGenesis(tc.genesisState)

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
