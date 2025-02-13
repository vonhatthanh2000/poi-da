package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

func U8aConcat(arrays ...[]byte) []byte {
	return bytes.Join(arrays, nil)
}

func NumberToU8a(num uint64) []byte {
	return []byte{byte(num)} // Simplified for demonstration
}

func HexToU8a(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr[2:]) // Remove '0x' prefix
}

func IsHex(s string) bool {
	_, err := hex.DecodeString(s[2:])
	return err == nil
}

func IsBase58(s string) bool {
	// Placeholder for actual base58 validation
	return true
}

func Base58Decode(s string) ([]byte, error) {
	// Placeholder for actual base58 decoding
	return []byte(s), nil
}

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func Blake2AsU8a(data []byte) []byte {
	// Placeholder for BLAKE2 hashing
	return data // No actual hashing for this example
}

// Helper function to convert []byte to hex string
func ByteToHex(b []byte) string {
	if b == nil {
		return ""
	}
	return hex.EncodeToString(b)
}
