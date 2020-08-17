package blockchain

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestPrevHash tests that the hash of the previous block is equal to the prevHash value of the current block
func TestPrevHash(t *testing.T) {
	b0 := Initial(2)
	b1 := b0.Next("Next Block")
	assert.Equal(t, b0.Hash, b1.PrevHash, "These hashes should be identical")
}

//TestSetProof tests that the setproof funciton works properly
func TestSetProof(t *testing.T) {
	b0 := Initial(19)
	b0.SetProof(87745)
	b1 := b0.Next("hash example 1234")
	b1.SetProof(1407891)
	assert.Equal(t, b1.ValidHash(), true, "This is a valid hash")
	b1.SetProof(346082)
	assert.Equal(t, b1.ValidHash(), false, "This is not a valid hash")
}

//TestMinedHash tests that the mined hash value and proof are correct.
func TestMinedHash(t *testing.T) {
	b0 := Initial(16)
	b0.Mine(1)
	assert.Equal(t, b0.Proof, uint64(56231), "These proofs should be identical")
	assert.Equal(t, hex.EncodeToString(b0.Hash), "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000", "These Hashes should be identical")
}

//TestValidity tests that the ValidHash() function works correctly
func TestValidity(t *testing.T) {
	b0 := Initial(16)
	b1 := b0.Next("Next block")
	b1.Mine(1)
	assert.Equal(t, b1.ValidHash(), true, "This should be true")

}

//TestBlockData tests that the generation, difficulty, hash, and data of a block is correct
func TestBlockData(t *testing.T) {
	b0 := Initial(16)
	b1 := b0.Next("Next block")

	assert.Equal(t, b1.Generation, uint64(1), "Wrong generation")
	assert.Equal(t, b1.Difficulty, uint8(16), "Wrong Difficulty")
	assert.Equal(t, b0.Hash, b1.PrevHash, "Hashes dont match up")
	assert.Equal(t, b1.Data, "Next block", "Data doesnt matche up")
}
