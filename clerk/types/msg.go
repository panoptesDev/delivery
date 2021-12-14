package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/maticnetwork/heimdall/types"
)

// MsgEventRecord - state msg
type MsgEventRecord struct {
	From            types.HeimdallAddress `json:"from"`
	TxHash          types.HeimdallHash    `json:"tx_hash"`
	LogIndex        uint64                `json:"log_index"`
	BlockNumber     uint64                `json:"block_number"`
	ContractAddress types.HeimdallAddress `json:"contract_address"`
	Data            types.HexBytes        `json:"data"`
	ID              uint64                `json:"id"`
	ChainID         string                `json:"bor_chain_id"`
	RootChainType   string                `json:"root_chain_type"`
}

var _ sdk.Msg = MsgEventRecord{}

// NewMsgEventRecord - construct state msg
func NewMsgEventRecord(
	from types.HeimdallAddress,
	txHash types.HeimdallHash,
	logIndex uint64,
	blockNumber uint64,
	id uint64,
	contractAddress types.HeimdallAddress,
	data types.HexBytes,
	chainID string,
	rootChainType string,

) MsgEventRecord {
	return MsgEventRecord{
		From:            from,
		TxHash:          txHash,
		LogIndex:        logIndex,
		BlockNumber:     blockNumber,
		ID:              id,
		ContractAddress: contractAddress,
		Data:            data,
		ChainID:         chainID,
		RootChainType:   rootChainType,
	}
}

// Route Implements Msg.
func (msg MsgEventRecord) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgEventRecord) Type() string { return "event-record" }

// ValidateBasic Implements Msg.
func (msg MsgEventRecord) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}

	if msg.TxHash.Empty() {
		return sdk.ErrInvalidAddress("missing tx hash")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEventRecord) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgEventRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{types.HeimdallAddressToAccAddress(msg.From)}
}

// GetTxHash Returns tx hash
func (msg MsgEventRecord) GetTxHash() types.HeimdallHash {
	return msg.TxHash
}

// GetLogIndex Returns log index
func (msg MsgEventRecord) GetLogIndex() uint64 {
	return msg.LogIndex
}

// getChainType  return tron type

func (msg MsgEventRecord) GetRootChainType() string {
	return msg.RootChainType
}

// GetSideSignBytes returns side sign bytes
func (msg MsgEventRecord) GetSideSignBytes() []byte {
	return nil
}
