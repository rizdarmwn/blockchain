package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Hash      []byte
	Prev      []byte
	Data      []byte
	Timestamp int64
	Nonce     int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{Data: []byte(data), Prev: prevHash, Timestamp: time.Now().Unix()}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) CreateHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.Prev, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
