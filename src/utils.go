package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"strconv"
	"testing"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const signerSalt string = ""

func initChaincode(test *testing.T) *shim.MockStub {
	stub := shim.NewMockStub("testingStub", new(SimpleChaincode))
	result := stub.MockInit("000", nil)

	if result.Status != shim.OK {
		test.FailNow()
	}
	return stub
}

func invoke(test *testing.T, stub *shim.MockStub, function string, args [][]byte) []byte {
	const transactionID = "000"

	// prepend the function name as the first item
	args = append([][]byte{[]byte(function)}, args...)

	// prepare the parameters for printing
	byteDivider := []byte{','}
	byteArrayToPrint := bytes.Join(args[1:], byteDivider)

	// print information just before the call
	fmt.Println("Call:    ", function, "(", string(byteArrayToPrint), ")")

	// perform the MockInvoke call
	result := stub.MockInvoke(transactionID, args)

	// print the Invoke results
	fmt.Println("RetCode: ", result.Status)
	fmt.Println("RetMsg:  ", result.Message)
	fmt.Println("Payload: ", string(result.Payload))

	if result.Status != shim.OK {
		fmt.Println("Invoke", function, "failed", string(result.Message))
		return nil
	}

	return []byte(result.Payload)
}

func getMd5(data string) string {
	data += signerSalt
	dataHash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	time.Sleep(10 * time.Millisecond)
	return dataHash
}

func getCrc32(data string) string {
	data += signerSalt
	chsum := crc32.ChecksumIEEE([]byte(data))
	dataHash := strconv.FormatUint(uint64(chsum), 10)
	time.Sleep(time.Second)
	return dataHash
}

// gets a crc32(data) + "~" + crc32(md5(data)) value
func getSingleHash(data string) string {
	var result string
	crc32Hash1 := getCrc32(data)
	md5Hash := getMd5(data)
	crc32Hash2 := getCrc32(md5Hash)
	result = crc32Hash1 + "~" + crc32Hash2
	return result
}

// gets a crc32(i + data) where i = 0..5
func getMultiHash(data string) string {
	var result string
	for i := 0; i < 6; i++ {
		sI := strconv.Itoa(i)
		result += getCrc32(sI + data)
	}
	return result
}
