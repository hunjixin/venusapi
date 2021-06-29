package main

import (
	"bytes"
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
)

func CheckStateLookupID(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateLookupID %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f3wylwd6pclppme4qmbgwled5xpsbgwgqbn2alxa7yahg2gnbfkipsdv6m764xm5coizujmwdmkxeugplmorha")
	//add3, _  := address.NewFromString("t3qbbsp5aczfsfvlbqd4fvo7csjkvl3qd7ufs23woycgseart636cwsixuuksktq2xpipywdzdpv5epewllfiq")

	for _, addr := range []address.Address{add1, add2} {
		venusId, err := venusFullNode.StateLookupID(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("venus fail to do StateLookupID %v", err)
			return
		}

		lotusId, err := nodeFullNode.StateLookupID(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("lotus fial to do StateLookupID %v", err)
			return
		}

		if venusId.String() != lotusId.String() {
			log.Errorf("StateLookupID address %s fail venus %v, lotus %v", addr, venusId, lotusId)
			return
		}
	}

	log.Infof("success to check api StateLookupID")
}

func CheckStateNetworkVersion(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateNetworkVersion %v", err)
		return
	}

	venusVersion, err := venusFullNode.StateNetworkVersion(ctx, headTs.Parents())
	if err != nil {
		log.Errorf("fial to do StateNetworkVersion %v", err)
		return
	}

	lotusVersion, err := nodeFullNode.StateNetworkVersion(ctx, headTs.Parents())
	if err != nil {
		log.Errorf("fial to do StateNetworkVersion %v", err)
		return
	}

	if venusVersion != lotusVersion {
		log.Errorf("StateNetworkVersion address fail venus %v, lotus %v", venusVersion, lotusVersion)
		return
	}

	log.Infof("success to check api StateNetworkVersion")
}

func CheckStateAccountKey(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateAccountKey %v", err)
		return
	}

	add1, _ := address.NewFromString("f0809528")
	add2, _ := address.NewFromString("f3wylwd6pclppme4qmbgwled5xpsbgwgqbn2alxa7yahg2gnbfkipsdv6m764xm5coizujmwdmkxeugplmorha")
	//add3, _  := address.NewFromString("t3qbbsp5aczfsfvlbqd4fvo7csjkvl3qd7ufs23woycgseart636cwsixuuksktq2xpipywdzdpv5epewllfiq")

	for _, addr := range []address.Address{add1, add2} {
		venusId, err := venusFullNode.StateAccountKey(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateAccountKey %v", err)
			return
		}

		lotusId, err := nodeFullNode.StateAccountKey(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateAccountKey %v", err)
			return
		}

		if venusId.String() != lotusId.String() {
			log.Errorf("StateAccountKey address %s fail venus %v, lotus %v", addr, venusId, lotusId)
			return
		}
	}

	log.Infof("success to check api StateAccountKey")
}

func CheckStateMinerRecoveries(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerRecoveries %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		venusRecovery, err := venusFullNode.StateMinerRecoveries(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerRecoveries %v", err)
			return
		}

		lotusRecovery, err := nodeFullNode.StateMinerRecoveries(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerRecoveries %v", err)
			return
		}

		venusBytes, _ := venusRecovery.MarshalJSON()
		lotusBytes, _ := lotusRecovery.MarshalJSON()
		if !bytes.Equal(venusBytes, lotusBytes) {
			log.Errorf("StateMinerRecoveries address %s fail venus %v, lotus %v", addr, string(venusBytes), string(lotusBytes))
			return
		}
	}

	log.Infof("success to check api StateMinerRecoveries")
}

func CheckStateMinerFaults(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerFaults %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		venuFaults, err := venusFullNode.StateMinerFaults(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerFaults %v", err)
			return
		}

		lotuFaults, err := nodeFullNode.StateMinerFaults(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerFaults %v", err)
			return
		}

		venusBytes, _ := venuFaults.MarshalJSON()
		lotusBytes, _ := lotuFaults.MarshalJSON()
		if !bytes.Equal(venusBytes, lotusBytes) {
			log.Errorf("StateMinerFaults address %s fail venus %v, lotus %v", addr, string(venusBytes), string(lotusBytes))
			return
		}
	}

	log.Infof("success to check api StateMinerFaults")
}

func CheckStateGetActor(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateGetActor %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f3sfyonhd3apsolzlpl5uy2a7j7jyktekp7v365l2uqo4chmmf7zmkmsry5qru562yhetnruzflmcnldwow6uq")
	add5, _ := address.NewFromString("f3qzprefkeragndcicaqgztojarm4pzohn7swwqtmtcx42wykpgxtz6rtpn7xsderun5kigfopv3tydhddx4na")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		venusActor, err := venusFullNode.StateGetActor(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateGetActor %v", err, addr.String())
			return
		}

		lotusActor, err := nodeFullNode.StateGetActor(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateGetActor %v", err)
			return
		}
		venusBuf := bytes.NewBufferString("")
		venusActor.MarshalCBOR(venusBuf)

		lotusBuf := bytes.NewBufferString("")
		lotusActor.MarshalCBOR(lotusBuf)
		if !bytes.Equal(lotusBuf.Bytes(), venusBuf.Bytes()) {
			log.Errorf("StateGetActor address %s fail venus %v, lotus %v", addr, venusActor, lotusActor)
			return
		}
	}

	log.Infof("StateGetActor to check api StateMinerFaults")
}

func CheckStateMinerSectorAllocated(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerSectorAllocated %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")

	for _, addr := range []address.Address{add1, add2, add3} {
		for i := 900; i < 1000; i++ {
			venusR, err := venusFullNode.StateMinerSectorAllocated(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectorAllocated %v", err)
				return
			}

			lotusR, err := nodeFullNode.StateMinerSectorAllocated(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectorAllocated %v", err)
				return
			}

			if venusR != lotusR {
				log.Errorf("StateMinerSectorAllocated address %s  sector %d fail venus %v, lotus %v", addr, i, venusR, lotusR)
				return
			}
		}

		for i := 900000000000; i < 900000000010; i++ {
			venusR, err := venusFullNode.StateMinerSectorAllocated(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectorAllocated %v", err)
				return
			}

			lotusR, err := nodeFullNode.StateMinerSectorAllocated(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectorAllocated %v", err)
				return
			}

			if venusR != lotusR {
				log.Errorf("StateMinerSectorAllocated address %s  sector %d fail venus %v, lotus %v", addr, i, venusR, lotusR)
				return
			}
		}
	}

	log.Infof("StateMinerSectorAllocated to check api StateMinerFaults")
}

/*
func CheckStateMinerInitialPledgeCollateral(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do CheckStateMinerInitialPledgeCollateral %v", err)
		return
	}

	add1, _  := address.NewFromString("f0128788")
	add2, _  := address.NewFromString("f0127595")
	add3, _  := address.NewFromString("f0123261")

	for _, addr := range []address.Address{add1, add2,add3}  {
			for i:=10;i<100;i++ {
				precomitInfo, err := nodeFullNode.StateSectorPreCommitInfo(ctx, addr, abi.SectorNumber(i), headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateMinerInitialPledgeCollateral %v", err)
					continue
				}
				venusR, err := venusFullNode.StateMinerInitialPledgeCollateral(ctx, addr, precomitInfo, headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateMinerInitialPledgeCollateral %v", err)
					return
				}

				lotusR, err := nodeFullNode.StateMinerInitialPledgeCollateral(ctx, addr,miner.SectorPreCommitInfo{

				}, headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateMinerInitialPledgeCollateral %v", err)
					return
				}

				if venusR != lotusR {
					log.Errorf("StateMinerSectorAllocated address %s  sector %d fail venus %v, lotus %v", addr, i, venusR, lotusR)
					return
				}
			}
	}

	log.Infof("StateMinerSectorAllocated to check api StateMinerFaults")
}*/

func CheckStateMinerProvingDeadline(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerProvingDeadline %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		venusDeadline, err := venusFullNode.StateMinerProvingDeadline(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerProvingDeadline %v", err)
			return
		}

		lotusDeadline, err := nodeFullNode.StateMinerProvingDeadline(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerProvingDeadline %v", err)
			return
		}

		if !checkRaw(venusDeadline, lotusDeadline) {
			log.Errorf("StateMinerProvingDeadline address %s fail venus %v, lotus %v", addr, venusDeadline, lotusDeadline)
			return
		}
	}

	log.Infof("success to check api StateMinerProvingDeadline")
}

func CheckStateMinerPartitions(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerPartitions %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		deadLines, err := nodeFullNode.StateMinerProvingDeadline(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerPartitions %v", err)
			return
		}

		venusPartitions, err := venusFullNode.StateMinerPartitions(ctx, addr, deadLines.Index, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerPartitions %v", err)
			return
		}

		lotusPartitions, err := nodeFullNode.StateMinerPartitions(ctx, addr, deadLines.Index, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerPartitions %v", err)
			return
		}

		if !checkRaw(venusPartitions, lotusPartitions) {
			log.Errorf("StateMinerPartitions address %s fail venus %v, lotus %v", addr, venusPartitions, lotusPartitions)
			return
		}
	}

	log.Infof("success to check api StateMinerPartitions")
}

func CheckStateMinerDeadlines(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerDeadlines %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {

		venusDeadLines, err := venusFullNode.StateMinerDeadlines(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerDeadlines %v", err)
			return
		}

		lotusDeadLines, err := nodeFullNode.StateMinerDeadlines(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerDeadlines %v", err)
			return
		}

		if !checkRaw(venusDeadLines, lotusDeadLines) {
			log.Errorf("StateMinerDeadlines address %s fail venus %v, lotus %v", addr, venusDeadLines, lotusDeadLines)
			return
		}
	}

	log.Infof("success to check api StateMinerDeadlines")
}

func CheckStateMinerInfo(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerInfo %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {

		venusMinerInfo, err := venusFullNode.StateMinerInfo(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerInfo %v", err)
			return
		}

		lotusMinerInfo, err := nodeFullNode.StateMinerInfo(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerInfo %v", err)
			return
		}

		if !checkRaw(venusMinerInfo, lotusMinerInfo) {
			log.Errorf("StateMinerInfo address %s fail venus %v, lotus %v", addr, venusMinerInfo, lotusMinerInfo)
			return
		}
	}

	log.Infof("success to check api StateMinerInfo")
}

func CheckStateSectorPartition(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateSectorPartition %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		for i := 200; i < 250; i++ {
			sinfo, err := nodeFullNode.StateSectorGetInfo(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateSectorPartition %v", err)
				return
			}

			if sinfo != nil {
				venusPartition, err := venusFullNode.StateSectorPartition(ctx, addr, abi.SectorNumber(i), headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateSectorPartition %v", err)
					return
				}
				lotusPartition, err := nodeFullNode.StateSectorPartition(ctx, addr, abi.SectorNumber(i), headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateSectorPartition %v", err)
					return
				}

				if !checkRaw(venusPartition, lotusPartition) {
					log.Errorf("StateSectorPartition address %s fail venus %v, lotus %v", addr, venusPartition, lotusPartition)
					return
				}
			}
		}

	}

	log.Infof("success to check api StateSectorPartition")
}

func CheckStateSectorGetInfo(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateSectorGetInfo %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		for i := 200; i < 250; i++ {
			has, err := nodeFullNode.StateMinerSectorAllocated(ctx, addr, abi.SectorNumber(i), headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateSectorGetInfo %v", err)
				return
			}

			if has {
				venusSectorInfo, err := venusFullNode.StateSectorGetInfo(ctx, addr, abi.SectorNumber(i), headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateSectorGetInfo %v", err)
					return
				}
				lotusSectorInfo, err := nodeFullNode.StateSectorGetInfo(ctx, addr, abi.SectorNumber(i), headTs.Parents())
				if err != nil {
					log.Errorf("fial to do StateSectorGetInfo %v", err)
					return
				}

				if !checkRaw(venusSectorInfo, lotusSectorInfo) {
					log.Errorf("StateSectorGetInfo address %s fail venus %v, lotus %v", addr, venusSectorInfo, lotusSectorInfo)
					return
				}
			}
		}

	}

	log.Infof("success to check api StateSectorGetInfo")
}

func CheckStateMinerSectors(ctx context.Context, venusFullNode, nodeFullNode api.FullNode) {
	headTs, err := nodeFullNode.ChainHead(ctx)
	if err != nil {
		log.Errorf("fail to do StateMinerSectors %v", err)
		return
	}

	add1, _ := address.NewFromString("f0128788")
	add2, _ := address.NewFromString("f0127595")
	add3, _ := address.NewFromString("f0123261")
	add4, _ := address.NewFromString("f0135467")
	add5, _ := address.NewFromString("f0142720")
	for _, addr := range []address.Address{add1, add2, add3, add4, add5} {
		deadlines, err := nodeFullNode.StateMinerProvingDeadline(ctx, addr, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerSectors %v", err)
			return
		}
		partitions, err := nodeFullNode.StateMinerPartitions(ctx, addr, deadlines.Index, headTs.Parents())
		if err != nil {
			log.Errorf("fial to do StateMinerSectors %v", err)
			return
		}

		for _, partition := range partitions {
			venusSectorInfo, err := venusFullNode.StateMinerSectors(ctx, addr, &partition.ActiveSectors, headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectors %v", err)
				return
			}
			lotusSectorInfo, err := nodeFullNode.StateMinerSectors(ctx, addr, &partition.ActiveSectors, headTs.Parents())
			if err != nil {
				log.Errorf("fial to do StateMinerSectors %v", err)
				return
			}

			if !checkRaw(venusSectorInfo, lotusSectorInfo) {
				log.Errorf("StateMinerSectors address %s fail venus %v, lotus %v", addr, venusSectorInfo, lotusSectorInfo)
				return
			}
		}
	}

	log.Infof("success to check api StateMinerSectors")
}
