package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const PrecisionForCoinsPerCU uint64 = 1000000

var (
	KeyMinStake            = []byte("MinStake")
	DefaultMinStake uint64 = 1000
)

var (
	KeyCoinsPerCU = []byte("CoinsPerCU")
	//this is divided by 1000000 later to get the coins per CU factor
	DefaultCoinsPerCU uint64 = 10000
)

var (
	KeyUnstakeHoldBlocks            = []byte("UnstakeHoldBlocks")
	DefaultUnstakeHoldBlocks uint64 = 10
)

var (
	KeyFraudStakeSlashingFactor = []byte("FraudStakeSlashingFactor")
	//this is divided by 1000000 later to get slashing factor
	DefaultFraudStakeSlashingFactor uint64 = 500000
)

var (
	KeyFraudSlashingAmount            = []byte("FraudSlashingAmount")
	DefaultFraudSlashingAmount uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minStake uint64,
	coinsPerCU uint64,
	unstakeHoldBlocks uint64,
	fraudStakeSlashingFactor uint64,
	fraudSlashingAmount uint64,
) Params {
	return Params{
		MinStake:                 minStake,
		CoinsPerCU:               coinsPerCU,
		UnstakeHoldBlocks:        unstakeHoldBlocks,
		FraudStakeSlashingFactor: fraudStakeSlashingFactor,
		FraudSlashingAmount:      fraudSlashingAmount,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinStake,
		DefaultCoinsPerCU,
		DefaultUnstakeHoldBlocks,
		DefaultFraudStakeSlashingFactor,
		DefaultFraudSlashingAmount,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinStake, &p.MinStake, validateMinStake),
		paramtypes.NewParamSetPair(KeyCoinsPerCU, &p.CoinsPerCU, validateCoinsPerCU),
		paramtypes.NewParamSetPair(KeyUnstakeHoldBlocks, &p.UnstakeHoldBlocks, validateUnstakeHoldBlocks),
		paramtypes.NewParamSetPair(KeyFraudStakeSlashingFactor, &p.FraudStakeSlashingFactor, validateFraudStakeSlashingFactor),
		paramtypes.NewParamSetPair(KeyFraudSlashingAmount, &p.FraudSlashingAmount, validateFraudSlashingAmount),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinStake(p.MinStake); err != nil {
		return err
	}

	if err := validateCoinsPerCU(p.CoinsPerCU); err != nil {
		return err
	}

	if err := validateUnstakeHoldBlocks(p.UnstakeHoldBlocks); err != nil {
		return err
	}

	if err := validateFraudStakeSlashingFactor(p.FraudStakeSlashingFactor); err != nil {
		return err
	}

	if err := validateFraudSlashingAmount(p.FraudSlashingAmount); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMinStake validates the MinStake param
func validateMinStake(v interface{}) error {
	minStake, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if minStake == 0 {
		return fmt.Errorf("invalid minStake value: %d, must be positive", minStake)
	}
	_ = minStake

	return nil
}

// validateCoinsPerCU validates the CoinsPerCU param
func validateCoinsPerCU(v interface{}) error {
	coinsPerCU, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	_ = coinsPerCU

	return nil
}

// validateUnstakeHoldBlocks validates the UnstakeHoldBlocks param
func validateUnstakeHoldBlocks(v interface{}) error {
	unstakeHoldBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if unstakeHoldBlocks > 15000000 {
		return fmt.Errorf("invalid unstakeHoldBlocks value: %d", unstakeHoldBlocks)
	}
	_ = unstakeHoldBlocks

	return nil
}

// validateFraudStakeSlashingFactor validates the FraudStakeSlashingFactor param
func validateFraudStakeSlashingFactor(v interface{}) error {
	fraudStakeSlashingFactor, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if fraudStakeSlashingFactor > 1000000 {
		return fmt.Errorf("invalid fraudStakeSlashingFactor value: %d must be [0-1000000]", fraudStakeSlashingFactor)
	}
	_ = fraudStakeSlashingFactor

	return nil
}

// validateFraudSlashingAmount validates the FraudSlashingAmount param
func validateFraudSlashingAmount(v interface{}) error {
	fraudSlashingAmount, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	_ = fraudSlashingAmount

	return nil
}
