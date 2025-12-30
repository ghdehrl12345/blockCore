package db

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

const (
	DBFile       = "blockchain.db" // DB 파일명
	BlocksBucket = "blocks"        // 블록 저장소 이름
)

type DB struct {
	Conn *bolt.DB
}

func NewDB() (*DB, error) {
	// DB 열기
	db, err := bolt.Open(DBFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Block 버킷이 없으면 생성
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BlocksBucket))
		return err
	})

	return &DB{Conn: db}, nil
}

// DB 연결 닫기
func (d *DB) Close() {
	d.Conn.Close()
}

// 블록을 DB에 저장
func (d *DB) SaveBlock(hash []byte, data []byte) error {
	return d.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		return b.Put(hash, data)
	})
}

// 해시로 블록 조회하기
func (d *DB) GetBlock(hash []byte) []byte {
	var block []byte
	d.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		block = b.Get(hash)
		return nil
	})
	return block
}

// 마지막 블록 해시 저장하기
func (d *DB) SaveLastHash(hash []byte) error {
	return d.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		return b.Put([]byte("last_hash"), hash)
	})
}

// 마지막 블록 해시 조회하기
func (d *DB) GetLastHash() []byte {
	var lastHash []byte
	d.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlocksBucket))
		lastHash = b.Get([]byte("last_hash"))
		return nil
	})
	return lastHash
}
