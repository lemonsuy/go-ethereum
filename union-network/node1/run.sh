#!/bin/bash
echo "Starting node1 ..."

../../build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock 'C75f1B2698cF9241988bB162C87E0C33d8EEA2Da' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock