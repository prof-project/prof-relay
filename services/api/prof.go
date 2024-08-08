package api

import (
	"github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// ToDo : bellatrix.Transaction unmarshalJson panics if the transaction is not pre allocated.
type ProfBundleRequest struct {
	Slot         uint64   `json:"slot"`
	Transactions []string `json:"transactions" ssz-max:"1048576,1073741824"` // 1MB, 1GB
}

type ProfBundleHTTPRequest struct {
	Transactions []string `json:"transactions" ssz-max:"1048576,1073741824"` // 1MB, 1GB
}

func NewEmptyProfBundleRequest(slot uint64) *ProfBundleRequest {
	return &ProfBundleRequest{
		Slot:         slot,
		Transactions: []string{},
	}
}

func NewProfBundleRequest(slot uint64, txs []string) *ProfBundleRequest {
	return &ProfBundleRequest{
		Slot:         slot,
		Transactions: txs,
	}
}

type ProfSimReq struct {
	PbsPayload            *deneb.ExecutionPayloadAndBlobsBundle
	ProfBundle            *ProfBundleRequest
	ParentBeaconBlockRoot *phase0.Root `json:"parent_beacon_block_root"`
	RegisteredGasLimit    uint64       `json:"registered_gas_limit,string"`
	ProposerFeeRecipient  common.Address
}

func NewProfSimReq(pbsPayload *deneb.ExecutionPayloadAndBlobsBundle, profBundle *ProfBundleRequest, parentBeaconBlockRoot *phase0.Root /*, registeredGasLimit uint64 */) *ProfSimReq {
	// TODO : check this request
	return &ProfSimReq{
		PbsPayload:            pbsPayload,
		ProfBundle:            profBundle,
		ParentBeaconBlockRoot: parentBeaconBlockRoot,
		RegisteredGasLimit:    0, // not checked yet
		ProposerFeeRecipient:  common.Address(pbsPayload.ExecutionPayload.FeeRecipient),
	}
}

type ProfSimResp struct {
	Value            *uint256.Int
	ExecutionPayload *deneb.ExecutionPayloadAndBlobsBundle
}

// type AppendProfResponse struct {
// 	value     *uint256.Int
// 	blockHash phase0.Hash32
// 	bid       *builderSpec.VersionedSignedBuilderBid
// }

// func (a *AppendProfResponse) BlockHash() (phase0.Hash32, error) {
// 	return a.blockHash, nil
// }
