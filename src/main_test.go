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

func TestParallel(test *testing.T) {
	stub := initChaincode(test)

	inputDataAsBytes := [][]byte{[]byte(strconv.Itoa(inputData[0])), []byte(strconv.Itoa(inputData[1]))}

	start := time.Now()
	resultBytes := invoke(test, stub, "runCalcParallel", inputDataAsBytes)
	end := time.Since(start)
	test.Logf("The code execution time is: %v\n", end)
	fmt.Println("Received result: \n", string(resultBytes))
}
