package internal

import (
	"crypto/aes"
	"fmt"
	"math"
)

var (
	// is used for padding when last block size is multiple of aes.BlockSize
	DefaultPaddingBlock = []byte{
		0x10, 0x10, 0x10, 0x10,
		0x10, 0x10, 0x10, 0x10,
		0x10, 0x10, 0x10, 0x10,
		0x10, 0x10, 0x10, 0x10,
	}
)

// Create cipher and encrypt data with it
func EncriptBlock(secretDataBlock, key []byte) ([]byte, error) {
	if len(secretDataBlock)%aes.BlockSize != 0 {
		secretDataBlock = append(secretDataBlock, make([]byte, aes.BlockSize-len(secretDataBlock)%aes.BlockSize)...)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("cant create cipher: %w", err)
	}
	encriptedData := make([]byte, len(secretDataBlock))
	c.Encrypt(encriptedData, secretDataBlock)
	return encriptedData, nil
}

// Create cipher and decrypt data with it
func DecriptBlock(secretDataBlock, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, fmt.Errorf("cant create cipher: %w", err)
	}
	decriptedData := make([]byte, len(secretDataBlock))
	c.Decrypt(decriptedData, secretDataBlock)
	return decriptedData, nil
}

// Split secretData to blocks with len equal to aes.BlockSize
func SplitToBlocks(secretData []byte) ([][]byte, int) {
	numBlocksFloat := float64(len(secretData)) / float64(aes.BlockSize)
	numBlocks := int(math.Ceil(numBlocksFloat))
	blocks := make([][]byte, numBlocks)
	for i := 0; i < numBlocks-1; i++ {
		blocks[i] = secretData[i*aes.BlockSize : (i+1)*aes.BlockSize]
	}
	blocks[numBlocks-1] = secretData[(numBlocks-1)*aes.BlockSize:]
	return blocks, numBlocks
}

// Add padding to last block. If last block len is equal to aes.BlockSize, add default padding and return true.
// If not block will be padded with numAddedBytes = aes.BlockSize - lastBlockLen.
func AddPadding(blocks [][]byte) ([][]byte, bool) {
	numBlocks := len(blocks)
	lastBlock := blocks[numBlocks-1]
	newLastBlock := make([]byte, aes.BlockSize)
	copy(newLastBlock, lastBlock)
	lastBlockLen := len(lastBlock)
	if lastBlockLen == aes.BlockSize {
		blocks = append(blocks, DefaultPaddingBlock)
		return blocks, true
	}
	numAddedBytes := aes.BlockSize - lastBlockLen
	for i := lastBlockLen; i < aes.BlockSize; i++ {
		newLastBlock[i] = byte(numAddedBytes)
	}
	blocks[numBlocks-1] = newLastBlock
	return blocks, false
}

// Remove padding from last block. It takes last byte from last block and removes num bytes equal to last byte.
func RemovePadding(blocks [][]byte) [][]byte {
	lastBlock := blocks[len(blocks)-1]
	numAddedBytes := lastBlock[aes.BlockSize-1]
	lastBlock = lastBlock[:aes.BlockSize-int(numAddedBytes)]
	newLastBlock := append(lastBlock[:0:0], lastBlock...) // nedded to strip unused capacity
	blocks[len(blocks)-1] = newLastBlock
	return blocks
}
