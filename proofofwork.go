package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	//区块
	block *Block

	//目标值,使用该类型是因为提供了多种方法：比较、将哈希值转换为big.int
	target *big.Int
}

/**
 * @description:新生成工作量证明,block由用户提供，target由系统提供
 * @param {*Block} block
 * @return {*ProofOfWork}
 */
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	//难度暂时写死，后续更新推导方式
	tempStr := "0001000000000000000000000000000000000000000000000000000000000000"

	tempBigInt := new(big.Int)

	tempBigInt.SetString(tempStr, 16)

	pow.target = tempBigInt

	return &pow

}

/**
 * @description: 挖矿
 * @return {*}
 */
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	var hash [32]byte
	fmt.Println("开始挖矿....")
	for {
		fmt.Printf("%x\r", hash[:])
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)

		tempBigInt := new(big.Int)

		tempBigInt.SetBytes(hash[:])

		if tempBigInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功!hash:%x,nonce:%d\n", hash[:], nonce)
			break
		} else {
			nonce++
		}
	}
	return hash[:], nonce

}

/**
 * @description: 准备用于计算哈希的数据：拼接区块头的各项数据+nonce
 * @param {uint64} nonce
 * @return {*}
 */
func (pow *ProofOfWork) PrepareData(nonce uint64) []byte {
	b := pow.block
	//data是block各个字段的拼接
	temp := [][]byte{
		uintToByte(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(nonce),
		//b.Hash,
		b.Data,
	}
	//注意使用bytes.join将二维切片转化为一维切片
	data := bytes.Join(temp, []byte{})
	return data
}

/**
 * @description: 校验有效性
 * @return {*}
 */
func (pow *ProofOfWork) IsValid() bool {
	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	tempInt := new(big.Int)
	tempInt.SetBytes(hash[:])
	return tempInt.Cmp(pow.target) == -1
}
