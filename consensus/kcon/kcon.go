package kcon

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	backend bind.ContractBackend
)

func GetSigners() ([]common.Address, error) {
	// conn, err := ethclient.Dial("/Users/lemon.pattharathon/lemon-fork-chain/posa/go-ethereum/posa-network/node1/geth.ipc")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }
	// posa, err := NewPoSA(common.HexToAddress(posaContractAddress), conn)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %v", err)
	// }
	// posa, err := NewPoSA(common.HexToAddress(posaContractAddress), backend)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a PoSA contract: %v", err)
	// }
	// signerList, err := posa.GetSigners(&bind.CallOpts{})
	// // fmt.Println("--------------Signer--------------:", signerList)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve signer list: %v", err)
	// }
	// fmt.Println("Signer List:", signerList)

	// signers := make([]common.Address, 0)
	// for _, address := range signerList {
	// 	signers = append(signers, address)
	// }

	addresses := []string{"13ec15a841332960aab03306484a0b0903479650"}
	signers := make([]common.Address, 0)
	for _, address := range addresses {
		signers = append(signers, common.HexToAddress(address))
	}
	return signers, nil
}
