package provider

import (
	"context"
	"demo/ucenter/provider/param"
	"demo/ucenter/service"
	"github.com/gutrse3321/aki-persit/remote"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 17:50
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Provider struct {
}

func (p *Provider) GetUserBaseInfo(ctx context.Context, args *param.GetUserBaseInfoArgs, codeRemote *remote.Remote) error {
	userBaseInfoDto := service.GetUserBaseInfo()
	remote.Init(codeRemote, userBaseInfoDto)
	return nil
}
