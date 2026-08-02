<?php
require_once __DIR__ . '/../../gen/php/vendor/autoload.php';

use Wallet\V1\WalletBalanceChanged;

$data = ''; // placeholder for received data

$wallet = new WalletBalanceChanged();
$wallet->mergeFromString($data);
echo "Received wallet: " . $wallet->getWalletId() . PHP_EOL;
