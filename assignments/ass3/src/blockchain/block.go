package blockchain

import "crypto/sha256"
import "encoding/hex"
import "fmt"
import "strings"

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	initialBlock := Block{
		PrevHash:   []byte(strings.Repeat("\x00", 32)),
		Generation: 0,
		Difficulty: difficulty,
		Data:       "",
	}
	return initialBlock
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	block := Block{
		PrevHash:   prev_block.Hash,
		Generation: prev_block.Generation + 1,
		Difficulty: prev_block.Difficulty,
		Data:       data,
	}
	return block
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	password := fmt.Sprintf("%s:%d:%d:%s:%d",
		hex.EncodeToString(blk.PrevHash),
		blk.Generation,
		blk.Difficulty,
		blk.Data,
		blk.Proof)
	println("password: %s\n", password)

	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	fmt.Printf("Hash is: %x\n", hash)
	return hash
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	length := len(blk.Hash)

	if (length-1)-int(nBytes) < 0 {
		return false
	}

	println("length of hash is: %d", length)
	for i := length - 1; i > (length-1)-int(nBytes); i-- {
		println("blk.Hash[%d] = %x", i, hex.EncodeToString(blk.Hash))
		if blk.Hash[i] != '\x00' {
			return false
		}
	}

	if blk.Hash[length-int(nBytes)]%(1<<nBits) != 0 {
		return false
	}

	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
