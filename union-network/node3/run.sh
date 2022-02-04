#!/bin/bash
echo "Starting node3 ..."

../../bkc/build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock '1b9526fe6607e4f9147b9c8137eb4740b0888b9c' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock \
  # --miner.sealerAddress '1b9526fe6607e4f9147b9c8137eb4740b0888b9c' \
  # --miner.etherbase '0x000BDFE0235eAb68af68aa43658edd0F5eFFab79'

# geth --datadir . \
#   --config ./config.toml \
#   --unlock '1b9526fe6607e4f9147b9c8137eb4740b0888b9c' \
#   --password ./password.txt \
#   --nodekey ./nodekey \
#   --mine \
#   --allow-insecure-unlock