package blockchain

import (
	"encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

// Add adds a block to the blockchain
func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	} else {
		if chain.IsValid() {
			chain.Chain = append(chain.Chain, blk)
		} else {
			panic("adding block with invalid chain")

		}
	}
}

//IsValid checks that the blockchain is valid
func (chain Blockchain) IsValid() bool {
	initial := chain.Chain[0]
	for a := range initial.PrevHash {
		if a != '\x00' {
			return false
		}
	}
	if initial.Generation != 0 {
		return false
	}

	difficulty := initial.Difficulty
	generation := uint64(1)
	previousHash := initial.Hash

	for _, block := range chain.Chain[1:] {
		if block.Difficulty != difficulty {
			return false
		}
		if block.Generation != generation {
			return false
		}
		if hex.EncodeToString(block.PrevHash) != hex.EncodeToString(previousHash) {
			return false
		}
		if hex.EncodeToString(block.Hash) != hex.EncodeToString(block.CalcHash()) {
			return false
		}
		if !block.ValidHash() {
			return false
		}
		generation++
		difficulty = block.Difficulty
		previousHash = block.Hash
	}
	return true
}
