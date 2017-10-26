// Package types contains data types related to go-loki consensus.
package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

// Block defines ...
type Block struct {
	Index        int64
	Timestamp    time.Time
	Data         []byte
	PreviousHash Hash
}

// Hash defines ...
func (b *Block) Hash() Hash {
	return rlpHash([]interface{}{
		b.Index,
		b.Timestamp,
		b.Data,
		b.PreviousHash,
	})
}

func rlpHash(x interface{}) Hash {
	b := []byte(fmt.Sprintf("%v", x))
	sum := sha256.Sum256(b)
	return sum
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewBlock(index int64, ts time.Time, data []byte, ph Hash) *Block {
	return &Block{
		Index:        index,
		Timestamp:    ts,
		Data:         data,
		PreviousHash: ph,
	}
}
