package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("main")

func main() {
	app := &cli.App{
		Name:     "chain-noise",
		Usage:    "Generate some spam transactions in the network",
		Flags:    []cli.Flag{},
		Commands: []*cli.Command{runCmd},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

var runCmd = &cli.Command{
	Name: "run",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "venus-url",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "venus-token",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "lotus-url",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "lotus-token",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		logging.SetAllLoggers(logging.LevelInfo)
		ctx := cctx.Context
		venusAPIInfo := APIInfo{
			Addr:  cctx.String("venus-url"),
			Token: []byte(cctx.String("venus-token")),
		}
		addr, err := venusAPIInfo.DialArgs("v1")
		if err != nil {
			return err
		}
		venusFullNode := &api.FullNodeStruct{}
		venusCloser, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin", []interface{}{&venusFullNode.CommonStruct.Internal, &venusFullNode.Internal}, venusAPIInfo.AuthHeader())
		if err != nil {
			return err
		}
		defer venusCloser()

		nodeAPIInfo := cliutil.APIInfo{
			Addr:  cctx.String("lotus-url"),
			Token: []byte(cctx.String("lotus-token")),
		}
		nodeAddr, err := nodeAPIInfo.DialArgs("v1")
		if err != nil {
			return err
		}
		nodeFullNode := &api.FullNodeStruct{}
		nodeCloser, err := jsonrpc.NewMergeClient(ctx, nodeAddr, "Filecoin", []interface{}{&nodeFullNode.CommonStruct.Internal, &nodeFullNode.Internal}, nodeAPIInfo.AuthHeader())
		if err != nil {
			return err
		}
		defer nodeCloser()

		var checkFuncs = []CheckFunc{
			CheckGetTipsetByKey,
			CheckChainGetRandomnessFromTickets,
			CheckChainGetRandomnessFromBeacon,
			CheckChainGetBlockMessages,
			CheckChainGetTipSetByHeight,
			CheckChainGetMessage,
			CheckChainReadObj,
			CheckChainHasObj,

			CheckStateLookupID,
			CheckStateNetworkVersion,
			CheckStateAccountKey,
			CheckStateMinerRecoveries,
			CheckStateMinerFaults,
			CheckStateGetActor,
			CheckStateMinerSectorAllocated,
			CheckStateMinerProvingDeadline,
			CheckStateMinerPartitions,
			CheckStateMinerDeadlines,
			CheckStateMinerInfo,
			CheckStateSectorPartition,
			CheckStateSectorGetInfo,
			CheckStateMinerSectors,

			CheckWalletBalance,
			CheckGasEstimateGasLimit,
			CheckGasEstimateGasPremium,
			CheckGasEstimateMessageGas,
			CheckGasEstimateFeeCap,

			CheckGetBaseInfo,
		}

		for _, cf := range checkFuncs {
			cf(ctx, venusFullNode, nodeFullNode)
		}
		return nil
	},
}

func checkRaw(v1, v2 interface{}) bool {
	v1Raw, _ := json.Marshal(v1)
	v2Raw, _ := json.Marshal(v2)
	return bytes.Equal(v1Raw, v2Raw)
}

type CheckFunc func(ctx context.Context, venusFullNode, nodeFullNode api.FullNode)
