package ttmethereum

import (
	"context"
	token "github.com/C0wS0ft/ttmethereum/erc20token"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"log"
	"math/big"
)

func Init(baseURL string) *EthereumRequest {
	client, err := ethclient.Dial(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	return &EthereumRequest{client: client}
}

// CurrentBlockNumber return current block number
func (e *EthereumRequest) CurrentBlockNumber(ctx context.Context) (uint64, error) {
	number, err := e.client.BlockNumber(ctx)

	return number, errors.Wrap(err, "unable to get current block number")
}

// GetBlockByNumber return block by number
func (e *EthereumRequest) GetBlockByNumber(ctx context.Context, num uint64) (*types.Block, error) {
	block, err := e.client.BlockByNumber(ctx, big.NewInt(int64(num)))

	return block, errors.Wrap(err, "unable to get block by number")
}

func (e *EthereumRequest) GetNativeBalance(ctx context.Context, address string) (uint64, error) {
	var res *big.Int
	var err error

	ethAddr := common.HexToAddress(address)
	res, err = e.client.BalanceAt(ctx, ethAddr, nil)

	if err != nil {
		return 0, err
	}

	return res.Uint64(), nil
}

func (e *EthereumRequest) GetERC20TokenName(ctx context.Context, tokenAddress string) (string, error) {
	contractAddr := common.HexToAddress(tokenAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: common.FromHex(Erc20NameSignature),
	}

	res, err := e.client.CallContract(ctx, callMsg, nil)

	if err != nil {
		return "", err
	}

	str := common.Bytes2Hex(res)
	name, err := DecodeConstantToSymbol(str)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (e *EthereumRequest) GetERC20TokenSymbol(ctx context.Context, tokenAddress string) (string, error) {
	contractAddr := common.HexToAddress(tokenAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: common.FromHex(Erc20SymbolSignature),
	}

	res, err := e.client.CallContract(ctx, callMsg, nil)

	if err != nil {
		return "", err
	}

	// 64
	//0000000000000000000000000000000000000000000000000000000000000020
	//0000000000000000000000000000000000000000000000000000000000000004
	//5a45544100000000000000000000000000000000000000000000000000000000

	str := common.Bytes2Hex(res)
	symbol, err := DecodeConstantToSymbol(str)

	if err != nil {
		return "", err
	}

	return symbol, nil
}

func (e *EthereumRequest) GetERC20TokenDecimals(ctx context.Context, tokenAddress string) (uint64, error) {
	contractAddr := common.HexToAddress(tokenAddress)
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: common.FromHex(Erc20DecimalsSignature),
	}

	res, err := e.client.CallContract(ctx, callMsg, nil)

	if err != nil {
		return 0, err
	}

	str := common.Bytes2Hex(res)
	decimals, err := HexToInt256(str)

	if err != nil {
		return 0, err
	}

	return decimals.Uint64(), nil
}

func (e *EthereumRequest) GetERC20TokenBalance(ctx context.Context, tokenAddress string, userAddress string) (uint64, error) {
	tokenAddr := common.HexToAddress(tokenAddress)
	instance, err := token.NewToken(tokenAddr, e.client)
	if err != nil {
		return 0, err
	}

	userAddr := common.HexToAddress(userAddress)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, userAddr)
	if err != nil {
		return 0, err
	}

	return bal.Uint64(), nil
}

func (e *EthereumRequest) GetTransactionFrom(ctx context.Context, transaction *types.Transaction) (string, error) {
	msg, err := transaction.AsMessage(types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return "", err
	}

	return msg.From().Hex(), nil // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
}
