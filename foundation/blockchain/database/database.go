package database

import (
	"github.com/ardanlabs/blockchain/foundation/blockchain/genesis"
	"sync"
)

// Database manages data related to accounts who have transacted on the blockchain.
type Database struct {
	mu      sync.RWMutex
	genesis genesis.Genesis
	// latestBlock Block
	accounts map[AccountID]Account
}
