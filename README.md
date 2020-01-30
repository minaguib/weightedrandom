# minaguib/weightedrandom

A Go (golang) library for efficient weighted random picking

[![GoDoc](https://godoc.org/github.com/minaguib/weightedrandom?status.svg)](https://godoc.org/github.com/minaguib/weightedrandom)
[![Build Status](https://travis-ci.org/minaguib/weightedrandom.svg)](https://travis-ci.org/minaguib/weightedrandom)
[![Coverage Status](https://coveralls.io/repos/github/minaguib/weightedrandom/badge.svg?branch=master)](https://coveralls.io/github/minaguib/weightedrandom?branch=master)
[![GoReport](https://goreportcard.com/badge/github.com/minaguib/weightedrandom)](https://goreportcard.com/report/github.com/minaguib/weightedrandom)

## About

Given a theoretical jar of 1,000 marbles, like so:

|Marble color|Count|
|---|---|
|Red|500|
|Blue|250|
|Green|125|
|Yellow|120|
|Transparent|4|
|Vantablack|1|

You want a solution that allows you to simulate picking a random marble while adhering to the above probabilities.  You want it to operate as fast as possible, require as little storage as possible, and efficiently handle large weights and large number of items.


## About the implementation

This library implements the Alias picking method as further optimized by Michael D. Vose

Ref. Paper: "[A Linear Algorithm For Generating Random Numbers With a Given Distribution](https://pdfs.semanticscholar.org/f65b/cde1fcf82e05388b31de80cba10bf65acc07.pdf)" by Michael D. Vose

Ref: [Darts, Dice, and Coins: Sampling from a Discrete Distribution](https://www.keithschwarz.com/darts-dice-coins/) by Keith Schwarz

## Performance characteristics

|Item|Cost|
|---|---|
|Time to initialize|O(n)|
|Data structure storage|O(n)|
|Time to pick an item|O(1)|

## Performance benchmarks

On a MacBook Pro, Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz, 2133 MHz LPDDR3 memory:

```
$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/minaguib/weightedrandom

BenchmarkNew/1_weights-8           126366      8967 ns/op
BenchmarkNew/10_weights-8          126315      9015 ns/op
BenchmarkNew/100_weights-8         112978      9985 ns/op
BenchmarkNew/1000_weights-8         58854     19891 ns/op
BenchmarkNew/10000_weights-8         9981    109303 ns/op
BenchmarkNew/100000_weights-8        1029   1001077 ns/op
BenchmarkNew/1000000_weights-8        112   9697761 ns/op
BenchmarkNew/10000000_weights-8        12  89310961 ns/op

BenchmarkPick/from_1_weights-8            76958226        13.8 ns/op
BenchmarkPick/from_10_weights-8           42556176        26.2 ns/op
BenchmarkPick/from_100_weights-8          43152367        26.0 ns/op
BenchmarkPick/from_1000_weights-8         42600124        26.0 ns/op
BenchmarkPick/from_10000_weights-8        39454466        27.1 ns/op
BenchmarkPick/from_100000_weights-8       31104279        32.5 ns/op
BenchmarkPick/from_1000000_weights-8      17665467        69.6 ns/op
BenchmarkPick/from_10000000_weights-8     11770742       104 ns/op
```

## Usage

For usage, examples and documentation, see [GoDoc](https://godoc.org/github.com/minaguib/weightedrandom)
