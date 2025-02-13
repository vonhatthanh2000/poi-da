package poi

import (
	"bytes"
	"errors"
	"fmt"
	"layerg-poi-da/internal/utils"
)

// Constants
var DEFAULT_BLOCK_HASH, _ = utils.HexToU8a("0x00")
var NULL_OPERATION_HASH, _ = utils.HexToU8a("0x00")

// PoiBlock struct
type PoiBlock struct {
	ID                uint64
	ChainBlockHash    []byte
	OperationHashRoot []byte
	ProjectID         string
	ParentHash        []byte
	Hash              []byte
}

// poiBlockHash function
func poiBlockHash(id uint64, chainBlockHash []byte, operationHashRoot []byte, projectId string, parentHash []byte) ([]byte, error) {
	if id == 0 || projectId == "" {
		return nil, errors.New("Proof of index: cannot generate block hash")
	}

	return utils.Blake2AsU8a(utils.U8aConcat(
		utils.NumberToU8a(id),
		chainBlockHash,
		operationHashRoot,
		[]byte(projectId),
		parentHash,
	)), nil
}

// CreateFirstBlock function for the first PoiBlock
func CreateFirstBlock(id uint64, chainBlockHash interface{}, operationHashRoot []byte, projectId string) (*PoiBlock, error) {
	if operationHashRoot == nil && chainBlockHash == nil {
		return nil, errors.New("Create first POI block failed, chainBlockHash must be defined.")
	}

	var _chainBlockHash []byte
	var err error

	switch v := chainBlockHash.(type) {
	case string:
		if utils.IsHex(v) {
			_chainBlockHash, err = utils.HexToU8a(v)
		} else if utils.IsBase58(v) && !bytes.Contains([]byte(v), []byte("=")) {
			_chainBlockHash, err = utils.Base58Decode(v)
		} else if utils.IsBase64(v) {
			_chainBlockHash, err = utils.Base64Decode(v)
		} else if v == "" {
			_chainBlockHash = nil
		} else {
			return nil, fmt.Errorf("Unable to create first PoiBlock, chainBlockHash was not valid. Received: %q", v)
		}
	case []byte:
		_chainBlockHash = v
	default:
		return nil, errors.New("Invalid chainBlockHash type")
	}

	hash, err := poiBlockHash(id, _chainBlockHash, operationHashRoot, projectId, nil)
	if err != nil {
		return nil, err
	}

	return &PoiBlock{id, _chainBlockHash, operationHashRoot, projectId, nil, hash}, nil
}

// CreateSubsequentBlock function for subsequent PoiBlocks
func CreateSubsequentBlock(id uint64, chainBlockHash interface{}, operationHashRoot []byte, projectId string, parentHash []byte) (*PoiBlock, error) {
	if operationHashRoot == nil && chainBlockHash == nil {
		return nil, errors.New("Create subsequent POI block failed, chainBlockHash must be defined.")
	}

	var _chainBlockHash []byte
	var err error

	switch v := chainBlockHash.(type) {
	case string:
		if utils.IsHex(v) {
			_chainBlockHash, err = utils.HexToU8a(v)
		} else if utils.IsBase58(v) && !bytes.Contains([]byte(v), []byte("=")) {
			_chainBlockHash, err = utils.Base58Decode(v)
		} else if utils.IsBase64(v) {
			_chainBlockHash, err = utils.Base64Decode(v)
		} else if v == "" {
			_chainBlockHash = nil
		} else {
			return nil, fmt.Errorf("Unable to create subsequent PoiBlock, chainBlockHash was not valid. Received: %q", v)
		}
	case []byte:
		_chainBlockHash = v
	default:
		return nil, errors.New("Invalid chainBlockHash type")
	}

	hash, err := poiBlockHash(id, _chainBlockHash, operationHashRoot, projectId, parentHash)
	if err != nil {
		return nil, err
	}

	return &PoiBlock{id, _chainBlockHash, operationHashRoot, projectId, parentHash, hash}, nil
}
