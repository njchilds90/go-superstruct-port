# go-superstruct-port

A Go port of the Python library 'superstruct', providing a simple way to define and validate structured data.

## Overview

This library allows you to define a structure for your data using a simple and declarative API. You can then use this structure to validate and coerce your data into the desired shape.

## Installation

To install the library, run the following command:

```go

go get github.com/go-superstruct-port/superstruct
```

## Usage

Here is an example of how to use the library:

    package main

    import (
        "fmt"
        "github.com/go-superstruct-port/superstruct"
    )

    func main() {
        // Define a structure for the data
        type User struct {
            Name  string `struct:"name"`
            Email string `struct:"email"`
        }

        // Create a new instance of the structure
        user := User{
            Name:  "John Doe",
            Email: "john@example.com",
        }

        // Validate the data
        err := superstruct.Validate(user)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Data is valid")
        }
    }

## API Reference

### type Struct

* `func NewStruct(name string) *Struct`
* `func (s *Struct) AddField(name string, typ reflect.Type) *Struct`
* `func (s *Struct) Validate(data interface{}) error`

### type Field

* `func NewField(name string, typ reflect.Type) *Field`
* `func (f *Field) SetTag(tag string) *Field`

## Contributing

Contributions are welcome! Please submit a pull request with your changes and a brief description of what you've added or fixed.