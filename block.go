package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

/* 定义区块
1.定义基础字段：前哈希，哈希，数据
2.第二阶段补充字段：版本号、时间戳、梅克尔根等等
*/

type Block struct {

	//版本号：表示本区块遵循的验证规则
	Version uint64

	//前一区块的哈希
	PrevHash []byte

	//梅克尔根，交易数据的根哈希
	MerkleRoot []byte

	//时间戳
	TimeStamp uint64

	//难度值，由系统提供，用于计算出一个哈希值
	Bits uint64

	//Nonce，挖矿所求的随机数
	Nonce uint64

	//该区块哈希，本应另外储存，方便起见一起写入
	Hash []byte

	//区块体数据
	Data []byte
}

/*
	创建区块
	Input: 数据，前哈希
	Output：Block
*/

func NewBlock(data string, prevHash []byte) *Block {

	b := Block{
		Version:    0,
		PrevHash:   prevHash,
		MerkleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       0,
		Nonce:      0,
		Data:       []byte(data),
		Hash:       nil,
	}
	b.setHash()
	return &b
}

/*
	计算Hash
*/
func (b *Block) setHash() {

	//data是block各个字段的拼接

	temp := [][]byte{
		uintToByte(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(b.Nonce),
		b.Hash,
		b.Data,
	}
	//注意使用bytes.join将二维切片转化为一维切片
	data := bytes.Join(temp, []byte{})
	hash := sha256.Sum256(data)
	//此处或许是将数组转化为切片？
	b.Hash = hash[:]
}
