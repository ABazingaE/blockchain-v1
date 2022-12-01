package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

type Transaction struct {
	TXID      []byte     //id
	TXInputs  []TXInput  //可以有多个输入
	TXOutputs []TXOutput //多个输出
	TimeStamp uint64     //创建交易的时间戳
}

type TXInput struct {
	Txid      []byte //引用的UTXO所在的交易id
	Index     int64  //引用的output在output数组中的索引
	ScriptSig string //付款人对当前交易的签名（后续更新签名+公钥，暂时简化）
}

type TXOutput struct {
	ScriptPubk string //收款人的公钥哈希（从地址推出，先理解为地址）
	Value      uint64 //转账金额
}

func (tx *Transaction) setHash() error {
	//对tx做gob编码得到字节流，做sha256，赋值id
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(tx)
	if err != nil {
		fmt.Println("encode error:", err)
	}

	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
	return nil
}

/**
 * @description: 生成挖矿交易
 * @param {string} miner
 * @return {*}
 */
//暂时写死挖矿奖励
var reward = 12.5

func NewCoinbaseTx(miner string, data string) *Transaction {
	tx := Transaction{}
	//特点：没有输入，只有一个输出，得到挖矿奖励
	//挖矿交易需要能够识别出来，没有input，所以不需要签名
	//挖矿交易不需要签名，签名字段可以写任意值，只有矿工有权利写

	input := TXInput{
		Txid:      nil,
		Index:     -1,
		ScriptSig: data,
	}

	output := TXOutput{
		Value:      uint64(reward),
		ScriptPubk: miner,
	}

	timeStamp := time.Now().Unix()

	tx.TXID = nil
	tx.TXInputs = []TXInput{input}
	tx.TXOutputs = []TXOutput{output}
	tx.TimeStamp = uint64(timeStamp)

	//哈希运算，填写id
	tx.setHash()
	return &tx
}
