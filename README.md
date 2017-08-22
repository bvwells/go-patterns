# go-patterns

[![Build Status](https://travis-ci.org/bvwells/go-patterns.svg?branch=master)](https://travis-ci.org/bvwells/go-patterns) 
[![GoDoc](http://godoc.org/github.com/bvwells/go-patterns?status.svg)](http://godoc.org/github.com/bvwells/go-patterns)

Design patterns for the Go programming language.

``` go
import "github.com/bvwells/go-patterns"
```

To install the packages on your system,

```
$ go get -u github.com/bvwells/go-patterns/...
```

Documentation and examples are available at http://godoc.org/github.com/bvwells/go-patterns

 * [Design Patterns](#design-patterns)
 * [Creational](#creational)
 * [Structural](#structural)
 * [Behavioral](#behavioral)
 * [Go Versions Supported](#go-versions-supported)

## Design Patterns

Pattern    | Package                                   | Description
-----------|-------------------------------------------|------------
Creational | [`github.com/bvwells/go-patterns/creational`][creational-ref] | 
Structural | [`github.com/bvwells/go-patterns/structural`][structural-ref] | 
Behavioral | [`github.com/bvwells/go-patterns/behavioral`][behavioral-ref] | 

## Creational [![GoDoc](https://godoc.org/github.com/bvwells/go-patterns/creational?status.svg)](https://godoc.org/github.com/bvwells/go-patterns/creational)

Name       | Description
-----------|-------------------------------------------
Abstract Factory | Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
Builder | Separate the construction of a complex object from its representation, allowing the same construction process to create various representations.
Factory Method | Define an interface for creating a single object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.
Object Pool | Avoid expensive acquisition and release of resources by recycling objects that are no longer in use. Can be considered a generalisation of connection pool and thread pool patterns.
Prototype | Specify the kinds of objects to create using a prototypical instance, and create new objects from the 'skeleton' of an existing object, thus boosting performance and keeping memory footprints to a minimum.
Singleton | Ensure a class has only one instance, and provide a global point of access to it.

## Structural [![GoDoc](https://godoc.org/github.com/bvwells/go-patterns/structural?status.svg)](https://godoc.org/github.com/bvwells/go-patterns/structutal)

Name       | Description
-----------|-------------------------------------------
Adaptor |
Bridge |
Composite |
Decorator |
Facade |
Flyweight |
Proxy |

## Behavioral [![GoDoc](https://godoc.org/github.com/bvwells/go-patterns/behavioral?status.svg)](https://godoc.org/github.com/bvwells/go-patterns/behavioral)

Name       | Description
-----------|-------------------------------------------
Chain of Responsibility |
Command |
Interpretor |
Iterator |
Mediator |
Memento |
Null Object |
Observer |
State | 
Strategy |
Template Method |
Visitor |

## Go Versions Supported

The most recent major version of Go is supported. You can see which versions are
currently supported by looking at the lines following `go:` in
[`.travis.yml`](.travis.yml).

[creational-ref]: https://godoc.org/github.com/bvwells/go-patterns/creational
[structural-ref]: https://godoc.org/github.com/bvwells/go-patterns/structural
[behavioral-ref]: https://godoc.org/github.com/bvwells/go-patterns/behavioral
