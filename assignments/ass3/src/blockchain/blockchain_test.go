package blockchain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: some useful tests of Blocks
func TestInitialBlock(t *testing.T) {
	block := Initial(5)

	assert.Equal(t, block, Block{
		PrevHash:   []byte(strings.Repeat("\x00", 32)),
		Generation: 0,
		Difficulty: 5,
		Data:       "",
	})
}

func TestHashing(t *testing.T) {
	block := Initial(16)
	assert.Equal(t, block.ValidHash(), false)
	block.SetProof(56231)
	assert.Equal(t, block.ValidHash(), true)
}

func TestAppendingToChain(t *testing.T) {
	chain := new(Blockchain)
	block := Initial(16)
	block.SetProof(56231)
	chain.Add(block)

	assert.Equal(t, chain.Chain[0], block)
}
