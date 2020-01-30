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

BenchmarkNew/1_weights-8                 129896          8881 ns/op
BenchmarkNew/10_weights-8                125815          9297 ns/op
BenchmarkNew/100_weights-8               106694         10781 ns/op
BenchmarkNew/1000_weights-8               50884         23185 ns/op
BenchmarkNew/10000_weights-8               6634        152428 ns/op
BenchmarkNew/100000_weights-8               698       1540491 ns/op
BenchmarkNew/1000000_weights-8               75      15374675 ns/op
BenchmarkNew/10000000_weights-8               7     161330317 ns/op

BenchmarkPick/from_1_weights-8         81434600          13.9 ns/op
BenchmarkPick/from_10_weights-8        42723463          26.0 ns/op
BenchmarkPick/from_100_weights-8       43014874          26.1 ns/op
BenchmarkPick/from_1000_weights-8      41333022          25.8 ns/op
BenchmarkPick/from_10000_weights-8     39687285          26.8 ns/op
BenchmarkPick/from_100000_weights-8    32212609          31.7 ns/op
BenchmarkPick/from_1000000_weights-8   18433252          68.7 ns/op
BenchmarkPick/from_10000000_weights-8  14053593          83.8 ns/op
```

## Usage

For usage, examples and documentation, see [GoDoc](https://godoc.org/github.com/minaguib/weightedrandom)
