package main

//gaia init 后生成config.toml、genesis.json、priv_validator.json这三个文件
import (
	"encoding/json"
	"fmt"
)

type Coin struct {
	Denom string `json:"denom"`;
	Amount int `json:"amount"`
}

type Account struct{
	Address string `json:"address"`;
	Coins []Coin `json:"coins"`

}

type AccountList struct {
	Accounts []Account `json:"accounts"`;
}

func main(){

	//生成n个余额为100的账户

}
