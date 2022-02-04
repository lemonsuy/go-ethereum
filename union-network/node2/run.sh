#!/bin/bash
echo "Starting node2 ..."

../../build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock '13ec15a841332960AAb03306484a0B0903479650' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock