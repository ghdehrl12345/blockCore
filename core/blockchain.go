package core

import (
	"fmt"
	"log"

	"github.com/ghdehrl12345/blockCore/db"
)

// DB로 블록 관리
type Blockchain struct {
	lastHash []byte // 마지막 블록 해시
	DB       *db.DB // DB 연결
	State    *StateDB
}

// 기존 체인을 로드하거나 새로 생성
func NewBlockchain(minerAddress []byte) *Blockchain {
	database, _ := db.NewDB()
	lastHash := database.GetLastHash()
	state := NewStateDB(database)

	if lastHash == nil {
		// 채굴자에게 보상
		coinbase := NewCoinbaseTX(minerAddress, "Genesis Block")
		genesis := NewBlock([]*Transaction{coinbase}, []byte{})

		database.SaveBlock(genesis.Hash, genesis.Serialize())
		database.SaveLastHash(genesis.Hash)
		lastHash = genesis.Hash

		// 채굴자에게 50 코인
		account := state.GetAccount(minerAddress)
		account.Balance += 50
		state.UpdateAccount(minerAddress, account)
	}

	return &Blockchain{
		lastHash: lastHash,
		DB:       database,
		State:    state,
	}
}

// 트랜잭션들을 검증하고 블록에 포함
func (bc *Blockchain) MineBlock(transactions []*Transaction, minerAddress []byte) *Block {
	// 트랜잭션 검증
	for _, tx := range transactions {
		if !bc.VerifyTransaction(tx) {
			log.Panic("Invalid transaction")
		}
	}

	// 채굴 보상 추가
	coinbase := NewCoinbaseTX(minerAddress, "")
	transactions = append([]*Transaction{coinbase}, transactions...)

	// 블록 생성
	newBlock := NewBlock(transactions, bc.lastHash)

	// 상태 업데이트
	bc.ApplyTransactions(transactions)

	// DB 저장
	bc.DB.SaveBlock(newBlock.Hash, newBlock.Serialize())
	bc.DB.SaveLastHash(newBlock.Hash)
	bc.lastHash = newBlock.Hash

	return newBlock
}

// 트랜잭션이 유효한지 검증
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	sender := bc.State.GetAccount(tx.From)

	// 잔액 확인
	if sender.Balance < tx.Amount {
		fmt.Printf("Insufficient balance: %d < %d\n", sender.Balance, tx.Amount)
		return false
	}

	// Nonce 확인
	if sender.Nonce != tx.Nonce {
		fmt.Printf("Invalid nonce: expected %d, got %d\n", sender.Nonce, tx.Nonce)
		return false
	}

	return true
}

// 트랜잭션들을 상태에 적용
func (bc *Blockchain) ApplyTransactions(transactions []*Transaction) {
	for _, tx := range transactions {
		if tx.IsCoinbase() {
			// 받는 사람에게만 추가
			receiver := bc.State.GetAccount(tx.To)
			receiver.Balance += tx.Amount
			bc.State.UpdateAccount(tx.To, receiver)
		} else {
			// 보내는 사람에서 빼고 받는 사람에게 추가
			sender := bc.State.GetAccount(tx.From)
			sender.Balance -= tx.Amount
			sender.Nonce++
			bc.State.UpdateAccount(tx.From, sender)

			receiver := bc.State.GetAccount(tx.To)
			receiver.Balance += tx.Amount
			bc.State.UpdateAccount(tx.To, receiver)
		}
	}
}

// 주소의 잔액을 반환
func (bc *Blockchain) GetBalance(address []byte) int {
	account := bc.State.GetAccount(address)
	return account.Balance
}

// 마지막 블록 해시를 반환하기
func (bc *Blockchain) GetLastHash() []byte {
	return bc.lastHash
}

// DB 연결 닫기
func (bc *Blockchain) Close() {
	bc.DB.Close()
}

// 블록체인 순회
type Iterator struct {
	currentHash []byte
	db          *db.DB
}

// 마지막 블록부터 시작하는 이터레이터 생성
func (bc *Blockchain) NewIterator() *Iterator {
	return &Iterator{
		currentHash: bc.lastHash,
		db:          bc.DB,
	}
}

// 다음 블록 반환
func (it *Iterator) Next() *Block {
	blockData := it.db.GetBlock(it.currentHash)
	if blockData == nil {
		return nil
	}

	block := DeserializeBlock(blockData)
	it.currentHash = block.PrevBlockHash

	return block
}
