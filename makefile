# https://ethdocs.org/en/latest/contracts-and-transactions/accessing-contracts-and-transactions.html
# https://goethereumbook.org/smart-contract-deploy/
# https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J#dd57ef90-f990-037e-5512-4929e7280d7c
#
# The environment has three accounts all using this same passkey (123).
# Geth is started with address 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd and is used as the coinbase address.
# The coinbase address is the account to pay mining rewards to.
# The coinbase address is give a LOT of money to start.
#
# These are examples of what you can do in the attach JS environment.
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd") or eth.getBalance(eth.coinbase)
# 	eth.getBalance("0x8e113078adf6888b7ba84967f299f29aece24c55")
# 	eth.getBalance("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
#   eth.sendTransaction({from:eth.coinbase, to:"0x8e113078adf6888b7ba84967f299f29aece24c55", value: web3.toWei(0.05, "ether")})
#   eth.sendTransaction({from:eth.coinbase, to:"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", value: web3.toWei(0.05, "ether")})
#   eth.blockNumber
#   eth.getBlockByNumber(8)
#   eth.getTransaction("0xaea41e7c13a7ea627169c74ade4d5ea86664ff1f740cd90e499f3f842656d4ad")
#
# check this out https://etherscan.io/gastracker
# If you want something to go through soon, have a base fee (~30) and the priority fee or tip of 1 gwei
# Priority fee doesn't do much i've noticed and is always around 1-3 gwei. I've been setting that base fee higher.
#
#  Unit	                Wei Value	 Wei
#  wei	                1 wei        1
#  Kwei (babbage)	    1e3 wei	     1,000
#  Mwei (lovelace)	    1e6 wei	     1,000,000
#  Gwei (shannon)	    1e9 wei	     1,000,000,000
#  microether (szabo)	1e12 wei	 1,000,000,000,000
#  milliether (finney)	1e15 wei	 1,000,000,000,000,000
#  ether	            1e18 wei	 1,000,000,000,000,000,000
#
# Visibility Quantifiers
# external − External functions are meant to be called by other contracts. They cannot be used for internal calls.
# public   − Public functions/variables can be used both externally and internally. For public state variables, Solidity automatically creates a getter function.
# internal − Internal functions/variables can only be used internally or by derived contracts.
# private  − Private functions/variables can only be used internally and not even by derived contracts.
#
# Variable Location Options
# Storage  - It is where all state variables are stored. Because state can be altered in a contract (for example, within a function), storage variables must be mutable. However, their location is persistent, and they are stored on the blockchain.
# Memory   - Reserved for variables that are defined within the scope of a function. They only persist while a function is called, and thus are temporary variables that cannot be accessed outside this scope (ie anywhere else in your contract besides within that function). However, they are mutable within that function.
# Calldata - Is an immutable, temporary location where function arguments are stored, and behaves mostly like memory.
#

# ==============================================================================
# Install dependencies
# https://geth.ethereum.org/docs/install-and-build/installing-geth
# https://docs.soliditylang.org/en/v0.8.17/installing-solidity.html

dev.setup:
	brew update
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity

dev.update:
	brew update
	brew list ethereum || brew upgrade ethereum
	brew list solidity || brew upgrade solidity

# ==============================================================================
# These commands build, deploy, and run the basic smart contract.

# This will compile the smart contract and produce the binary code. Then with the
# abi and binary code, a Go source code file can be generated for Go API access.

basic-build:
	solc --abi app/basic/contract/src/basic/basic.sol -o app/basic/contract/abi/basic --overwrite
	solc --bin app/basic/contract/src/basic/basic.sol -o app/basic/contract/abi/basic --overwrite
	abigen --bin=app/basic/contract/abi/basic/Basic.bin --abi=app/basic/contract/abi/basic/Basic.abi --pkg=basic --out=app/basic/contract/go/basic/basic.go

# This will deploy the smart contract to the locally running Ethereum environment.
basic-deploy:
	CGO_ENABLED=0 go run app/basic/cmd/deploy/main.go

# This will execute a simple program to test access to the smart contract API.
basic-write:
	CGO_ENABLED=0 go run app/basic/cmd/write/main.go

# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

# This is start Ethereum in developer mode. Only when a transaction is pending will
# Ethereum mine a block. It provides a minimal environment for development.
geth-up:
	geth --dev --ipcpath zarf/ethereum/geth.ipc --http.corsdomain '*' --http --allow-insecure-unlock --rpc.allow-unprotected-txs --mine --miner.threads 1 --verbosity 5 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

# This will signal Ethereum to shutdown.
geth-down:
	kill -INT $(shell ps -eo pid,comm | grep " geth" | awk '{print $$1}')

# This will remove the local blockchain and let you start new.
geth-reset:
	rm -rf zarf/ethereum/geth/

# This is a JS console environment for making geth related API calls.
geth-attach:
	geth attach --datadir zarf/ethereum/

# This will add a new account to the keystore. The account will have a zero
# balance until you give it some money.
geth-new-account:
	geth --datadir zarf/ethereum/ account new

# This will deposit 1 ETH into the two extra accounts from the coinbase account.
# Do this if you delete the geth folder and start over or if the accounts need money.
geth-deposit:
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x8E113078ADF6888B7ba84967F299F29AeCe24c55", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x7FDFc99999f1760e8dBd75a480B93c7B8386B79a", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x000cF95cB5Eb168F57D0bEFcdf6A201e3E1acea9", "value":"0x1000000000000000000"}], "id":1}' localhost:8545

# ==============================================================================
# These commands provide Go related support.

test:
	CGO_ENABLED=0 go test -count=1 ./...
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...
	govulncheck ./...

# This will tidy up the Go dependencies.
tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor