package blockchain

import "work_queue"

type miningWorker struct {
	block Block // use this block to pass info
	start uint64
	end   uint64
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	wq := work_queue.Create(uint(workers), uint(chunks))
	chunk := ((end - start) / chunks)
	if end < chunks {
		mineWorker := miningWorker{blk, start, end}
		wq.Enqueue(mineWorker)
	} else {
		for i := start; i < end; i += chunk {
			mineWorker := miningWorker{blk, i, i + chunk}

			wq.Enqueue(mineWorker)
		}
	}

	mr := new(MiningResult)
	for res := range wq.Results {
		if res.(MiningResult).Found == true {
			wq.Shutdown()
			return res.(MiningResult)
		}
	}
	return *mr

}

func (miningworker miningWorker) Run() interface{} {
	mr := new(MiningResult)
	for i := miningworker.start; i <= miningworker.end; i++ {
		miningworker.block.SetProof(i)
		if miningworker.block.ValidHash() {
			mr.Proof = i
			mr.Found = true
			return *mr
		}
	}
	mr.Found = false
	return *mr
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
