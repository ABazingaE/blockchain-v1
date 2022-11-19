package main

/*
	定义区块链结构（数组模拟）
*/
type BlockChain struct {
	Blocks []*Block
}

/*
	创建区块链
	1.创建BlockChain
	2.添加一个创世区块
*/
//创世语
const genesisInfo = "Hello,Web3!"

func NewBlockChain() *BlockChain {

	genesisBlock := NewBlock(genesisInfo, nil)
	bc := BlockChain{
		Blocks: []*Block{genesisBlock},
	}
	return &bc
}

/*
	添加区块
*/
func (bc *BlockChain) AddBlock(data string) {
	//拿到最后一个区块的哈希值作为新区块的前置哈希
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := lastBlock.Hash
	newBlock := NewBlock(data, prevHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
