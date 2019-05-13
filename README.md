# dbender
[![Go Report Card](https://goreportcard.com/badge/github.com/cohenjo/dbender)](https://goreportcard.com/report/github.com/cohenjo/dbender)
[![GoDoc](https://godoc.org/github.com/cohenjo/dbender?status.svg)](https://godoc.org/github.com/cohenjo/dbender)
[![build status](https://travis-ci.org/cohenjo/dbender.svg)](https://travis-ci.org/cohenjo/dbender) [![downloads](https://img.shields.io/github/downloads/cohenjo/dbender/total.svg)](https://github.com/cohenjo/dbender/releases) 
[![release](https://img.shields.io/github/release/cohenjo/dbender.svg)](https://github.com/cohenjo/dbender/releases)

# Db Bot ENtity Doing Emergency Routines
(or just a bot to end your db ;) )

give it some conf, spin it have and have a nice conversation about db stuff 


# The simple slack messanger relay
simple code with grpc to relay grpc to slack...
example:
```bash
go generate github.com/cohenjo/dbender/pkg/messanger/...
go run cmd/messanger/main.go
grpc_cli call localhost:50051 SendMessage "msg: 'test', body: 'life', channel: '@cohenjo'
```


## Bunch of things it will try to do:
- [ ] a report file on your cluster
- [ ] locking report
- [ ] get useful info
- [ ] understand what the heck are you talking about
- [ ] be active part of your slack channels
- [ ] be mildely offensive


- [ ] external hooks?
- [ ] volume sizes (resizing?)
- [ ] operatoion with locks/long running
- [ ] call huston to annoy anna
- [ ] check server resource cpu/memory/iops
