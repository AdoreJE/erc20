package controller

import (
	"encoding/json"
	"fmt"

	"github.com/erc20/repository"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// TotalSupply is query function
// params - tokenName
// returns the amount of token in existence
func (cc *Controller) TotalSupply(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// the number of params must be one
	if len(params) != 1 {
		return shim.Error("inccorect number of parameters")
	}

	tokenName := params[0]

	// Get ERC20 TotalSupply
	totalSupply, err := repository.GetERC20TotalSupply(stub, tokenName)
	if err != nil {
		return shim.Error(err.Error())
	}

	// convert TotalSupply to Bytes
	totalSupplyBytes, err := json.Marshal(totalSupply)
	if err != nil {
		return shim.Error("failed to Marshal totalSupply, error: " + err.Error())
	}
	fmt.Println(tokenName + "'s totalSupply is " + string(totalSupplyBytes))

	return shim.Success(totalSupplyBytes)
}

// BalanceOf is query function
// params - address
// Returns the amount of tokens owned by address
func (cc *Controller) BalanceOf(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// the number of params must be one
	if len(params) != 1 {
		return shim.Error("inccorect number of parameter")
	}

	address := params[0]

	// get Balance
	amountBytes, err := stub.GetState(address)
	if len(amountBytes) == 0 {
		return shim.Error("failed to GetState, error: " + err.Error())
	}
	if err != nil {
		return shim.Error("failed to GetState, error: " + err.Error())
	}

	fmt.Println(address + "'s balance is " + string(amountBytes))

	return shim.Success(amountBytes)
}

// Allowance is query function
// params - owner's address, spender's address
// returns the remaining amount of token to invoke {transferFrom}
func (cc *Controller) Allowance(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// check the number of params is 2
	if len(params) != 2 {
		return shim.Error("incorrect number of parameters")
	}
	ownerAddress, spenderAddress := params[0], params[1]

	// get amount
	amountBytes, err := repository.GetAllowanceBytes(stub, ownerAddress, spenderAddress, true)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(amountBytes)
}

// ApprovalList is query function
// params - owner's address
// Returns the approval list approved by owner
func (cc *Controller) ApprovalList(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// check the number of params is 1
	if len(params) != 1 {
		return shim.Error("incorrect number of parameters")
	}

	ownerAddress := params[0]

	// get approval List
	approvalSlice, err := repository.GetApprovalList(stub, ownerAddress)
	if err != nil {
		return shim.Error(err.Error())
	}

	// convert approvalSlice to bytes for return
	response, err := json.Marshal(approvalSlice)
	if err != nil {
		return shim.Error("failed to Marshal approvalSlice, error: " + err.Error())
	}

	return shim.Success(response)
}
