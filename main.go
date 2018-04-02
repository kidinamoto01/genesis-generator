package main

//gaia init 后生成config.toml、genesis.json、priv_validator.json这三个文件
import (
	"encoding/json"
	"fmt"
)

const(
	password = "1234567890"
)


func main(){

	//生成n个余额为100的账户


	num := 1
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
		fmt.Println(account)
		fmt.Println(seed)
		accts = append(accts,account)
		seedList = append(seedList,seed)
	}

	accList.Accounts =accts

	result := GenerateGenesis("test",accList)
	resultJSONBytes,err := json.Marshal(&result)
	if err != nil{

	}else{
		fmt.Println(string(resultJSONBytes))
	}
	exportContent(&result,"./data/genesis.json")

	//fmt.Printf("%+v\n", result)

}


func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


