module github.com/hunjixin/venusapi

go 1.16

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.1-0.20210506134452-99b279731c48
	github.com/filecoin-project/lotus v1.10.0
	github.com/google/martian v2.1.0+incompatible
	github.com/ipfs/go-log/v2 v2.1.3 // indirect
	github.com/multiformats/go-multiaddr v0.3.1 // indirect
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi