package internal_test

import (
	"crypto/aes"
	"fmt"
	"testing"

	"github.com/avran02/aes-any-bytes/internal"
	"github.com/stretchr/testify/assert"
)

func getBlock(lenth int) []byte {
	if lenth > aes.BlockSize {
		panic("invalid block size")
	}
	block := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		block[i] = byte(i)
	}
	return block
}

func getBlocks(numBlocks, lenth int) [][]byte {
	var blocks [][]byte
	if lenth == 0 {
		blocks = make([][]byte, numBlocks-1)
	} else {
		blocks = make([][]byte, numBlocks)
	}
	for i := 0; i < numBlocks-1; i++ {
		blocks[i] = getBlock(aes.BlockSize)
	}
	if lenth == 0 {
		return blocks
	}
	blocks[numBlocks-1] = getBlock(lenth)
	return blocks
}

func TestAddPadding(t *testing.T) {
	assert := assert.New(t)
	for lastBlockSize := 0; lastBlockSize <= aes.BlockSize; lastBlockSize++ {
		t.Run("lastBlockSize: "+fmt.Sprint(lastBlockSize), func(t *testing.T) {
			blocks := getBlocks(3, lastBlockSize)
			lastBlock := blocks[len(blocks)-1]
			paddedBlocks, blockAdded := internal.AddPadding(blocks)

			if blockAdded {
				assert.Equal(len(blocks), len(paddedBlocks)-1)
				assert.Equal(paddedBlocks[len(paddedBlocks)-1], internal.DefaultPaddingBlock)
				return
			}

			lastPaddedBlock := paddedBlocks[len(paddedBlocks)-1]
			paddingByte := lastPaddedBlock[aes.BlockSize-1]

			paddedBytes := make([]byte, paddingByte)
			for i := byte(0); i < paddingByte; i++ {
				paddedBytes[i] = paddingByte
			}

			exceptedLastBlock := append(lastBlock[:0:0], lastBlock[:aes.BlockSize-paddingByte]...)
			exceptedLastBlock = append(exceptedLastBlock, paddedBytes...)
			assert.Equal(len(blocks), len(paddedBlocks))
			assert.Equal(lastPaddedBlock, exceptedLastBlock)
		})
	}
}
