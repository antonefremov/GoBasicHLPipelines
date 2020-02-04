package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var inputData = []int{0, 1}

func TestSeq(test *testing.T) {
	stub := initChaincode(test)

	inputDataAsBytes := [][]byte{[]byte(strconv.Itoa(inputData[0]))}

	start := time.Now()
	resultBytes := invoke(test, stub, "runCalcSeq", inputDataAsBytes)
	end := time.Since(start)
	test.Logf("The code execution time is: %v\n", end)
	fmt.Println("Received result: \n", string(resultBytes))
}
