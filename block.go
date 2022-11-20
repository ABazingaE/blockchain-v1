package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

	//计算hash
	// b.setHash()
	pow := NewProofOfWork(&b)
	hash, nonce := pow.Run()
	b.Hash = hash
	b.Nonce = nonce
	return &b
}

func (b *Block) Serialize() []byte {
	//用于指定容器，告诉编码器往哪里放
	var buffer bytes.Buffer

	//create encoder
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(b)

	if err != nil {
		fmt.Printf("encode error:", err)
		return nil
	}

	return buffer.Bytes()
}

func Deserialize(src []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(src))

	err := decoder.Decode(&block)

	if err != nil {
		fmt.Printf("decode error:", err)
		return nil
	}

	return &block
}
