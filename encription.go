package aesanybytes

import (
	"crypto/aes"
	"errors"
	"fmt"

	"github.com/avran02/aes-any-bytes/internal"
)

var (
	emptyDataError = errors.New("secret data is empty")
)

// Encript secretData with key. If secretData is empty return emptyDataError
// If len of key is not equal to 16, 24 or 32 return error
func Encript(secretData, key []byte) ([]byte, error) {
	if len(secretData) == 0 {
		return []byte{}, emptyDataError
	}

	blocks, numBlocks := internal.SplitToBlocks(secretData)
	blocks, blockWasAdded := internal.AddPadding(blocks)
	if blockWasAdded {
		numBlocks++
	}
	encriptedBlocks := make([][]byte, numBlocks)
	for blockIdx := 0; blockIdx < numBlocks; blockIdx++ {
		encriptedBlock, err := internal.EncriptBlock(blocks[blockIdx], key)
		if err != nil {
			return []byte{}, fmt.Errorf("cant encript block: %w", err)
		}
		encriptedBlocks[blockIdx] = encriptedBlock
	}
	encriptedBytes := make([]byte, numBlocks*aes.BlockSize)
	for blockIdx := 0; blockIdx < numBlocks; blockIdx++ {
		for i := 0; i < aes.BlockSize; i++ {
			encriptedBytes[blockIdx*aes.BlockSize+i] = encriptedBlocks[blockIdx][i]
		}
	}
	return encriptedBytes, nil
}

// Decript secretData with key. If secretData is empty return emptyDataError
// If len of key is not equal to 16, 24 or 32 return error
func Decript(secretData, key []byte) ([]byte, error) {
	blocks, numBlocks := internal.SplitToBlocks(secretData)
	decriptedBlocks := make([][]byte, numBlocks)
	for blockIdx := 0; blockIdx < numBlocks; blockIdx++ {
		decriptedBlock, err := internal.DecriptBlock(blocks[blockIdx], key)
		if err != nil {
			return []byte{}, fmt.Errorf("cant encript block: %w", err)
		}
		decriptedBlocks[blockIdx] = decriptedBlock
	}
	decriptedBlocks = internal.RemovePadding(decriptedBlocks)

	decriptedBytesLen := (numBlocks-1)*aes.BlockSize + len(decriptedBlocks[numBlocks-1])
	decriptedBytes := make([]byte, decriptedBytesLen)
	for blockIdx := 0; blockIdx < numBlocks; blockIdx++ {
		for i := 0; i < len(decriptedBlocks[blockIdx]); i++ {
			decriptedBytes[blockIdx*aes.BlockSize+i] = decriptedBlocks[blockIdx][i]
		}
	}
	return decriptedBytes, nil
}
