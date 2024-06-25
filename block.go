package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
