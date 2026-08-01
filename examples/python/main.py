import sys

sys.path.append("../../gen/python")


from wallet.v1 import wallet_pb2


wallet = wallet_pb2.WalletBalanceChanged()


wallet.wallet_id = "wallet-123"

wallet.asset = "BTC"


wallet.balance.currency = "USD"

wallet.balance.amount = "10000"


wallet.metadata.request_id = "req-001"


data = wallet.SerializeToString()


print(data)


new_wallet = wallet_pb2.WalletBalanceChanged()

new_wallet.ParseFromString(data)


print(new_wallet)