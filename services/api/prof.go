package api

import (
	"github.com/attestantio/go-builder-client/api/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

type ProfBundleRequest struct {
	slot         uint64
	Transactions [][]byte `ssz-max:"1048576,1073741824"`
	bundleHash   phase0.Hash32
}

func NewEmptyProfBundleRequest(slot uint64) *ProfBundleRequest {
	return &ProfBundleRequest{
		slot:         slot,
		Transactions: [][]byte{},
		bundleHash:   phase0.Hash32{},
	}
}

func (p *ProfBundleRequest) Slot() (uint64, error) {
	return p.slot, nil
}

func (p *ProfBundleRequest) BundleHash() (phase0.Hash32, error) {
	return p.bundleHash, nil
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
	Value     *uint256.Int
	NewHeader *types.Header
}

// type AppendProfResponse struct {
// 	value     *uint256.Int
// 	blockHash phase0.Hash32
// 	bid       *builderSpec.VersionedSignedBuilderBid
// }

// func (a *AppendProfResponse) BlockHash() (phase0.Hash32, error) {
// 	return a.blockHash, nil
// }
