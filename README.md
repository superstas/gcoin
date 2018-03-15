# GCoin :heavy_dollar_sign:
[![GoDoc](https://godoc.org/github.com/superstas/gcoin/gcoin?status.svg)](http://godoc.org/github.com/superstas/gcoin/gcoin) [![Build Status](https://travis-ci.org/superstas/gcoin.svg?branch=master)](https://travis-ci.org/superstas/gcoin)  
A very simple proof-of-concept implementation of a PoW cryptocurrency written in Go.  
This has been written as a part of the talk at the [GopherCon Russia 2018](https://gophercon-russia.ru).  


:warning: This is not production ready implementation!  
:warning: Use it for learning purposes only.  

## Features :smiley:
:heavy_check_mark: Simple P2PKH transactions  
:heavy_check_mark: DSHA256 PoW consensus algorithm  
:heavy_check_mark: Block explorer  
:heavy_check_mark: Manual network discovering  
:heavy_check_mark: CLI tool  
:heavy_check_mark: Wallet ( with keys and PKH address generation )

## Demo


## Requirements
Go 1.9 or newer  

## Installation
```
go get -u github.com/superstas/gcoin/cmd/...
```

## Usage
### Daemon
```
$> gcoind -h
NAME:

                                                                                                dddddddd
        GGGGGGGGGGGGG        CCCCCCCCCCCCC                   iiii                               d::::::d
     GGG::::::::::::G     CCC::::::::::::C                  i::::i                              d::::::d
   GG:::::::::::::::G   CC:::::::::::::::C                   iiii                               d::::::d
  G:::::GGGGGGGG::::G  C:::::CCCCCCCC::::C                                                      d:::::d
 G:::::G       GGGGGG C:::::C       CCCCCC   ooooooooooo   iiiiiiinnnn  nnnnnnnn        ddddddddd:::::d
G:::::G              C:::::C               oo:::::::::::oo i:::::in:::nn::::::::nn    dd::::::::::::::d
G:::::G              C:::::C              o:::::::::::::::o i::::in::::::::::::::nn  d::::::::::::::::d
G:::::G    GGGGGGGGGGC:::::C              o:::::ooooo:::::o i::::inn:::::::::::::::nd:::::::ddddd:::::d
G:::::G    G::::::::GC:::::C              o::::o     o::::o i::::i  n:::::nnnn:::::nd::::::d    d:::::d
G:::::G    GGGGG::::GC:::::C              o::::o     o::::o i::::i  n::::n    n::::nd:::::d     d:::::d
G:::::G        G::::GC:::::C              o::::o     o::::o i::::i  n::::n    n::::nd:::::d     d:::::d
 G:::::G       G::::G C:::::C       CCCCCCo::::o     o::::o i::::i  n::::n    n::::nd:::::d     d:::::d
  G:::::GGGGGGGG::::G  C:::::CCCCCCCC::::Co:::::ooooo:::::oi::::::i n::::n    n::::nd::::::ddddd::::::dd
   GG:::::::::::::::G   CC:::::::::::::::Co:::::::::::::::oi::::::i n::::n    n::::n d:::::::::::::::::d
     GGG::::::GGG:::G     CCC::::::::::::C oo:::::::::::oo i::::::i n::::n    n::::n  d:::::::::ddd::::d
        GGGGGG   GGGG        CCCCCCCCCCCCC   ooooooooooo   iiiiiiii nnnnnn    nnnnnn   ddddddddd   ddddd

gcoin-daemon - Run Gcoin as a daemon

USAGE:
   gcoind [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --add-node value       this flag used to connect to an another node
   --listen value         this flag starts the node on given listen address (default: "0.0.0.0:10000")
   --http-listen value    this flag starts a block explorer on given listen address
   --miner-address value  an address for mining reward ( for mining mode only ) (default: "12PLGoQb9usohESQAeB6rrYXPcua9M36Pc")
   --run-mining           this flag enables mining mode
   --help, -h             show help
   --version, -v          print the version

```
### Client
```
$> gcoin-cli -h
NAME:

        GGGGGGGGGGGGG        CCCCCCCCCCCCC                   iiii
     GGG::::::::::::G     CCC::::::::::::C                  i::::i
   GG:::::::::::::::G   CC:::::::::::::::C                   iiii
  G:::::GGGGGGGG::::G  C:::::CCCCCCCC::::C
 G:::::G       GGGGGG C:::::C       CCCCCC   ooooooooooo   iiiiiii nnnn  nnnnnnnn
G:::::G              C:::::C               oo:::::::::::oo i:::::i n:::nn::::::::nn
G:::::G              C:::::C              o:::::::::::::::o i::::i n::::::::::::::nn
G:::::G    GGGGGGGGGGC:::::C              o:::::ooooo:::::o i::::i nn:::::::::::::::n
G:::::G    G::::::::GC:::::C              o::::o     o::::o i::::i   n:::::nnnn:::::n
G:::::G    GGGGG::::GC:::::C              o::::o     o::::o i::::i   n::::n    n::::n
G:::::G        G::::GC:::::C              o::::o     o::::o i::::i   n::::n    n::::n
 G:::::G       G::::G C:::::C       CCCCCCo::::o     o::::o i::::i   n::::n    n::::n
  G:::::GGGGGGGG::::G  C:::::CCCCCCCC::::Co:::::ooooo:::::oi::::::i  n::::n    n::::n
   GG:::::::::::::::G   CC:::::::::::::::Co:::::::::::::::oi::::::i  n::::n    n::::n
     GGG::::::GGG:::G     CCC::::::::::::C oo:::::::::::oo i::::::i  n::::n    n::::n
        GGGGGG   GGGG        CCCCCCCCCCCCC   ooooooooooo   iiiiiiii  nnnnnn    nnnnnn

gcoin cli tool - command-line client for gcoin

USAGE:
   gcoin-cli [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
     send, s            this command sends coins
     getbalance, gblnc  this command returns all known addresses with UTXOs
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --node-addr value  an hostname:port of running node (default: "localhost:10000")
   --help, -h         show help
   --version, -v      print the version
```
## Run demo :arrow_forward:


## Limitations :disappointed:
There are no following things:  
:heavy_minus_sign: A syncing protocol between nodes  
:heavy_minus_sign: Blockchain forking and conflicts resolving  
:heavy_minus_sign: A permanent blockchain storage  
:heavy_minus_sign: A real P2P discovering  
:heavy_minus_sign: An internal locking/unlocking mechanism  


## Credits :+1:
- https://github.com/btcsuite - A very helpful set of packages
- https://github.com/urfave/cli - A library that helps a lot in creating CLI tools
- https://github.com/olekukonko/tablewriter - ASCII tables in Go :fire:
- https://en.bitcoin.it/wiki/Main_Page - The main source of the information about BitCoin
- https://bitcoin.stackexchange.com - The most important source to find out some hidden details
- https://jeiwan.cc/posts - Amazing articles about BitCoin details and implementation in Go
- https://github.com/bitcoinbook/bitcoinbook - If you want to know more, read this book