package main

import (
	"fmt"

	walletv1 "github.com/example/poly-contracts/gen/go/wallet/v1"
)

func main() {
	// This example would normally consume a message from a queue.
	// For demonstration, we just create a sample message.
	data := []byte{} // placeholder for received data

	var wallet walletv1.WalletBalanceChanged
	if err := wallet.Unmarshal(data); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Received wallet:", wallet)
}
