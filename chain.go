package main

import (
	"bytes"
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/api"
)

func CheckGetTipsetByKey(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := venusFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do CheckGetTipsetByKey %v", err)
		return
	}

	venusTs, err := venusFullNode.ChainGetTipSet(ctx, headTs.Key())
	if err != nil {
		log.Errorf("fail to do CheckGetTipsetByKey %v", err)
		return
	}

	nodeTs, err := nodeFullNode.ChainGetTipSet(ctx, headTs.Key())
	if err != nil {
		log.Errorf("fail to do CheckGetTipsetByKey %v", err)
		return
	}
	if !checkRaw(venusTs, nodeTs) {
		log.Errorf("ts in lotus and venus is not same venus:%v, lotus:%v", venusTs, nodeTs)
		return
	}
	log.Infof("success to check api GetTipsetByKey")

}

func CheckChainGetRandomnessFromTickets(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	ts, err := venusFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do CheckGetTipsetByKey %v", err)
		return
	}
	addr, _ := address.NewFromString("f0128788")

	for i := crypto.DomainSeparationTag_TicketProduction; i <= crypto.DomainSeparationTag_PoStChainCommit; i++ {
		venusTicket, err := venusFullNode.ChainGetRandomnessFromTickets(ctx, ts.Key(), i, ts.Height()-10, addr.Bytes())
		if err != nil {
			log.Errorf("fail to do ChainGetRandomnessFromTickets %v", err)
			return
		}

		lotusTicket, err := nodeFullNode.ChainGetRandomnessFromTickets(ctx, ts.Key(), i, ts.Height()-10, addr.Bytes())
		if err != nil {
			log.Errorf("fail to do ChainGetRandomnessFromTickets %v", err)
			return
		}
		if !bytes.Equal(venusTicket, lotusTicket) {
			log.Errorf("ticket from venus and lotus domainTag %d is not same venus: %v lotus %v %v", i, venusTicket, lotusTicket)
			return
		}
	}
}

func CheckChainGetRandomnessFromBeacon(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	ts, err := venusFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainGetRandomnessFromBeacon %v", err)
		return
	}
	addr, _ := address.NewFromString("f0128788")

	for i := crypto.DomainSeparationTag_TicketProduction; i <= crypto.DomainSeparationTag_PoStChainCommit; i++ {
		venusRandomness, err := venusFullNode.ChainGetRandomnessFromBeacon(ctx, ts.Key(), i, ts.Height()-10, addr.Bytes())
		if err != nil {
			log.Errorf("fail to do ChainGetRandomnessFromBeacon %v", err)
			return
		}

		lotusRandomness, err := nodeFullNode.ChainGetRandomnessFromBeacon(ctx, ts.Key(), i, ts.Height()-10, addr.Bytes())
		if err != nil {
			log.Errorf("fail to do ChainGetRandomnessFromBeacon %v", err)
			return
		}
		if !bytes.Equal(venusRandomness, lotusRandomness) {
			log.Errorf("beacon from venus and lotus domainTag %d is not same venus: %v lotus %v %v", i, venusRandomness, lotusRandomness)
			return
		}
	}
}

func CheckChainGetBlockMessages(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := venusFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainGetBlockMessages %v", err)
		return
	}

	bid := headTs.Parents().Cids()[0]

	venusMessages, err := venusFullNode.ChainGetBlockMessages(ctx, bid)
	if err != nil {
		log.Errorf("fail to do ChainGetBlockMessages %v", err)
		return
	}

	lotusMessage, err := nodeFullNode.ChainGetBlockMessages(ctx, bid)
	if err != nil {
		log.Errorf("fail to do ChainGetBlockMessages %v", err)
		return
	}
	if !checkRaw(venusMessages, lotusMessage) {
		log.Errorf("ChainGetBlockMessages blocks message of %s in lotus and venus is not same venus:%v, lotus:%v", bid, lotusMessage, lotusMessage)
		return
	}

	log.Infof("success to check api ChainGetBlockMessages")
}

func CheckChainGetTipSetByHeight(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := venusFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainGetTipSetByHeight %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		height := headTs.Height() - abi.ChainEpoch(i)

		venusTs, err := venusFullNode.ChainGetTipSetByHeight(ctx, height, headTs.Key())
		if err != nil {
			log.Errorf("fail to do ChainGetTipSetByHeight %v", err)
			return
		}

		lotusTs, err := nodeFullNode.ChainGetTipSetByHeight(ctx, height, headTs.Key())
		if err != nil {
			log.Errorf("fail to do ChainGetTipSetByHeight %v", err)
			return
		}
		if !checkRaw(venusTs, lotusTs) {
			log.Errorf("ChainGetTipSetByHeight tipset at height %d in lotus and venus is not same venus:%v, lotus:%v", height, venusTs, lotusTs)
			return
		}
	}

	log.Infof("success to check api ChainGetBlockMessages")
}

func CheckChainGetMessage(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainGetBlockMessages %v", err)
		return
	}

	bid := headTs.Parents().Cids()[0]

	messages, err := nodeFullNode.ChainGetBlockMessages(ctx, bid)

	for _, msg := range messages.Cids {
		venusMessage, err := venusFullNode.ChainGetMessage(ctx, msg)
		if err != nil {
			log.Errorf("fail to do ChainGetBlockMessages %v", err)
			return
		}

		lotusMessage, err := nodeFullNode.ChainGetMessage(ctx, msg)
		if err != nil {
			log.Errorf("fail to do ChainGetBlockMessages %v", err)
			return
		}

		if !checkRaw(venusMessage, lotusMessage) {
			log.Errorf("ChainGetBlockMessages message of %s in lotus and venus is not same venus:%v, lotus:%v", msg, venusMessage, lotusMessage)
			return
		}
	}

	log.Infof("success to check api ChainGetBlockMessages")
}

func CheckChainReadObj(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainReadObj %v", err)
		return
	}

	mid := headTs.Blocks()[0].Messages

	venusRaw, err := venusFullNode.ChainReadObj(ctx, mid)
	if err != nil {
		log.Errorf("fial to do ChainReadObj %v", err)
	}

	lotusRaw, err := nodeFullNode.ChainReadObj(ctx, mid)
	if err != nil {
		log.Errorf("fial to do ChainReadObj %v", err)
	}

	if !bytes.Equal(venusRaw, lotusRaw) {
		log.Errorf("ChainReadObj fail venus %v, lotus %v", venusRaw, lotusRaw)
		return
	}
	log.Infof("success to check api ChainReadObj")
}

func CheckChainHasObj(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do ChainHasObj %v", err)
		return
	}

	mid := headTs.Blocks()[0].Messages

	venusRaw, err := venusFullNode.ChainHasObj(ctx, mid)
	if err != nil {
		log.Errorf("fial to do ChainHasObj %v", err)
	}

	lotusRaw, err := nodeFullNode.ChainHasObj(ctx, mid)
	if err != nil {
		log.Errorf("fial to do ChainHasObj %v", err)
	}

	if venusRaw != lotusRaw {
		log.Errorf("ChainHasObj fail venus %v, lotus %v", venusRaw, lotusRaw)
		return
	}
	log.Infof("success to check api ChainHasObj")
}
