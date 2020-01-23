# minaguib/weightedrandom

A Go (golang) library for efficient weighted random picking

[![GoDoc](https://godoc.org/github.com/minaguib/weightedrandom?status.svg)](https://godoc.org/github.com/minaguib/weightedrandom)
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

## Usage

For usage, examples and documentation, see [GoDoc](https://godoc.org/github.com/minaguib/weightedrandom)
