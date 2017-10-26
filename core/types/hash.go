// Package types contains data types related to go-loki consensus.
package types

import "crypto/sha256"

// Hash defines a  hash
type Hash [sha256.Size]byte
