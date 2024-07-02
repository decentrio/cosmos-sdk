#!/usr/bin/env bash

set -o errexit
set -o nounset
set -x

ROOT=$PWD

SIMD="$ROOT/build/simdv2"
CONFIG="${CONFIG:-$HOME/.simappv2/config}"

COSMOS_BUILD_OPTIONS=v2 make build     

if [ -d "$($SIMD config home)" ]; then rm -rv $($SIMD config home); fi

$SIMD cometbft-server status --node tcp://127.0.0.1:26657

$SIMD cometbft-server status --node tcp://127.0.0.1:26658

$SIMD keys list --home ./testnet/node0/simdv2

$SIMD keys list --home ./testnet/node1/simdv2

$SIMD keys list --home ./testnet/node2/simdv2


