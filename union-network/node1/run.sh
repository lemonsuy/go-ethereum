#!/bin/bash
echo "Starting node1 ..."

../../bkc/build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock 'C75f1B2698cF9241988bB162C87E0C33d8EEA2Da' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock \
  # --miner.sealerAddress 'C75f1B2698cF9241988bB162C87E0C33d8EEA2Da' \
  # --miner.etherbase '0xB23068c6412CBE7bd94db54A5176Cc23222B3356'

# geth --datadir . \
#   --config ./config.toml \
#   --unlock 'C75f1B2698cF9241988bB162C87E0C33d8EEA2Da' \
#   --password ./password.txt \
#   --nodekey ./nodekey \
#   --mine \
#   --allow-insecure-unlock
