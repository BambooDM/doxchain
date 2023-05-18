package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccessTokenRegistries:       []AccessTokenRegistry{},
		AuthorizationCodeRegistries: []AuthorizationCodeRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in AccessTokenRegistry
	AccessTokenRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccessTokenRegistries {
		tenant := string(AccessTokenRegistryKey(elem.Owner.Creator))
		if _, ok := AccessTokenRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated tenant for AccessTokenRegistry")
		}
		AccessTokenRegistryIndexMap[tenant] = struct{}{}
	}
	// Check for duplicated index in authorizationCodeRegistry
	authorizationCodeRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.AuthorizationCodeRegistries {
		tenant := string(AuthorizationCodeRegistryKey(elem.Owner.Creator))
		if _, ok := authorizationCodeRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated tenant for authorizationCodeRegistry")
		}
		authorizationCodeRegistryIndexMap[tenant] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
