package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	"gopkg.in/yaml.v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter store key
var (
	DefaultOperators = []didTypes.Did(nil) // none allowed
	DefaultBlockExpireOffset = sdk.NewInt(100000)

	ParamStoreKeyOperators = []byte("Operators")
	ParamStoreKeyBlockExpireOffset = []byte("BlockExpireOffset")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(operators []didTypes.Did, blockExpireOffset sdk.Int) Params {
	return Params{
		Operators: operators,
		BlockExpireOffset: blockExpireOffset,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultOperators, DefaultBlockExpireOffset)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{		
		paramtypes.NewParamSetPair(ParamStoreKeyOperators, &p.Operators, validateOperators),
		paramtypes.NewParamSetPair(ParamStoreKeyBlockExpireOffset, &p.BlockExpireOffset, validateBlockExpireOffset),
	}
}

func validateBlockExpireOffset(i interface{}) error {
	_, ok := i.(sdk.Int)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateOperators(i interface{}) error {
	_, ok := i.([]didTypes.Did)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBlockExpireOffset(p.BlockExpireOffset); err != nil {
		return err
	}

	if err := validateOperators(p.Operators); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
