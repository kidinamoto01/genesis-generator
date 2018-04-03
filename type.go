package main


import (
	"encoding/json"
	cmn "github.com/tendermint/tmlibs/common"
	"io/ioutil"
	cryptoKeys "github.com/tendermint/go-crypto/keys"
	"fmt"
	"github.com/tendermint/go-crypto"
	"time"
	"sync"
	"github.com/cosmos/cosmos-sdk/client"
	dbm "github.com/tendermint/tmlibs/db"
	)

type Coin struct {
	Denom string `json:"denom"`
	Amount int `json:"amount"`
}

type Account struct{
	Address crypto.Address  `json:"address"`
	Coins []Coin `json:"coins"`

}

type AccountList struct {
	Accounts []Account `json:"accounts"`
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
	AppOptions      AppOptions    `json:"app_state,omitempty"`
}


type AppOptions struct{

	Accounts []Account `json:"accounts"`
  //  Options []interface{}   `json:"plugin_options"`
}

type Signer interface {
	Sign(msg []byte) (crypto.Signature, error)
}

type PrivValidator struct {
	Address       crypto.Address          `json:"address"`
	PubKey        crypto.PubKey    `json:"pub_key"`
	LastHeight    int64            `json:"last_height"`
	LastRound     int              `json:"last_round"`
	LastStep      int8             `json:"last_step"`
	LastSignature crypto.Signature `json:"last_signature,omitempty"` // so we dont lose signatures
	LastSignBytes cmn.HexBytes     `json:"last_signbytes,omitempty"` // so we dont lose signatures

	// PrivKey should be empty if a Signer other than the default is being used.
	PrivKey crypto.PrivKey `json:"priv_key"`
	Signer  `json:"-"`

	// For persistence.
	// Overloaded for testing.
	filePath string
	mtx      sync.Mutex
}

var (
	manager cryptoKeys.Keybase
)


func GenerateGenesis(id string,acclist AccountList) GenesisDoc{

	genDoc := GenesisDoc{}

	genDoc.AppHash = cmn.HexBytes{}

	genDoc.GenesisTime = time.Now()

	genDoc.ChainID = id


	genDoc.Validators = generateValidators("test1")

	genDoc.AppOptions.Accounts = acclist.Accounts


	//genDoc.AppOptions.Options = generateOption(acclist.Accounts[0].Address)
	return genDoc

}

func generateValidators(name string)([]GenesisValidator){

	priVals := []GenesisValidator{}

	priVal :=GenesisValidator{}


	privKey := crypto.GenPrivKeyEd25519().Wrap()

	//priVal.PubKey = privKey.PubKey()

	//privValidator := types.GenPrivValidatorFS("")
	privValidator :=PrivValidator{
		Address:  privKey.PubKey().Address(),
		PubKey:   privKey.PubKey(),
		PrivKey:  privKey,
		LastStep: 0,
		Signer:   NewDefaultSigner(privKey),
		filePath: "",
	}


	exportContent(&privValidator,"./data/priv_validator.json")

	//types.GenesisDocFromFile("")

	priVal.PubKey = privValidator.PubKey
	priVal.Name = name
	//固定权重
	priVal.Power = 10

	priVals = append(priVals, priVal)


	return priVals


}

type Opt struct{
	App string `json:"app"`
	Addr   string `json:"addr"`

}

func generateOption(add crypto.Address) []interface{}{


	var result []interface{}
	issuer := &Opt{App:"sigs",Addr:add.String()}
	//issuerstr, err := json.Marshal(issuer)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//issuer := "{ app: sigs,addr: "+add.String()+"}"
	result = append(result,"coin/issuer")
	result = append(result,"stake/allowed_bond_denom")
	result = append(result,issuer)

	result = append(result,"fermion")
	return result
}


func exportContent(input interface{},path string){

	privValidatorJSONBytes, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		panic(err)
	}
	// write the whole body at once
	err = ioutil.WriteFile(path, privValidatorJSONBytes, 0644)
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
	//kb, err := keys.GetKeyBase() // dbm.NewMemDB()) // :(
	//kb := client.MockKeyBase()
	kb,err := GetKeyBase()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//var info cryptokeys.Info
	algo := cryptoKeys.AlgoEd25519

	fmt.Println("algo: ",algo)
	info, seed, err := kb.Create(name, pass, algo)

	//fmt.Println(seed)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fermion := Coin{Denom:"fermion",Amount:100}
	coins :=[]Coin{fermion}
	var add crypto.Address
	add = info.Address()
	account := Account{Address:add,Coins:coins}
	//fmt.Println(account)

	return account,seed



}


// NewDefaultSigner returns an instance of DefaultSigner.
func NewDefaultSigner(priv crypto.PrivKey) *DefaultSigner {
	return &DefaultSigner{
		PrivKey: priv,
	}
}

type DefaultSigner struct {
	PrivKey crypto.PrivKey `json:"priv_key"`
}

// Sign implements Signer. It signs the byte slice with a private key.
func (ds *DefaultSigner) Sign(msg []byte) (crypto.Signature, error) {
	return ds.PrivKey.Sign(msg), nil
}


// GetKeyManager initializes a key manager based on the configuration
func GetKeyBase() (cryptoKeys.Keybase, error){
	rootDir := ""
	KeyDBName :="test"
	db, err := dbm.NewGoLevelDB(KeyDBName, rootDir)
	if err != nil {
		return nil, err
	}
	keybase := client.GetKeyBase(db)

	return keybase,nil
}