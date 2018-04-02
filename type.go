package main


import (
	"encoding/json"
	//"github.com/tendermint/go-crypto"
	"github.com/tendermint/tendermint/types"
	cmn "github.com/tendermint/tmlibs/common"
	"io/ioutil"
	//	"time"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/tendermint/go-crypto"
	"time"
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

// GenesisValidator is an initial validator.
type GenesisValidator struct {
	PubKey crypto.PubKey `json:"pub_key"`
	Power  int64         `json:"power"`
	Name   string        `json:"name"`
}

type GenesisDoc struct {
	GenesisTime     time.Time          `json:"genesis_time"`
	ChainID         string             `json:"chain_id"`
	Validators      []GenesisValidator `json:"validators"`
	AppHash         cmn.HexBytes       `json:"app_hash"`
	AppOptions      AppOptions    `json:"app_options,omitempty"`
}


type AppOptions struct{

	Accounts []Account `json:"accounts"`
    Options []string   `json:"plugin_options"`
}


func GenerateGenesis(id string,acclist AccountList){

	genDoc := GenesisDoc{}

	genDoc.GenesisTime = time.Now()

	genDoc.ChainID = id


	genDoc.Validators = generateValidators("test1")

	//genDoc.AppOptions =


}

func generateValidators(name string)([]GenesisValidator){

	priVals := []GenesisValidator{}

	priVal :=GenesisValidator{}


	//privKey := crypto.GenPrivKeyEd25519().Wrap()
	//
	//
	//priVal.PubKey = privKey.PubKey()
	privValidator := types.GenPrivValidatorFS("")

	exportVaidatorKey(privValidator)

	types.GenesisDocFromFile("")

	priVal.Name = name
	//固定权重
	priVal.Power = 10

	priVals[0] = priVal


	return priVals


}

func exportVaidatorKey(input *types.PrivValidatorFS){

	privValidatorJSONBytes, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		panic(err)
	}
	// write the whole body at once
	err = ioutil.WriteFile("output.txt", privValidatorJSONBytes, 0644)
	if err != nil {
		panic(err)
	}

}

func GenerateOption(){

	//
	//acc, seed := GenerateAccount("alice",password)
	//
	//var accList []Account


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