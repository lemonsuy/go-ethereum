#!/bin/bash
echo "Starting node2 ..."

../../bkc/build/bin/geth --datadir . \
  --config ./config.toml \
  --unlock '13ec15a841332960AAb03306484a0B0903479650' \
  --password ./password.txt \
  --nodekey ./nodekey \
  --mine \
  --allow-insecure-unlock \
  # --miner.sealerAddress '13ec15a841332960AAb03306484a0B0903479650' \
  # --miner.etherbase '0x57054419a0d667C74c3404F52119344ec018021A'

# geth --datadir . \
#   --config ./config.toml \
#   --unlock '13ec15a841332960AAb03306484a0B0903479650' \
#   --password ./password.txt \
#   --nodekey ./nodekey \
#   --mine \
#   --allow-insecure-unlock