echo "Initial genesis file to [node 1]"
../../build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/node1 /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/genesis.json

echo "Initial genesis file to [node 2]"
../../build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/node2 /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/genesis.json

echo "Initial genesis file to [node 3]"
../../build/bin/geth init --datadir /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/node3 /Users/lemon.pattharathon/lemon-research/go-ethereum/union-network/genesis.json

sleep 1;
echo "Done!"
