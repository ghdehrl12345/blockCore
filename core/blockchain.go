package core

import (
	"github.com/ghdehrl12345/blockCore/db"
)

// DB로 블록 관리
type Blockchain struct {
	lastHash []byte // 마지막 블록 해시
	DB       *db.DB // DB 연결
}

// 기존 체인을 로드하거나 새로 생성
func NewBlockchain() *Blockchain {
	database, _ := db.NewDB()
	lastHash := database.GetLastHash()

	// DB에 체인이 없으면 제네시스 블록 생성
	if lastHash == nil {
		genesis := NewBlock("Genesis Block", []byte{})
		database.SaveBlock(genesis.Hash, genesis.Serialize())
		database.SaveLastHash(genesis.Hash)
		lastHash = genesis.Hash
	}

	return &Blockchain{
		DB:       database,
		lastHash: lastHash,
	}
}

// 새 블록을 채굴하고 DB에 저장
func (bc *Blockchain) AddBlock(data string) *Block {
	newBlock := NewBlock(data, bc.lastHash)

	bc.DB.SaveBlock(newBlock.Hash, newBlock.Serialize())
	bc.DB.SaveLastHash(newBlock.Hash)
	bc.lastHash = newBlock.Hash

	return newBlock
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
