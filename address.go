package ethflags

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

var zeroAddress = common.HexToAddress("0x0")

func (a *Address) Set(value string) error {
	if common.IsHexAddress(value) == false {
		return fmt.Errorf("invalid hex address value")
	}
	*a = Address(common.HexToAddress(value))
	return nil
}

func (a *Address) Type() string {
	return "common.Address"
}

func (a Address) String() string {
	return (common.Address(a)).String()
}

func (a *Address) Address() common.Address {
	return common.Address(*a)
}
