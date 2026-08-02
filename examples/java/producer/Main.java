package main;

import wallet.v1.WalletBalanceChanged;
import wallet.v1.Money;
import wallet.v1.Metadata;

public class Main {
    public static void main(String[] args) throws Exception {
        WalletBalanceChanged wallet = WalletBalanceChanged.newBuilder()
                .setWalletId("wallet-123")
                .setAsset("BTC")
                .setBalance(Money.newBuilder()
                        .setCurrency("USD")
                        .setAmount("10000")
                        .build())
                .setMetadata(Metadata.newBuilder()
                        .setRequestId("req-001")
                        .build())
                .build();

        byte[] data = wallet.toByteArray();
        System.out.println("Serialized data: " + java.util.Arrays.toString(data));

        WalletBalanceChanged newWallet = WalletBalanceChanged.parseFrom(data);
        System.out.println("Deserialized: " + newWallet);
    }
}
