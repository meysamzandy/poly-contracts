package main;

import wallet.v1.WalletBalanceChanged;

public class Main {
    public static void main(String[] args) throws Exception {
        // placeholder for received data
        byte[] data = new byte[0];

        WalletBalanceChanged wallet = WalletBalanceChanged.parseFrom(data);
        System.out.println("Received wallet: " + wallet);
    }
}
