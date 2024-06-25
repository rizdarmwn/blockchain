package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Hash         []byte
	Prev         []byte
	Transactions []*Transaction
	Timestamp    int64
	Nonce        int
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{Transactions: transactions, Prev: prevHash, Timestamp: time.Now().Unix()}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic("encoder error")
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewBuffer(d))

	err := decoder.Decode(&block)
	if err != nil {
		panic("decoder error")
	}
	return &block
}
