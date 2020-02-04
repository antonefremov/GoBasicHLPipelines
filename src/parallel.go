package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func runCalcParallel(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var resAsBytes []byte
	var hash1, hash2 string

	for _, val := range args {
		hash1 = getSingleHashParallel(val)
		hash2 = getMultiHashParallel(hash1)
	}

	resAsBytes = []byte(hash2)

	return shim.Success(resAsBytes)
}
