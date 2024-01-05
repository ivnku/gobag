<p align="center">
    <img src="gobag.png" alt="GoBag logo">
</p>

<h1 align="center">GoBag</h1>

The GoBag library is a set of helpers and types which is basically just a bunch of syntax sugar for Golang.

Pull requests are highly appreciated, so if you want to add something - feel free to create it.

## Getting Started

### Prerequisites

For this library to work you need a Golang version 1.18 or higher

### Installing

To install GoBag just run

```
go get github.com/ivnku/gobag
```

### Documentation

* [Collections](#Collections)
    * [Filter](#Filter)
    * [GroupBy](#GroupBy)
    * [Map](#Map)
    * [Reduce](#Reduce)
* [Maps](#Collections)
    * [GetKeys](#GetKeys)
    * [GetValues](#GetValues)
    * [Reform](#Reform)

### Collections

#### Filter

Filters collection of type `A` by condition `cond()`

```go
Filter[A any](in []A, cond func (A) bool) []A
```

#### GroupBy

Groups a collection of type `I` by the key of type `K` which is formed by the `keyFn(I) K` function and with the value
of type `[]V` each item of which is formed by the `valueFn(I) V` function

```go
GroupBy[I any, K comparable, V any](
    in []I,
    keyFn func (I) K,
    valueFn func (I) V,
) map[K][]V
```

#### Map

Transforms a collection of the type `A` into the collection of the type `B` with the mapping function `fn(A) B`

```go
Map[A any, B any](in []A, fn func (A) B) []B
```

#### Reduce

Accumulates the data from a collection of the type `[]A` with the function `fn func(acc B, item A) B`

```go
Reduce[A any, B any](in []A, acc B, fn func (acc B, item A) B) B
```

### Maps

#### GetKeys

Returns a slice of keys of a map

```go
GetKeys[K comparable, V any](in map[K]V) []K
```

#### GetValues

Returns a slice of values of a map

```go
GetValues[K comparable, V any](in map[K]V) []V
```

#### Reform

Reformat map's keys and values from one type to another

```go
Reform[K1 comparable, K2 comparable, V1 any, V2 any](
    in map[K1]V1,
    keyFn func(K1) K2,
    valFn func(V1) V2,
) map[K2]V2
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
