package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
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

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

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

// DB에 블록을 저장하려면 Block 구조체를 바이트 배열로 변환해야함
// 블록을 바이트 배열로 변환
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// 바이트 배열을 블록으로 복원
func DeserializeBlock(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
