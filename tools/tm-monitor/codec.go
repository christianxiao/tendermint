package main

import (
	ctypes "github.com/christianxiao/tendermint/rpc/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
