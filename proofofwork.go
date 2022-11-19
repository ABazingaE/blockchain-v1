package main

import "math/big"

type ProofOfWork struct {
	//区块
	block *Block

	//目标值,使用该类型是因为提供了多种方法：比较、将哈希值转换为big.int
	target big.Int
}
