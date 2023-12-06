# Example DApp using EggRoll

Minimum template for [EggRoll](https://github.com/gligneul/eggroll/) projects.

## Generate schema

```sh
go generate
```

## Build

```sh
sunodo build
```

## Run

```sh
sunodo run
```

## Send input

```sh
ARGS='{"value": "hi"}' \
INPUT=$(eggroll schema encode --kind advanceEcho --args $ARGS) ; \
sunodo send generic \
    --mnemonic-index=0 \
    --mnemonic-passphrase="test test test test test test test test test test test junk" \
    --chain-id=31337 \
    --rpc-url=http://localhost:8545 \
    --dapp=0x70ac08179605AF2D9e75782b8DEcDD3c22aA4D0C \
    --input=$INPUT \
    --input-encoding=hex
```
