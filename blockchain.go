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
func (bc *BlockChain) AddBlock(data string) error {
	lastblockHash := bc.tail
	db := bc.db

	err := db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockBucket))

		if bucket == nil {
			return errors.New("bucket must not be empty when adding a block")
		}

		//1. create new block
		block := NewBlock(data, lastblockHash)

		//2.insert into db
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(lastblockKey), block.Hash)

		//3.update last hash in memory
		bc.tail = block.Hash

		return nil

	})

	return err
}

//++++++++++++++++++++++迭代器相关+++++++++++++++++++++++++++
type Iterator struct {
	db          *bolt.DB
	currentHash []byte
}

func (bc *BlockChain) NewIterator() *Iterator {
	it := Iterator{
		db:          bc.db,
		currentHash: bc.tail,
	}
	return &it
}

func (it *Iterator) Next() *Block {
	db := it.db
	var block *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			return errors.New("找不到Bucket")
		}
		data := bucket.Get(it.currentHash)
		block = Deserialize(data)
		return nil
	})
	it.currentHash = block.PrevHash
	return block
}
