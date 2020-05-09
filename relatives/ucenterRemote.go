package relatives

import (
	"context"
	"demo/pkg/transports/rpc"
	"demo/ucenter/provider/param"
	"github.com/gutrse3321/aki/persit/remote"
	"github.com/pkg/errors"
	"github.com/smallnest/rpcx/client"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:06
 * @Title:
 * --- --- ---
 * @Desc:
 */

type UCenterRemote struct {
	client  *rpc.Client
	xClient client.XClient
}

func NewUCenterRemote(client *rpc.Client) (*UCenterRemote, error) {
	xClient, err := client.Connect("UCenterProvider")
	if err != nil {
		return nil, errors.Wrap(err, "connected ucenter-provider error")
	}
	return &UCenterRemote{client, xClient}, nil
}

func (u *UCenterRemote) GetUserBaseInfo(args *param.GetUserBaseInfoArgs) (*remote.Remote, error) {
	codeRemote := &remote.Remote{}
	err := u.xClient.Call(context.Background(), "GetUserBaseInfo", args, codeRemote)
	if err != nil {
		return nil, err
	}

	return codeRemote, nil
}
