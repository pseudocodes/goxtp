# goxtp
Golang wrapper bindings for XTP.

## Why Golang
> https://golang.org/doc/
The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient.
Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines,
while its novel type system enables flexible and modular program construction.
Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection.
It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

## Why Swig
> http://www.swig.org/Doc3.0/Go.html#Go
Go does not support direct calling of functions written in C/C++.
The cgo program may be used to generate wrappers to call C code from Go, but there is no convenient way to call C++ code.
SWIG fills this gap.

## Why goXTP
This is a matter of course.

## How to Build goXTP
> Only support Linux X64 now.

```
go get -u -v github.com/pseudocodes/goxtp

./build.sh
```

## How to Build example

```
make example
```