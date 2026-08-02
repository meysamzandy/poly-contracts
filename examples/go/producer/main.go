package main

import (
	"fmt"

	walletv1 "github.com/example/poly-contracts/gen/go/wallet/v1"
)

func main() {
	wallet := &walletv1.WalletBalanceChanged{
		WalletId: "wallet-123",
		Asset:    "BTC",
		Balance: &walletv1.Money{
			Currency: "USD",
			Amount:   "10000",
		},
		Metadata: &walletv1.Metadata{
			RequestId: "req-001",
		},
	}

	data, err := wallet.Marshal()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Serialized data:", data)

	newWallet := &walletv1.WalletBalanceChanged{}
	if err := newWallet.Unmarshal(data); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deserialized:", newWallet)
}
