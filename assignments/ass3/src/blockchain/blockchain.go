package blockchain

import (
	"encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	// TODO
	block := Initial(1)
	length := len(chain.Chain)
	difficulty := chain.Chain[0].Difficulty

	if hex.EncodeToString(chain.Chain[0].PrevHash) != hex.EncodeToString(block.PrevHash) {
		return false
	}

	for i := 1; i < length; i++ {
		if chain.Chain[i].Difficulty != difficulty {
			return false
		}
		if chain.Chain[i].Generation != chain.Chain[i-1].Generation+1 {
			return false
		}
		if hex.EncodeToString(chain.Chain[i].PrevHash) != hex.EncodeToString(chain.Chain[i-1].Hash) {
			return false
		}
		// Other two requirements are met with the block ValidHash method
	}
	return true
}
