package main

type Transaction struct {
	TXID      []byte     //id
	TXInputs  []TXInput  //可以有多个输入
	TXOutputs []TXOutput //多个输出
	TimeStamp uint64     //创建交易的时间戳
}

type TXInput struct {
	Txid      []byte //引用的UTXO所在的交易id
	Index     uint64 //引用的output在output数组中的索引
	ScriptSig string //付款人对当前交易的签名（后续更新签名+公钥，暂时简化）
}

type TXOutput struct {
	ScriptPubk string //收款人的公钥哈希（从地址推出，先理解为地址）
	Value      uint64 //转账金额
}
