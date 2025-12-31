package core

import (
	bolt "go.etcd.io/bbolt"

	"bytes"
	"encoding/gob"

	"github.com/ghdehrl12345/blockCore/db"
)

const StateBucket = "state"

type Account struct {
	Balance int
	Nonce   int
}

type StateDB struct {
	db *db.DB
}

func NewStateDB(database *db.DB) *StateDB {
	database.Conn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(StateBucket))
		return err
	})
	return &StateDB{db: database}
}

func (s *StateDB) GetAccount(address []byte) *Account {
	var account Account
	data := s.getFromBucket(address)
	if data == nil {
		return &Account{Balance: 0, Nonce: 0}
	}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&account)
	return &account
}

func (s *StateDB) UpdateAccount(address []byte, account *Account) {
	var encoded bytes.Buffer
	encoder := gob.NewEncoder(&encoded)
	encoder.Encode(account)
	s.saveToBucket(address, encoded.Bytes())
}

func (s *StateDB) getFromBucket(key []byte) []byte {
	var result []byte
	s.db.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(StateBucket))
		if b != nil {
			result = b.Get(key)
		}
		return nil
	})
	return result
}

func (s *StateDB) saveToBucket(key, value []byte) {
	s.db.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(StateBucket))
		return b.Put(key, value)
	})
}
