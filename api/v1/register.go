package v1

import (
	"A11Smile/db/model"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Register(user *model.User) error {
	key := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	passwd := user.Passwd

	//创建一个钱包用户
	create_account, err := key.NewAccount(passwd)
	if err != nil {
		return err
	}
	log.Fatal(create_account.Address)
	log.Fatal(create_account.URL)
	return nil
}
