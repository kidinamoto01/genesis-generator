package main

//gaia init 后生成config.toml、genesis.json、priv_validator.json这三个文件
import (
	//"encoding/json"
	"fmt"

	//"github.com/tendermint/go-crypto"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/tendermint/go-crypto"
)

const(
	password = "1234567890"
)

type Coin struct {
	Denom string `json:"denom"`;
	Amount int `json:"amount"`
}

type Account struct{
	Address crypto.Address  `json:"address"`;
	Coins []Coin `json:"coins"`

}

type AccountList struct {
	Accounts []Account `json:"accounts"`;
}

func main(){

	//生成n个余额为100的账户


	//num := 10
	//accList := []Account{}
	//
	//
	//
	//bl := AccountList{
	//	Accounts: []Account{
	//		Account{},
	//	},
	//}
	//data, _ := json.MarshalIndent(bl, "", "  ");
	//fmt.Println(string(data));

	account,seed := GenerateAccount("abc",password)
	fmt.Println(account)
	fmt.Println(seed)

}



func GenerateAccount(name string,pass string) (Account,string){

	kb, err := keys.GetKeyBase() // dbm.NewMemDB()) // :(
	if err != nil {
		fmt.Println(err)
		panic(err)
	}


	//var info cryptokeys.Info
	info, seed, err := kb.Create(name, pass, "ed25519")

	fmt.Println(seed)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fermion := Coin{Denom:"fermion",Amount:100}
	coins :=[]Coin{fermion}
	var add crypto.Address
	add = info.Address()
	account := Account{Address:add,Coins:coins}
	fmt.Println(account)

	return account,seed



}