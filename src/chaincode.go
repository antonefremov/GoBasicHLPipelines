package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode - %s", err)
	}
}

// Init is called during Instantiate transaction
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called to update or query the ledger in a proposal transaction
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	switch function {

	case "init":
		return t.Init(stub)
	case "runCalcSeq":
		return runCalcSeq(stub, args)
	case "runCalcParallel":
		return runCalcParallel(stub, args)
	default:
		return shim.Error("Invoke: no such method " + function)

	}
}
