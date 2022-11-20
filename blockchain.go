package main

import (
	"errors"

	"github.com/boltdb/bolt"
)

/*
	定义区块链结构（数组模拟）
*/
type BlockChain struct {
	//用于存储数据
	db *bolt.DB

	//最后一个区块的哈希值
	tail []byte
}

/*
	创建区块链
*/
//创世语
const genesisInfo = "Hello,Web3!"
const blockchainDbFile = "blockchain"
const blockBucket = "blockBucket"
const lastblockKey = "lastBlockKey"

func CreateBlockChain() error {
	db, err := bolt.Open(blockchainDbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	//start to create
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		//若没有bucket，则创建，添加创世块
		if bucket == nil {
			//create bucket
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				return err
			}

			//create genisisBlock
			genisisBlock := NewBlock(genesisInfo, nil)

			//put into the bucket,key is the hash of the block, value is the block
			bucket.Put(genisisBlock.Hash, genisisBlock.Serialize())
			bucket.Put([]byte(lastblockKey), genisisBlock.Hash)
		}

		return nil

	})

	return err

}

/**
 * @description: 获取已有的blockchain实例
 * @return {*}
 */
func GetBlockChainInstance() (*BlockChain, error) {

	var lastHash []byte

	db, err := bolt.Open(blockchainDbFile, 0400, nil)
	if err != nil {
		return nil, err
	}

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))

		//若没有bucket，则创建，添加创世块
		if bucket == nil {
			return errors.New("bucket 为空")
		}

		lastHash = bucket.Get([]byte(lastblockKey))

		return nil

	})

	blockchain := BlockChain{
		db:   db,
		tail: lastHash,
	}

	return &blockchain, nil

}

/*
	添加区块
*/
func (bc *BlockChain) AddBlock(data string) {
	//拿到最后一个区块的哈希值作为新区块的前置哈希
	// lastBlock := bc.Blocks[len(bc.Blocks)-1]
	// prevHash := lastBlock.Hash
	// newBlock := NewBlock(data, prevHash)
	// bc.Blocks = append(bc.Blocks, newBlock)
}
