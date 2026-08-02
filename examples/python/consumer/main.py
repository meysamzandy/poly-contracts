import sys

sys.path.append("../../gen/python")

from wallet.v1 import wallet_pb2

def main():
    # placeholder for received data
    data = b''

    wallet = wallet_pb2.WalletBalanceChanged()
    wallet.ParseFromString(data)
    print(wallet)

if __name__ == "__main__":
    main()
