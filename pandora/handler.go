package main

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"micah.wiki/pandora/kitex_gen/dragon"
)

// PandoraImpl implements the last service interface defined in the IDL.
type PandoraImpl struct{}

// DragonSay implements the PandoraImpl interface.
func (s *PandoraImpl) DragonSay(ctx context.Context, req *dragon.DragonSayRequest) (resp *dragon.DragonSayResponse, err error) {
	if len(req.GetMessage()) == 0 {
		logger.CtxErrorf(ctx, "req is empty")
	}
	resp.SetMessage("Hello")
	return resp, nil
}
