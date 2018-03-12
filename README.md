# GCoin :heavy_dollar_sign:
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

## Installation
```
go get -u github.com/superstas/gcoin/cmd/...
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

## TODO
- Beauty BE
- Doc
- Fill out README
- Add badges
- Linters
- Tests
- Publish with v0.1
