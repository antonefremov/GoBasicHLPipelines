package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func runCalcParallel(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var resAsBytes []byte
	var hash1, hash2 string

	start := time.Now()

	for _, val := range args {
		hash1 = getSingleHashParallel(val)
		hash2 = getMultiHashParallel(hash1)
	}

	end := time.Since(start)
	fmt.Printf("Code execution time for 'runCalcParallel' is %v\n", end)

	resAsBytes = []byte(hash2)

	return shim.Success(resAsBytes)
}
