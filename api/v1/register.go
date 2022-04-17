package v1

import (
	"A11Smile/db/model"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Register(user *model.User) error {
	//key := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	//passwd := user.Passwd
	//
	////创建一个钱包用户
	//create_account, err := key.NewAccount(passwd)
	//if err != nil {
	//	return err
	//}
	//log.Fatal(create_account.Address)
	//log.Fatal(create_account.URL)
	CreateAccount()
	fmt.Println("success")
	return nil
}
func CreateAccount() (string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err)
		return "", nil
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	log.Println("address: ", address)

	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("privateKey: ", privateKey)
	return address, nil
}
