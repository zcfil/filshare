package config

import (
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/spf13/viper"
)

var FullAPI *YungoRpc

type YungoRpc struct {
	Rpc    api.FullNode
	Closer jsonrpc.ClientCloser
}

type Lotus struct {
	Host  string
	Token string
}

func InitLotus(cfg *viper.Viper) *Lotus {
	return &Lotus{
		Host:  cfg.GetString("host"),
		Token: cfg.GetString("token"),
	}
}

var LotusConfig = new(Lotus)
