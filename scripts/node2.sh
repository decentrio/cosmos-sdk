#!/usr/bin/env bash

set -o errexit
set -o nounset
set -x

ROOT=$PWD

SIMD="$ROOT/build/simdv2"
# CONFIG="${CONFIG:-$HOME/.simappv2/config}"

COSMOS_BUILD_OPTIONS=v2 make build     

# if [ -d "$($SIMD config home)" ]; then rm -rv $($SIMD config home); fi

# $SIMD testnet init-files --chain-id=testing --output-dir=./testnet --validator-count=2 --keyring-backend=test --minimum-gas-prices=0.000001stake --commit-timeout=900ms --single-host
# $SIMD testnet init-files --chain-id=testing --output-dir=./testnet --validator-count=4 --keyring-backend=test --commit-timeout=900ms --single-host


# $SIMD start --log_level=info --home /Users/hieuvu/Documents/Decentrio/cosmos-sdk-decentrio/testnet/node0/simdv2
$SIMD start --log_level=info --home /Users/hieuvu/Documents/Decentrio/cosmos-sdk-decentrio/testnet/node2/simdv2
# $SIMD start --log_level=info --home /Users/hieuvu/Documents/Decentrio/cosmos-sdk-decentrio/testnet/node2/simdv2
# $SIMD start --log_level=info --home /Users/hieuvu/Documents/Decentrio/cosmos-sdk-decentrio/testnet/node3/simdv2

# cd "$CONFIG"

# # to change the voting_period
# jq '.app_state.gov.voting_params.voting_period = "600s"' genesis.json > temp.json && mv temp.json genesis.json

# # to change the inflation
# jq '.app_state.mint.minter.inflation = "0.300000000000000000"' genesis.json > temp.json && mv temp.json genesis.json

# $SIMD config set client chain-id simapp-v2-chain
# $SIMD keys add test_validator --indiscreet
# VALIDATOR_ADDRESS=$($SIMD keys show test_validator -a --keyring-backend test)

# $SIMD genesis add-genesis-account "$VALIDATOR_ADDRESS" 1000000000stake
# $SIMD genesis gentx test_validator 1000000000stake --keyring-backend test
# $SIMD genesis collect-gentxs

# $SIMD start &
# SIMD_PID=$!

# cnt=0
# while ! $SIMD query block --type=height 5; do
#   cnt=$((cnt + 1))
#   if [ $cnt -gt 30 ]; then
#     kill -9 "$SIMD_PID"
#     exit 1
#   fi
#   sleep 1
# done

# kill -9 "$SIMD_PID"
