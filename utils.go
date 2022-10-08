package ttmethereum

import (
	"encoding/hex"
	"github.com/pkg/errors"
	"math/big"
)

func HexToInt256(hex string) (*big.Int, error) {
	var bignum, ok = new(big.Int).SetString(hex, 16)

	if ok {
		return bignum, nil
	}

	return nil, errors.New("failed to convert")
}

func DecodeConstantToSymbol(in string) (string, error) {
	if len(in) < 192 {
		return "", errors.New("input should be >192")
	}
	if len(in)%64 != 0 {
		return "", errors.New("input should divide by 64")
	}

	length, err := HexToInt256(in[64:128])

	if err != nil {
		return "", err
	}

	symbolLength := length.Uint64()

	bs, err := hex.DecodeString(in[128 : 192-64+symbolLength*2])
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
