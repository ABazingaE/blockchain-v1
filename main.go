package main

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

}
