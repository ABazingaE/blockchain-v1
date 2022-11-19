package main

import "math/big"

type ProofOfWork struct {
	//区块
	block *Block

	//目标值,使用该类型是因为提供了多种方法：比较、将哈希值转换为big.int
	target big.Int
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
	tempStr := "0000100000000000000000000000000000000000000000000000000000000000"

	tempBigInt := new(big.Int)

	tempBigInt.SetString(tempStr, 16)

	pow.target = *tempBigInt

	return &pow

}
