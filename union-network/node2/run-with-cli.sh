#!/bin/bash
echo "Starting node2 ..."

geth --datadir <DATA_DIR> \
  --identity node2 \
  --syncmode full \
  --port 30002 \
  --ws \
  --ws.addr 0.0.0.0 \
  --ws.port 8556 \
  --ws.origins "*" \
  --http \
  --http.addr 0.0.0.0 \
  --http.port 8555 \
  --http.corsdomain "*" \
  --http.api eth,net,web3,txpool,miner,admin,clique,debug \
  --networkid <NETWORK_ID> \
  --miner.gasprice 0 \
  --unlock '<MINER_ADDRESS>' \
  --password <DATA_DIR>/password.txt \
  --mine \
  --verbosity 3 \
  --allow-insecure-unlock
