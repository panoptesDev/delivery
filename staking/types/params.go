package types

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/maticnetwork/heimdall/params/subspace"
)

const (

	// DefaultProposerBonusPercent - Proposer Signer Reward Ratio
	DefaultStakingBufferTime = 600 * time.Second
)

// Parameter keys
var (
	KeyStakingBufferTime = []byte("StakingBufferTime")
)

var _ subspace.ParamSet = &Params{}

// Params defines the parameters for the auth module.
type Params struct {
	StakingBufferTime time.Duration `json:"staking_buffer_time" yaml:"staking_buffer_time"`
}

// NewParams creates a new Params object
func NewParams(stakingBufferTime time.Duration) Params {
	return Params{
		StakingBufferTime: stakingBufferTime,
	}
}

// ParamKeyTable for auth module
func ParamKeyTable() subspace.KeyTable {
	return subspace.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return subspace.ParamSetPairs{
		{KeyStakingBufferTime, &p.StakingBufferTime},
	}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{StakingBufferTime: DefaultStakingBufferTime}
}

// String implements the stringer interface.
func (p Params) String() string {
	var sb strings.Builder
	sb.WriteString("Params: \n")
	sb.WriteString(fmt.Sprintf("CheckpointBufferTime: %s\n", p.StakingBufferTime))
	return sb.String()
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if p.StakingBufferTime == 0 {
		return fmt.Errorf("StakingBufferTime, AvgCheckpointLength should be non-zero")
	}
	return nil
}
