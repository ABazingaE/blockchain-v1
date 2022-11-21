package main

import (
	"fmt"
	"os"
)

type CLI struct {
}

const Usage = `

	Usage:
		./blockchain create  : 创建区块链
		./blockchain addblock <data> : 添加新区块
		./blockchain print			 : 遍历打印区块链
`

func (cli *CLI) Run() {
	command := os.Args
	if len(command) < 2 {
		fmt.Printf("输入参数不合法")
		fmt.Println(Usage)
		return
	}
	switch command[1] {
	case "create":
		fmt.Printf("创建区块链")
		cli.createBlockChain()

	case "addblock":
		fmt.Printf("添加区块")
		data := command[2]
		cli.addBlock(data)
	case "print":
		fmt.Printf("打印区块链")
		cli.print()
	default:
		fmt.Printf("命令无效")
		fmt.Println(Usage)
	}

}
