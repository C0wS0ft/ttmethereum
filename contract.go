package ttmethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	Erc20TransferMethodSignature = "0xa9059cbb"
	Erc20TransferEventSignature  = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	Erc20NameSignature           = "0x06fdde03"
	Erc20SymbolSignature         = "0x95d89b41"
	Erc20DecimalsSignature       = "0x313ce567"
	Erc20BalanceOf               = "0x70a08231"
)

type (
	Transaction types.Transaction

	EthereumRequest struct {
		client *ethclient.Client
	}

	TTMTron interface {
		CurrentBlockNumber(context.Context) (uint64, error)
		GetBlockByNumber(context.Context, uint64) (*types.Block, error)
		GetNativeCoinBalance(context.Context, string, string) (uint64, error)

		GetERC20TokenSymbol(context.Context, string, string) (string, error)
		GetERC20TokenDecimals(context.Context, string, string) (uint64, error)
		GetERC20TokenBalance(context.Context, string, string) (uint64, error)
		GetERC20TokenName(context.Context, string) (string, error)

		GetTransactionFrom(ctx context.Context, transaction *types.Transaction) (string, error)
	}
)
