/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/erc20/model"
	"github.com/erc20/repository"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const initAmount = 100000
const tokenName = "dappToken"
const address = "dappcampus"

func Test_Init_success(t *testing.T) {
	cc := NewChaincode()
	stub := shim.NewMockStub("erc20", cc)
	res := stub.MockInit("1", [][]byte{[]byte("init"), []byte(tokenName), []byte("dt"), []byte(address), []byte(strconv.Itoa(initAmount))})
	if res.Status != shim.OK {
		t.FailNow()
	}

	// check totalSupply
	erc20 := model.ERC20Metadata{}
	erc20Bytes, _ := stub.GetState(tokenName)
	json.Unmarshal(erc20Bytes, &erc20)
	if *erc20.GetTotalSupply() != initAmount {
		t.FailNow()
	}

	// check dappcampus balance
	balance, _ := repository.GetBalance(stub, address, false)
	if *balance != initAmount {
		t.FailNow()
	}
}

func TestInvoke(t *testing.T) {
	cc := new(ERC20Chaincode)
	stub := shim.NewMockStub("chaincode", cc)
	res := stub.MockInit("1", [][]byte{[]byte("initFunc")})
	if res.Status != shim.OK {
		t.Error("Init failed", res.Status, res.Message)
	}
	res = stub.MockInvoke("1", [][]byte{[]byte("invokeFunc")})
	if res.Status != shim.OK {
		t.Error("Invoke failed", res.Status, res.Message)
	}
}
