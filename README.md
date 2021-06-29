# for api test

StateMinerSectors(context.Context, address.Address, *bitfield.BitField, types.TipSetKey) ([]*miner.SectorOnChainInfo, error)

    StateSectorPreCommitInfo(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (miner.SectorPreCommitOnChainInfo, error)

StateSectorGetInfo(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (*miner.SectorOnChainInfo, error)

StateSectorPartition(ctx context.Context, maddr address.Address, sectorNumber abi.SectorNumber, tok types.TipSetKey) (*miner.SectorLocation, error)

StateMinerInfo(context.Context, address.Address, types.TipSetKey) (miner.MinerInfo, error)

StateMinerDeadlines(context.Context, address.Address, types.TipSetKey) ([]api.Deadline, error)

StateMinerPartitions(context.Context, address.Address, uint64, types.TipSetKey) ([]api.Partition, error)

StateMinerProvingDeadline(context.Context, address.Address, types.TipSetKey) (*dline.Info, error)

    StateMinerPreCommitDepositForPower(context.Context, address.Address, miner.SectorPreCommitInfo, types.TipSetKey) (types.BigInt, error)

    StateMinerInitialPledgeCollateral(context.Context, address.Address, miner.SectorPreCommitInfo, types.TipSetKey) (types.BigInt, error)

StateMinerSectorAllocated(context.Context, address.Address, abi.SectorNumber, types.TipSetKey) (bool, error)

StateGetActor(ctx context.Context, actor address.Address, ts types.TipSetKey) (*types.Actor, error)

    StateMarketStorageDeal(context.Context, abi.DealID, types.TipSetKey) (*api.MarketDeal, error)

StateMinerFaults(context.Context, address.Address, types.TipSetKey) (bitfield.BitField, error)

StateMinerRecoveries(context.Context, address.Address, types.TipSetKey) (bitfield.BitField, error)

StateAccountKey(context.Context, address.Address, types.TipSetKey) (address.Address, error)

StateNetworkVersion(context.Context, types.TipSetKey) (network.Version, error)

StateLookupID(context.Context, address.Address, types.TipSetKey) (address.Address, error)



ChainHead(context.Context) (*types.TipSet, error)

ChainGetRandomnessFromTickets(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)

ChainGetRandomnessFromBeacon(ctx context.Context, tsk types.TipSetKey, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)

ChainGetTipSetByHeight(context.Context, abi.ChainEpoch, types.TipSetKey) (*types.TipSet, error)

ChainGetBlockMessages(context.Context, cid.Cid) (*api.BlockMessages, error)

ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error)

ChainReadObj(context.Context, cid.Cid) ([]byte, error)

ChainHasObj(context.Context, cid.Cid) (bool, error)

ChainGetTipSet(ctx context.Context, key types.TipSetKey) (*types.TipSet, error)



WalletSign(context.Context, address.Address, []byte) (*crypto.Signature, error)

WalletHas(context.Context, address.Address) (bool, error)

StateCall(context.Context, *types.Message, types.TipSetKey) (*api.InvocResult, error)

MpoolPushMessage(context.Context, *types.Message, *api.MessageSendSpec) (*types.SignedMessage, error)

GasEstimateMessageGas(context.Context, *types.Message, *api.MessageSendSpec, types.TipSetKey) (*types.Message, error)

GasEstimateFeeCap(context.Context, *types.Message, int64, types.TipSetKey) (types.BigInt, error)

GasEstimateGasPremium(_ context.Context, nblocksincl uint64, sender address.Address, gaslimit int64, tsk types.TipSetKey) (types.BigInt, error)

WalletBalance(context.Context, address.Address) (types.BigInt, error)

StateSearchMsg(ctx context.Context, from types.TipSetKey, msg cid.Cid, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)

StateWaitMsg(ctx context.Context, cid cid.Cid, confidence uint64, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)


