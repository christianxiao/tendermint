package main

import (
	"fmt"
	"os"

	cryptoAmino "github.com/christianxiao/tendermint/crypto/encoding/amino"
	amino "github.com/tendermint/go-amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
