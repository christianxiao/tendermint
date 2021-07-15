package privval

import (
	cryptoAmino "github.com/christianxiao/tendermint/crypto/encoding/amino"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterRemoteSignerMsg(cdc)
}
