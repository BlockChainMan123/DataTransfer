# Code description for DataFly
## Code directory structure
| name       | Clarification                  |
| --------------- | ---------------------- |
|ISE	| The implementation of ISE scheme is presented in the folder of ISE|
| Hyperledger_Extract | The implementation of the extract phase of our DataFly is presented in extractPhase. This folder contains all executable code for both plaintext and ciphertext read and write operations in Hyperledger Fabric, including the chaincode layer, SDK-go layer, API layer, and so on.|
| Pri-ETH_Transfer_Load   |  The implementation of the load phase of our DataFly is presented in the folder of Pri-ETH_Transger_Load|


### Introduction to extractPhase

- `chaincode`: writes the underlying chaining code to realize the up-linking operation of adding, deleting, changing and checking the underlying plaintext and ciphertext.
- `fixtures`: contains configuration files written by docker and docker-compose.
  In fixtures, the network comprised two organizations
  (Org#1 and Org#2), each contributing two peers (mocked up
  as a docker image). The ordering service, run by a third-party
  (Org#3), followed an endorsement policy requiring at least
  one peer commitment from each organization for successful
  transaction commitment. The Ordering Service Node (OSN)
  operates in solo mode with default parameters.
- `vendor`: defaults configuration files, including the go package, Hyperledger Fabric CA, 
Hyperledger Fabric Go, and other executable environments.
- `API`: encapsulates API interface functions for reading and writing on-chain data.

#### Environmental settings

```bash
Install the go environment
```
```bash
Install docker and docker-compose
```
generate docker images under the fixtures/docker-compose.yml
```bash
$ docker-compose -f docker-compose.yml up -d 
```
```bash
execute Makefile to initialize the program
```

