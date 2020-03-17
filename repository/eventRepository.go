package repository

import (
	"encoding/json"

	"github.com/erc20/model"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	TransferEventKey = "transferEvent"
	ApprovalEventKey = "approvalEvent"
)

func EmitTransferEvent(stub shim.ChaincodeStubInterface, sender, spender string, amount int) error {
	// emit transfer event
	// transferEvent := TransferEvent{Sender: callerAddress, Recipient: recipientAddress, Amount: transferAmountInt}
	transferEvent := model.NewTransferEvent(sender, spender, amount)
	transferEventBytes, err := json.Marshal(transferEvent)
	if err != nil {
		return model.NewCustomError(model.MarshalErrorType, "transferEvent", err.Error())
	}
	err = stub.SetEvent("transferEvent", transferEventBytes)
	if err != nil {
		return model.NewCustomError(model.SetEventErrorType, "transferEvent", err.Error())
	}
	return nil
}

func EmitApprovalEvent(stub shim.ChaincodeStubInterface, owner, spender string, allowance int) error {
	// emit approval event
	// approvalEvent := Approval{Owner: ownerAddress, Spender: spenderAddress, Allowance: allowanceAmountInt}
	approvalEvent := model.NewApproval(owner, spender, allowance)
	approvalBytes, err := json.Marshal(approvalEvent)
	if err != nil {
		return model.NewCustomError(model.MarshalErrorType, "approvalEvent", err.Error())
	}
	err = stub.SetEvent("approvalEvent", approvalBytes)
	if err != nil {
		return model.NewCustomError(model.SetEventErrorType, "approvalEvent", err.Error())
	}
	return nil
}
