package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func User_ETHforAs(uid int, as *model.PostETHforAS) error {
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w, "id = ?", uid)

	// 在合约中存入用户病历信息
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(w.BlockAddress))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(w.PrivateKey)
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
	auth.Value = big.NewInt(int64(as.Quantity * int(math.Pow10(18))))

	noncex, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(model.Deployer.Address))
	if err != nil {

		return err
	}

	privateKeyx, err := crypto.HexToECDSA(model.Deployer.PrivateKey)
	if err != nil {

		return err
	}

	authx, err := bind.NewKeyedTransactorWithChainID(privateKeyx, eth.ChainID)
	if err != nil {
		return err
	}

	authx.GasPrice = eth.GasPrice
	authx.GasLimit = uint64(6000000)
	authx.Nonce = big.NewInt(int64(noncex))

	_, err = eth.AS.Transfer(authx, common.HexToAddress(w.BlockAddress), big.NewInt(int64(2*as.Quantity)))

	_, err = eth.AS.EthGetAs(auth, common.HexToAddress(model.Deployer.Address))
	if err != nil {
		return err
	}

	return err

}
