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

	err = eth.Init([]string{"0xFFBf214564d6a90b95F3390afc0899B12862a580", "0x30c96a4A2180D3B39385cad5b28e68C065DB6d05"})

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
