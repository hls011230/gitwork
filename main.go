package main

import (
	"A11Smile/db"
	"A11Smile/eth"
	"A11Smile/service"
)

// 程序的入口
func main() {

	// 初始化数据库
	err := db.Init()
	if err != nil {
		panic(err)
	}

	// 初始化区块链(传入合约地址)

	err = eth.Init([]string{"0xE1585130e00Ce35Cdd3b3e17d4879FEe2485934E", "0x94E11665F0F457e4068e9E0388800dE27B5fF239"})

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
