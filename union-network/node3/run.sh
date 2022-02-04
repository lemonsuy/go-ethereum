#!/bin/bash
echo "Starting node3 ..."

../../build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock '1b9526fe6607e4f9147b9c8137eb4740b0888b9c' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock