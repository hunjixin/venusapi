package main

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
)

func CheckGetBaseInfo(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do GasEstimateFeeCap %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		for i := 0; i < 10; i++ {
			lotusBaseInfo, err := nodeFullNode.MinerGetBaseInfo(ctx, addr, headTs.Height(), headTs.Parents())
			if err != nil {
				log.Errorf("fail to do GasEstimateFeeCap %v", err)
				return
			}

			venusBaseInfo, err := venusFullNode.MinerGetBaseInfo(ctx, addr, headTs.Height(), headTs.Parents())
			if err != nil {
				log.Errorf("fail to do GasEstimateFeeCap %v", err)
				return
			}
			if !checkRaw(lotusBaseInfo, venusBaseInfo) {
				log.Errorf("GasEstimateFeeCap fail venus %v, lotus %v", lotusBaseInfo, venusBaseInfo)
				return
			}
		}
	}

	log.Infof("success to check api MinerGetBaseInfo")
}
