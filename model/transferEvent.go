package model

// TransferEvent is the event definition of Transfer
type TransferEvent struct {
	Sender    string `json:"sender"`
	Recipient string `josn:"recipient"`
	Amount    int    `json:"amount"`
}

func NewTransferEvent(sender, recipient string, amount int) *TransferEvent {
	return &TransferEvent{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}
