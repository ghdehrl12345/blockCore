package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

// Block 구조체
type Block struct {
	Timestamp     int64
	Data          []byte // 나중에 TX로 대체할거임
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock 함수(블록 생성, 해시 계산)
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          nil,
		Nonce:         0,
	}
	block.Hash = block.CalculateHash()
	return block
}

// 블록 해시 계산 함수
func (b *Block) CalculateHash() []byte {
	timestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(timestamp, uint64(b.Timestamp)) //

	headers := bytes.Join(
		[][]byte{
			b.PrevBlockHash,
			b.Data,
			timestamp,
		},
		[]byte{},
	)

	hash := sha256.Sum256(headers)
	return hash[:]
}
