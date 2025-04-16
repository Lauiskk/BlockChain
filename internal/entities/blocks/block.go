package blocks

import (
	"RedisLike/internal/entities/transaction"
	"bytes"
	"encoding/gob"
	"github.com/boltdb/bolt"
	"math/big"
)

const blocksBucket = "blocks"
const targetBits = 24

type Blockchain struct {
	tip []byte
	Db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}
type Block struct {
	Timestamp     int64
	Transactions  []*transaction.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}
type CLI struct {
	Bc *Blockchain
}
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
