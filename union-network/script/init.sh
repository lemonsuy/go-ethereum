echo "Initial genesis file to [node 1]"
../../bkc/build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/node1 /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/genesis.json

echo "Initial genesis file to [node 2]"
../../bkc/build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/node2 /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/genesis.json

echo "Initial genesis file to [node 3]"
../../bkc/build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/node3 /Users/lemon.pattharathon/lemon-fork-chain/bkc-posa-network/genesis.json

sleep 1;
echo "Done!"
