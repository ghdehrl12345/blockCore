package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	ID     []byte // 트랜잭션 해시
	From   []byte // 보내는 주소
	To     []byte // 받는 주소
	Amount int    // 금액
	Nonce  int    // 이중 지불 방지용 (송신자 기준 순서)
}

func NewTransaction(from, to []byte, amount int, nonce int) *Transaction {
	tx := &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
		Nonce:  nonce,
	}
	tx.ID = tx.Hash()
	return tx
}

// 트랜잭션의 해시를 계산합니다.
func (tx *Transaction) Hash() []byte {
	var encoded bytes.Buffer
	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(encoded.Bytes())
	return hash[:]
}

// 채굴 보상 트랜잭션인지 확인
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.From) == 0 && len(tx.To) == 0
}

// 채굴 보상 트랜잭션 생성
func NewCoinbaseTX(to []byte, data string) *Transaction {
	tx := &Transaction{
		From:   nil,
		To:     to,
		Amount: 50, // 채굴 보상
		Nonce:  0,
	}
	tx.ID = tx.Hash()
	return tx
}
