package main

import "fmt"

func (cli *CLI) addBlock(data string) {
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("添加失败:", err)
	}
	err = bc.AddBlock(data)

	if err != nil {
		fmt.Println("添加失败:", err)
	}

	fmt.Println("添加成功")
}

func (cli *CLI) createBlockChain() {
	err := CreateBlockChain()
	if err != nil {
		fmt.Println("创建区块链失败：", err)
		return
	}
	fmt.Println("创建区块链成功")
}

func (cli *CLI) print() {
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("打印失败:", err)
	}
	it := bc.NewIterator()
	for {
		block := it.Next()

		fmt.Printf("\n++++++++++Height++++++++++++++++++\n")
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Bits: %d\n", block.Bits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", string(block.Data))

		pow := NewProofOfWork(block)
		fmt.Printf("is valid: %v\n", pow.IsValid())

		if block.PrevHash == nil {
			break
		}
	}
}
