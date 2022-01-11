package kcon

import "github.com/ethereum/go-ethereum/common"

func GetSigners() ([]common.Address, error) {
	addresses := []string{"13ec15a841332960aab03306484a0b0903479650"}
	signers := make([]common.Address, 0)
	for _, address := range addresses {
		signers = append(signers, common.HexToAddress(address))
	}
	return signers, nil
}
