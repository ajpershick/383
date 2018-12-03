package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: some useful tests of Blocks
func TestInitialBlock(t *testing.T) {
	block := Initial(5)

	assert.Equal(t, block, Block{
		PrevHash:   []byte("\x00"),
		Generation: 0,
		Difficulty: 5,
		Data:       "",
	})

	assert.Equal(t, block.PrevHash, []byte("\x00"))
}

func TestHashing(t *testing.T) {
	block := Initial(16)
	assert.Equal(t, block.ValidHash(), false)
	block.SetProof(56231)
	assert.Equal(t, block.ValidHash(), true)
}
