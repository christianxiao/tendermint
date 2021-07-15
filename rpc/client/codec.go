package client

import (
	"github.com/christianxiao/tendermint/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterEvidences(cdc)
}
