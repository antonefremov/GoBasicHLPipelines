package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func runCalcSeq(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var resAsBytes []byte
	var hash1, hash2 string

	for _, val := range args {
		hash1 = getSingleHash(val)
		hash2 = getMultiHash(hash1)
	}

	resAsBytes = []byte(hash2)

	return shim.Success(resAsBytes)
}
