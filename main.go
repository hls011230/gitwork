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

	err = eth.Init([]string{"0xe1f8d8679DDA6f820b202CcF3b418c2650A3503A", "0xD8cA4CCbeA2D55605555aaa010b6e8550C1aE028"})

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
