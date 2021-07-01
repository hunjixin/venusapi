package main

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

func CheckWalletBalance(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	add1, _ := address.NewFromString("f1ojyfm5btrqq63zquewexr4hecynvq6yjyk5xv6q")
	add2, _ := address.NewFromString("f01000")
	add3, _ := address.NewFromString("f3qfrxne7cg4ml45ufsaxqtul2c33kmlt4glq3b4zvha3msw4imkyi45iyhcpnqxt2iuaikjmmgx2xlr5myuxa")
	for _, addr := range []address.Address{add1, add2, add3} {
		venusBalance, err := venusFullNode.WalletBalance(ctx, addr)
		if err != nil {
			log.Errorf("fial to do WalletBalance %v", err)
			return
		}
		lotusBalance, err := nodeFullNode.WalletBalance(ctx, addr)
		if err != nil {
			log.Errorf("fial to do WalletBalance %v", err)
			return
		}

		if venusBalance.String() != lotusBalance.String() {
			log.Errorf("WalletBalance address %s fail venus %v, lotus %v", addr, venusBalance.String(), lotusBalance.String())
			return
		}
	}

	log.Infof("success to check api WalletBalance")
}

func CheckGasEstimateGasLimit(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do GasEstimateGasLimit %v", err)
		return
	}

	addr, _ := address.NewFromString("f1kc5ami5ejqept4ydpjme5ytzywzvrgfxxsjy57i")
	actor, err := nodeFullNode.StateGetActor(ctx, addr, headTs.Key())
	if err != nil {
		log.Errorf("fail to do GasEstimateGasLimit %v", err)
		return
	}
	msg := types.Message{
		Version:    0,
		To:         addr,
		From:       addr,
		Nonce:      actor.Nonce + 3,
		Value:      abi.NewTokenAmount(100),
		GasLimit:   0,
		GasFeeCap:  abi.TokenAmount{},
		GasPremium: abi.TokenAmount{},
		Method:     0,
		Params:     nil,
	}

	venuslimit, err := venusFullNode.GasEstimateGasLimit(ctx, &msg, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateGasLimit %v", err)
		return
	}
	lotusBalance, err := nodeFullNode.GasEstimateGasLimit(ctx, &msg, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateGasLimit %v", err)
		return
	}

	if venuslimit != lotusBalance {
		log.Errorf("GasEstimateGasLimit fail venus %v, lotus %v", venuslimit, lotusBalance)
		return
	}

	log.Infof("success to check api GasEstimateGasLimit")
}

func CheckGasEstimateGasPremium(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do GasEstimateGasPremium %v", err)
		return
	}

	addr, _ := address.NewFromString("f1kc5ami5ejqept4ydpjme5ytzywzvrgfxxsjy57i")

	venuslimit, err := venusFullNode.GasEstimateGasPremium(ctx, 5, addr, 1000000, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateGasPremium %v", err)
		return
	}
	lotusBalance, err := nodeFullNode.GasEstimateGasPremium(ctx, 5, addr, 1000000, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateGasPremium %v", err)
		return
	}

	if venuslimit != lotusBalance {
		log.Errorf("GasEstimateGasPremium fail venus %v, lotus %v", venuslimit, lotusBalance)
		return
	}

	log.Infof("success to check api GasEstimateGasPremium")
}

func CheckGasEstimateMessageGas(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do GasEstimateMessageGas %v", err)
		return
	}

	addr, _ := address.NewFromString("f1kc5ami5ejqept4ydpjme5ytzywzvrgfxxsjy57i")
	actor, err := nodeFullNode.StateGetActor(ctx, addr, headTs.Key())
	if err != nil {
		log.Errorf("fail to do GasEstimateMessageGas %v", err)
		return
	}
	msg := types.Message{
		Version:    0,
		To:         addr,
		From:       addr,
		Nonce:      actor.Nonce + 3,
		Value:      abi.NewTokenAmount(100),
		GasLimit:   0,
		GasFeeCap:  abi.TokenAmount{},
		GasPremium: abi.TokenAmount{},
		Method:     0,
		Params:     nil,
	}

	venuslimit, err := venusFullNode.GasEstimateMessageGas(ctx, &msg, &api.MessageSendSpec{}, headTs.Key())
	if err != nil {
		log.Errorf("venus fial to do GasEstimateMessageGas %v", err)
		return
	}
	lotusBalance, err := nodeFullNode.GasEstimateMessageGas(ctx, &msg, &api.MessageSendSpec{}, headTs.Key())
	if err != nil {
		log.Errorf("lotus fial to do GasEstimateMessageGas %v", err)
		return
	}

	if checkRaw(venuslimit, lotusBalance) {
		log.Errorf("GasEstimateMessageGas fail venus %v, lotus %v", venuslimit, lotusBalance)
		return
	}

	log.Infof("success to check api GasEstimateMessageGas")
}

func CheckGasEstimateFeeCap(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do GasEstimateFeeCap %v", err)
		return
	}

	addr, _ := address.NewFromString("f1kc5ami5ejqept4ydpjme5ytzywzvrgfxxsjy57i")
	actor, err := nodeFullNode.StateGetActor(ctx, addr, headTs.Key())
	if err != nil {
		log.Errorf("fail to do GasEstimateFeeCap %v", err)
		return
	}
	msg := types.Message{
		Version:    0,
		To:         addr,
		From:       addr,
		Nonce:      actor.Nonce + 3,
		Value:      abi.NewTokenAmount(100),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(50000),
		GasPremium: abi.TokenAmount{},
		Method:     0,
		Params:     nil,
	}

	venuslimit, err := venusFullNode.GasEstimateFeeCap(ctx, &msg, 20, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateFeeCap %v", err)
		return
	}
	lotusBalance, err := nodeFullNode.GasEstimateFeeCap(ctx, &msg, 20, headTs.Key())
	if err != nil {
		log.Errorf("fial to do GasEstimateFeeCap %v", err)
		return
	}

	if venuslimit.String() != lotusBalance.String() {
		log.Errorf("GasEstimateFeeCap fail venus %v, lotus %v", venuslimit, lotusBalance)
		return
	}

	log.Infof("success to check api CheckGasEstimateFeeCap")
}

func CheckGasBatchEstimateFeeCap(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do CheckGasBatchEstimateFeeCap %v", err)
		return
	}

	addr, _ := address.NewFromString("f1kc5ami5ejqept4ydpjme5ytzywzvrgfxxsjy57i")
	actor, err := nodeFullNode.StateGetActor(ctx, addr, headTs.Key())
	if err != nil {
		log.Errorf("fail to do CheckGasBatchEstimateFeeCap %v", err)
		return
	}

	msg := types.Message{
		Version:    0,
		To:         addr,
		From:       addr,
		Nonce:      actor.Nonce + 3,
		Value:      abi.NewTokenAmount(100),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(50000),
		GasPremium: abi.TokenAmount{},
		Method:     0,
		Params:     nil,
	}

	estimateMsg := api.EstimateMessage{
		Msg: &msg,
		Spec: &api.MessageSendSpec{
			MaxFee:            abi.NewTokenAmount(100000000000000000),
			GasOverEstimation: 1.25,
		},
	}

	esitimates := []*api.EstimateMessage{&estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg, &estimateMsg}

	venuslimit, err := venusFullNode.GasBatchEstimateMessageGas(ctx, esitimates, actor.Nonce, headTs.Key())
	if err != nil {
		log.Errorf("fial to do CheckGasBatchEstimateFeeCap %v", err)
		return
	}
	lotusBalance, err := nodeFullNode.GasBatchEstimateMessageGas(ctx, esitimates, actor.Nonce, headTs.Key())
	if err != nil {
		log.Errorf("fial to do CheckGasBatchEstimateFeeCap %v", err)
		return
	}

	if !checkRaw(venuslimit, lotusBalance) {
		fmt.Println(showJson(venuslimit))
		fmt.Println(showJson(lotusBalance))
		log.Errorf("CheckGasBatchEstimateFeeCap fail venus %v, lotus %v", showJson(venuslimit), showJson(lotusBalance))
		return
	}

	log.Infof("success to check api CheckGasBatchEstimateFeeCap")
}
