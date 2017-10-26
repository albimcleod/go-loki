// Package types contains data types related to go-loki consensus.
package types

// BlockChain defines the blockchain
type BlockChain []Block

//Push will append a block to the blockchain
func (bc BlockChain) Push(b Block) {
	bc = append(bc, b)
}
