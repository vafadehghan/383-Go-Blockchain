package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

//Initial Creates a new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	block := new(Block)
	block.Difficulty = difficulty
	for i := 0; i < 32; i++ {
		block.PrevHash = append(block.PrevHash, '\x00')
	}
	return *block
}

// Next Creates a new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	block := new(Block)
	block.Difficulty = prev_block.Difficulty
	block.PrevHash = prev_block.Hash
	block.Generation = prev_block.Generation + 1
	block.Data = data
	return *block
}

// CalcHash Calculates the block's hash.
func (blk Block) CalcHash() []byte {
	str := fmt.Sprintf("%s:%d:%d:%s:%d", hex.EncodeToString(blk.PrevHash), blk.Generation, blk.Difficulty, blk.Data, blk.Proof)
	sha := sha256.Sum256([]byte(str))
	return sha[:]
}

// ValidHash checks Is this block's hash valid?
func (blk Block) ValidHash() bool {
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	for i := (len(blk.Hash) - int(nBytes)); i < len(blk.Hash); i++ {
		if blk.Hash[i] != '\x00' {
			return false
		}
	}
	return blk.Hash[(len(blk.Hash)-int(nBytes)-1)]%(1<<nBits) == 0
}

// SetProof Sets the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
