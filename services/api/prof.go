package api

import (
	builderApi "github.com/attestantio/go-builder-client/api"
	builderSpec "github.com/attestantio/go-builder-client/spec"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

type ProfBundleRequest struct {
	slot         uint64
	Transactions [][]byte `ssz-max:"1048576,1073741824"`
	bundleHash   phase0.Hash32
}

func (p *ProfBundleRequest) Slot() (uint64, error) {
	return p.slot, nil
}

func (p *ProfBundleRequest) BundleHash() (phase0.Hash32, error) {
	return p.bundleHash, nil
}

type ProfSimReq struct {
	pbsPayload *builderApi.VersionedSubmitBlindedBlockResponse
	profBundle *ProfBundleRequest
}

type ProfSimResp struct {
	value     *uint256.Int
	blockHash phase0.Hash32
}

type AppendProfResponse struct {
	value     *uint256.Int
	blockHash phase0.Hash32
	bid       *builderSpec.VersionedSignedBuilderBid
}

func (a *AppendProfResponse) BlockHash() (phase0.Hash32, error) {
	return a.blockHash, nil
}
