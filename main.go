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

	err = eth.Init([]string{"0x07FFDc22cb928f3dadfa5B8473F3A841A1E4E8aC", "0x4fcbd61Ef5E2f919195559c4aF864BBD0A508F0F"})

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
