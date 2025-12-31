package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

// 해시 앞에 0이 몇 개 있어야 하는지
const TargetBits = 16

// 작업증명 구조체
type ProofOfWork struct {
	block  *Block
	target *big.Int // 해시가 이 값보다 작아야 함
}

// 작업증명 생성 함수
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-TargetBits)
	return &ProofOfWork{block: b, target: target}
}

// 해시 계산을 위한 데이터 준비
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	timestamp := make([]byte, 8)
	binary.BigEndian.PutUint64(timestamp, uint64(pow.block.Timestamp))

	targetBits := make([]byte, 8)
	binary.BigEndian.PutUint64(targetBits, uint64(TargetBits))

	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, uint64(nonce))

	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			timestamp,
			targetBits,
			nonceBytes,
		},
		[]byte{},
	)

	return data
}

// 채굴 시작, Nonce를 찾을 떄까지 반복
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		// 해시값이 target보다 작으면 성공
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

// 블록의 PoW가 유효한지 검증
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
