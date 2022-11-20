package main

import "fmt"

func main() {
	// bc := NewBlockChain()

	// bc.AddBlock("第二个区块")
	// bc.AddBlock("顺便熟悉一下Go")

	// for i, block := range bc.Blocks {
	// 	fmt.Printf("\n++++++++++Height+++++++++++++++: %d\n", i)
	// 	fmt.Printf("Version: %d\n", block.Version)
	// 	fmt.Printf("PrevHash: %x\n", block.PrevHash)
	// 	fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
	// 	fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
	// 	fmt.Printf("Bits: %d\n", block.Bits)
	// 	fmt.Printf("Nonce: %d\n", block.Nonce)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Data: %s\n", string(block.Data))

	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("is valid: %v\n", pow.IsValid())

	// }

	CreateBlockChain()

	bc, err := GetBlockChainInstance()

	if err != nil {
		fmt.Printf("error:", err)
	}
	defer bc.db.Close()

	bc.AddBlock("test add function")

	bc.AddBlock("have a cold today")

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
