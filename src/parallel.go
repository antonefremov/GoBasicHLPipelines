package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func runCalcParallel(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var resAsBytes []byte
	var result string

	hashJobs := []job{
		job(func(in, out chan interface{}) {
			for _, seqNum := range args {
				out <- seqNum
			}
		}),
		job(getSingleHashParallel),
		job(getMultiHashParallel),
		job(func(in, out chan interface{}) {
			for input := range in {
				data, ok := input.(string)
				if !ok {
					shim.Error("Can't convert result data to string")
				}
				fmt.Println("Appending a new hash " + data)
				result += data
			}
		}),
	}

	start := time.Now()

	runPipeline(hashJobs...)

	end := time.Since(start)
	fmt.Printf("Code execution time for 'runCalcParallel' is %v\n", end)

	resAsBytes = []byte(result)

	return shim.Success(resAsBytes)
}
