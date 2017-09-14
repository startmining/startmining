//go:generate abigen --sol ./contracts/WrapperOraclize.sol --pkg gocontract  --out ./gocontract/gocontract.go
package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"io/ioutil"
	"strings"
	"log"

	"AlexStudio_ICO_2/gocontract"
)

const (
	KEYSTORE_DIR = `/Library/Ethereum/rinkeby/keystore`
	SIGN_PASSPHRASE = `dbnfkbq199529`
)

type User struct{
	Auth *bind.TransactOpts
	Conn *ethclient.Client
	Sim *backends.SimulatedBackend
	Nonce uint64
	JsonKey []byte
}

func CreateUser(key, pass string) *User {

	auth , err := bind.NewTransactor(strings.NewReader(key),pass)
	if err != nil {
		log.Fatal(err.Error())
	}

	sim, err := ethclient.Dial("/Users/user/.rinkeby/geth.ipc")
	if err != nil {
		log.Fatal(err.Error())
	}


	return &User{Auth:auth,Conn:sim}
}

func RealUsersInit() []*User {
	users := []*User{}

	// Init a keystore
	ks := keystore.NewKeyStore(
		KEYSTORE_DIR,
		keystore.LightScryptN,
		keystore.LightScryptP,
	)

	accounts := ks.Accounts()

	for _, v := range accounts{
		// Unlock the signing account
		errUnlock := ks.Unlock(v, SIGN_PASSPHRASE)
		if errUnlock != nil {
			fmt.Println("account unlock error:")
			panic(errUnlock.Error())
		}
	}




	for _, v := range accounts{
		keyJsonFile := v.URL.Path

		// Open the account key file
		keyJson, readErr := ioutil.ReadFile(keyJsonFile)
		if readErr != nil {
			fmt.Println("key json read error:")
			panic(readErr)
		}
		// Get the private key
		keyWrapper, keyErr := keystore.DecryptKey(keyJson, SIGN_PASSPHRASE)
		if keyErr != nil {
			fmt.Println("key decrypt error:")
			panic(keyErr)
		}
		k:= keyWrapper.PrivateKey
		hesh := common.BigToHash(k.D)
		fmt.Printf("key extracted: addr=%s, key=%s\n", keyWrapper.Address.String(), hesh.String())

		u := CreateUser(string(keyJson),SIGN_PASSPHRASE)
		users = append(users,u)
	}
	return users
}

// deploy
func (u *User)Deploy() (string,string, error) {

	//u.Auth.GasLimit = big.NewInt(3000000)

	// deploy contracts
	addr, tr, _, err := gocontract.DeployWrapperOraclize(u.Auth, u.Conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	// interact with contract
	fmt.Printf("Contract deployed to %s\n", addr.String())
	return addr.String(), tr.String(), nil
}

func main() {
	users := RealUsersInit()
	owner := users[0]
	addr, tr, err := owner.Deploy()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(addr)
	log.Print(tr)
}
