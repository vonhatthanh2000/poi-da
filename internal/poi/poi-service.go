package poi

import (
	"layerg-poi-da/internal/utils"
)

// ProofOfIndexHuman struct
type ProofOfIndexHuman struct {
	ID                uint64 `json:"id"`
	ChainBlockHash    string `json:"chainBlockHash,omitempty"`
	OperationHashRoot string `json:"operationHashRoot,omitempty"`
	ProjectID         string `json:"projectId"`
	ParentHash        string `json:"parentHash,omitempty"`
	Hash              string `json:"hash,omitempty"`
}

// PoiToHuman function
func PoiToHuman(proofOfIndex PoiBlock) ProofOfIndexHuman {
	return ProofOfIndexHuman{
		ID:                proofOfIndex.ID,
		ChainBlockHash:    utils.ByteToHex(proofOfIndex.ChainBlockHash),
		OperationHashRoot: utils.ByteToHex(proofOfIndex.OperationHashRoot),
		ProjectID:         proofOfIndex.ProjectID,
		ParentHash:        utils.ByteToHex(proofOfIndex.ParentHash),
		Hash:              utils.ByteToHex(proofOfIndex.Hash),
	}
}
