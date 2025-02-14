package core

import (
	ctypes "github.com/christianxiao/tendermint/rpc/core/types"
	rpctypes "github.com/christianxiao/tendermint/rpc/lib/types"
)

// Get node health. Returns empty result (200 OK) on success, no response - in
// case of an error.
//
// ```shell
// curl 'localhost:26657/health'
// ```
//
// ```go
// client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
// err := client.Start()
// if err != nil {
//   // handle error
// }
// defer client.Stop()
// result, err := client.Health()
// ```
//
// > The above command returns JSON structured like this:
//
// ```json
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
// ```
func Health(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
