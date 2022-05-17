package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Ganiner_ETHforAs(gid int, As *model.PostETHforAS) error {
	DB := db.Get()

	// 在合约中存入用户病历信息
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(model.Deployer.Address))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(model.Deployer.PrivateKey)
	if err != nil {

		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))

	var x model.Wallet
	DB.Table("gainers").First(&x, "id = ?", gid)

	noncex, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(x.BlockAddress))
	if err != nil {

		return err
	}

	privateKeyx, err := crypto.HexToECDSA(x.PrivateKey)
	if err != nil {

		return err
	}

	authx, err := bind.NewKeyedTransactorWithChainID(privateKeyx, eth.ChainID)
	if err != nil {
		return err
	}

	value := strconv.Itoa(As.Quantity)

	valuef, err := strconv.ParseFloat(value, 64) //先转换为 float64

	if err != nil {
		log.Println("is not a.abi number")
	}

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}

	authx.GasPrice = eth.GasPrice
	authx.GasLimit = uint64(6000000)
	authx.Nonce = big.NewInt(int64(noncex))
	authx.Value = valueWei

	_, err = eth.AS.EthGetAs(authx, common.HexToAddress(model.Deployer.Address))
	_, err = eth.AS.Transfer(auth, common.HexToAddress(x.BlockAddress), big.NewInt(int64(As.Quantity*2)))
	if err != nil {
		return err
	}

	return err

}
