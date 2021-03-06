package blockchain

import (
	"fmt"
	"work_queue"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	worker work_queue.Worker
	blk    Block
	start  uint64
	end    uint64
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	q := new(work_queue.WorkQueue)
	fmt.Printf("Queue created for MineRange %p\n", q)
	chunkSize := (end-start)/chunks + 1
	fmt.Printf("Chunksize = %d\n", chunkSize)

	// for i := start
	return MiningResult{Found: false}
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

func (worker miningWorker) Run() MiningResult {
	for i := worker.start; i <= worker.end; i++ {
		worker.blk.SetProof(i)
		if worker.blk.ValidHash() {
			return MiningResult{worker.blk.Proof, true}
		}
	}
	return MiningResult{Found: false}
}
