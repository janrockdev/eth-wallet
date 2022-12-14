# Eth-Wallet Generator

Simple Ethereum wallet generator with a RESTful API server to interact with ethereum type of network, Etherscan and node. 

ToDo & Bugs
--------
- fix send transactions 
- option for Etherscan (Rinkeby testnet)

Features
--------
- create keystore with a wallet details
- get address's balance(eth, erc20), transaction, nonce from node
- estimate transaction's gas limit
- get gasprice from node
- get address's transaction history(eth, erc20) from etherscan api
- send eth, erc20 from your address to another address
- sign message, sign transaction
- verify message

Setup
------

Create account with Infura (Rinkeby Testnet API):

https://infura.io/dashboard

Top-up Rinkeby Faucet:

https://rinkebyfaucet.com

Etherscan Rinkeby:

https://rinkeby.etherscan.io

## Setup config

### config path

Environment variable: ETHEREUM_WALLET_CONFIG_PATH

```shell
export ETHEREUM_WALLET_CONFIG_PATH="/Users/janrock/eth-wallet/tmp"
```
(if you don't set path to config.json, default path be project root)

### config.json

#### Variable

- `network(not necessary)`: defalut network, you can choose ropsten, rinkeby or mainnet.  
- `ropsten, rinkeby, mainnet`: network's node_url and etherscan_api_url need to be set in here, `node_url`
  is ethereum node's url, you can choose to use infura(https://infura.io/), `etherscan_api_url` is
  etherscan's developer api url.  
- `etherscan_api_key`: etherscan's api key, you can register at etherscan(https://etherscan.io/apis)
- `server_url(not necessary)`: server's url when use start server command, default is set in http://127.0.0.1:8080  
- `keyfile`: keystore's path, you can create keystore from cli create command  
- `passphrase(not necessary)`: keystore's passphrase, it's a fast way to unlock keyfile, or you can input in terminal when cli need unlock wallet  
- `address(not necessary)`: default query address  
- `erc20_list`: erc20 token's list, you need provide token's decimals, name, symbol, I put some popular 
  erc20 token in project's erc20_list.json file, user can add token you need in config file.

#### Example
```
{
"rinkeby":{
"node_url": "https://rinkeby.infura.io/v3/a827d9b9e1ff4d17b3dc46698821b8cb",
"etherscan_api_url": "https://api-rinkeby.etherscan.io/api"
},
  "server_url": "http://127.0.0.1:8080",
  "etherscan_api_key": "NFZAYYT9CJ1AYFPAZ7HF2134E63U9UUBBCS",
  "keyfile": "./tests/key/test",
  "passphrase": "pass",
  "address": "0xd3ba5c8d3a81f601bbb5e085773725a685139aa7",
  "erc20_list": [
    {
      "decimals": 18,
      "name": "Weenus",
      "symbol": "Weenus",
      "address": "0x10184ad5c5bbaa18e6b4431eedf6b95e9adf82fa"
    }
  ]
}
```

## Build 

```shell script
go build ./cmd/cli
```

## Run

### run server first
```shell script
open /Applications/Google\ Chrome.app --args --user-data-dir="/var/tmp/Chrome dev session" --disable-web-security
./cli server start -network rinkeby
```

### do what you want to do
```shell script
./cli
```

Example
-------

## cli command

### Server command

```shell script
./cli server start -port 8080 -network rinkeby
```

### Wallet command

#### get keystore address

```shell script
./cli wallet address -keyfile "./keystore/test"
```

#### create keystore

```shell script
./cli wallet create -p ""./keystore"
```

#### sign message

```shell script
./cli wallet signmessage --keyfile "./keystore/test" --msg "hello"
```

#### sign transaction

```shell script
./cli wallet signtx --keyfile "./keystore/test" --transaction  "{\"nonce\":160,\"gasprice\":2000000000,\"gaslimit\":21000,\"to\":\"0xe5664b93ad268393d1f695c4180993e60c59fc3e\",\"value\":1000000000000,\"data\":\"\"}"
```

#### verifymessage

```shell script
./cli wallet verifymessage -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B" -signature "0x61c01b1a23624f176cbc42feda9c394ce0c9c8dd80b46ab4ca3d5dfb95a4e60335ec0f8c1bcc475dfc5bdafa697b10e56c329fdf136fee4ec800898be2412d4f00" -message "hello"
```

### Node command

#### get address's balance

```shell script
./cli node balance -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```

#### get address's nonce

```shell script
./cli node nonce -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```

#### get address's erc20balance

```shell script
./cli node erc20balance -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```
#### get address's txhistory

```shell script
./cli node txhistory -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```

#### get address's internaltxhistory

```shell script
./cli node internaltxhistory -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```

#### get address's erc20txhistory

```shell script
./cli node erc20txhistory -address "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B"
```

#### get gasprice from node

```shell script
./cli node gasprice 
```

#### get transaction's estimate gas

- The unit needs to be expressed in the form of wei

```shell script
./cli node gaslimit --transaction ./cli node gaslimit -transaction "{\"from\":\"0xb60e8dd61c5d32be8058bb8eb970870f07233155\",\"to\":\"0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B\",\"gasPrice\":10000000000,\"value\":100000000}"
```
#### send raw transaction

```shell script
./cli node sendrawtx -raw "0x1234"
```

### NodeWallet wallet

#### send ether to other address

- The unit needs to be expressed in the form of wei

```shell script
./cli nodewallet sendether -keyfile "./keystore/test" -to "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B" -value 10000000000
```

### send erc20 to other address

```shell script
./cli nodewallet senderc20 -keyfile "./keystore/test" -symbol "Weenus" -to "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B" -value 1
```
