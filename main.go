package main

//gaia init 后生成config.toml、genesis.json、priv_validator.json这三个文件
import (
	//"encoding/json"
	"fmt"
	//"github.com/tendermint/go-crypto"

)

const(
	password = "1234567890"
)




func main(){

	//生成n个余额为100的账户


	num := 10
	var accts []Account
	var accList AccountList
	var seedList []string

	//bl := AccountList{
	//	Accounts: []Account{
	//		Account{},
	//	},
	//}
	//data, _ := json.MarshalIndent(bl, "", "  ");
	//fmt.Println(string(data));

	var (
		account Account
		seed string
	)
	for i:=0;i<num;i++{

		account,seed = GenerateAccount("abc",password)
		//fmt.Println(account)
		//fmt.Println(seed)
		accts = append(accts,account)
		seedList = append(seedList,seed)
	}

	accList.Accounts =accts

	GenerateGenesis("test",accList)

}


func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


