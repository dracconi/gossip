# gossip
Peer-to-peer chat written in Go

[![Github All Releases](https://img.shields.io/github/downloads/dracconi/gossip/total.svg)]()
[![GitHub license](https://img.shields.io/github/license/dracconi/gossip.svg)](https://github.com/dracconi/gossip/blob/master/LICENSE)

[![GitHub stars](https://img.shields.io/github/stars/dracconi/gossip.svg)](https://github.com/dracconi/gossip/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracconi/gossip)](https://goreportcard.com/report/github.com/dracconi/gossip)


## How to compile

[Install Go compiler](https://golang.org/doc/install).

`go get github.com/dracconi/gossip` (Note: master is currently unstable branch that doesn't always work!)

`go install github.com/dracconi/gossip`

It's depends on only:
* tui-go

## Usage

If you have working executable, then turn on the server by `gossip s`, you may specifiy port with `--port [port]`. Then start the client `gossip c [ip address with port]`.