package kcon

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetSigners() ([]common.Address, error) {

	conn, err := ethclient.Dial("/Users/lemon.pattharathon/lemon-fork-chain/posa/go-ethereum/posa-network/node1/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	posa, err := NewPoSA(common.HexToAddress("0xDB80B3722575bB75dE312c4A1619F2fa38D39F2f"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	signerList, err := posa.GetSigners(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve signer list: %v", err)
	}
	fmt.Println("Signer List:", signerList)

	signers := make([]common.Address, 0)
	for _, address := range signerList {
		signers = append(signers, address)
	}

	// addresses := []string{"13ec15a841332960aab03306484a0b0903479650"}
	// signers := make([]common.Address, 0)
	// for _, address := range addresses {
	// 	signers = append(signers, common.HexToAddress(address))
	// }
	return signers, nil
}
