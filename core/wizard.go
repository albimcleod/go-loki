package core

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/albimcleod/go-loki/core/types"
)

func MakeGenesis() *types.Block {
	s := "0"
	var a [sha256.Size]byte
	copy(a[:], s)
	return types.NewBlock(0, time.Now(), []byte("Genesis Block"), a)
}

func NextBlock(last *types.Block) *types.Block {
	return types.NewBlock(last.Index+1, time.Now(), []byte(fmt.Sprintf("New Block: %v", last.Index+1)), last.Hash())
}
