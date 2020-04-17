package rpc

import (
	"github.com/google/wire"
	"github.com/smallnest/rpcx/client"
	"github.com/spf13/viper"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/17 16:20
 * @Title:
 * --- --- ---
 * @Desc:
 */
type ClientOptions struct {
	BasePath    string
	EtcdAddress []string
}

func NewClientOptions(v *viper.Viper) (*ServerOptions, error) {
	opt := &ServerOptions{}

	if err := v.UnmarshalKey("rpc", opt); err != nil {
		return nil, err
	}
	return opt, nil
}

type Client struct {
	opt *ClientOptions
}

func NewClient(opt *ClientOptions) (*Client, error) {
	return &Client{opt: opt}, nil
}

func (c *Client) Connect(service string) (client.XClient, error) {
	d := client.NewEtcdV3Discovery(c.opt.BasePath, service, c.opt.EtcdAddress, nil)
	xClient := client.NewXClient(service, client.Failbackup, client.WeightedRoundRobin, d, client.DefaultOption)
	return xClient, nil
}

var ClientProviderSet = wire.NewSet(NewClientOptions, NewClient)
