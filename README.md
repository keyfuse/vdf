# Verifiable Delay Function (VDF) Implementation in Golang.

[![Build Status](https://travis-ci.org/keyfuse/vdf.png)](https://travis-ci.org/keyfuse/vdf) [![Go Report Card](https://goreportcard.com/badge/github.com/keyfuse/vdf)](https://goreportcard.com/report/github.com/keyfuse/vdf)

## What is a VDF?

Verifiable Delay Functions take a prescribed time to compute, even on a parallel computer, yet produce a unique output that can be efficiently and publicly verified.

## Description

We implement the approaches described in the following papers:
1. [Sloth(Quadratic Residue) Verifiable Delay Functions](https://eprint.iacr.org/2018/601.pdf). Boneh, 2018
2. [Simple Verifiable Delay Functions](https://eprint.iacr.org/2018/627.pdf). Pietrzak, 2018
3. [Efficient Verifiable Delay Functions](https://eprint.iacr.org/2018/623.pdf). Wesolowski, 2018


## How to Build

To build  from the source code you need to have a working
Go environment with [version 1.12 or greater installed](https://golang.org/doc/install).

```
$ git clone https://github.com/keyfuse/vdf
$ cd vdf
$ make build
```

## How to Run

```
./bin/vdf -x=7 -t=1000
T:1000, x:7, P:8740186361822062383
-->
Type:	Sloth Quadratic Residue
Value:	37602270514657052337634791495305873950174757777621818362566683935496339681801
Delay:	0.10983 secs
Verify:	0.00064 secs
Verify Result:	true
Delay/Verify:	170.32087
```

## License

VDF is released under the GPLv3 License.
