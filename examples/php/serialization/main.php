<?php
require_once __DIR__ . '/../../gen/php/vendor/autoload.php';

use Wallet\V1\WalletBalanceChanged;
use Wallet\V1\Money;
use Wallet\V1\Metadata;

$wallet = new WalletBalanceChanged();
$wallet->setWalletId('wallet-123');
$wallet->setAsset('BTC');

$balance = new Money();
$balance->setCurrency('USD');
$balance->setAmount('10000');
$wallet->setBalance($balance);

$metadata = new Metadata();
$metadata->setRequestId('req-001');
$wallet->setMetadata($metadata);

$data = $wallet->serializeToString();
echo "Serialized data: " . bin2hex($data) . PHP_EOL;

$newWallet = new WalletBalanceChanged();
$newWallet->mergeFromString($data);
echo "Deserialized: " . $newWallet->getWalletId() . PHP_EOL;
