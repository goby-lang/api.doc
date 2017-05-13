# Goby API Documentation (WIP)

[![Code Climate](https://codeclimate.com/github/goby-lang/api.doc/badges/gpa.svg)](https://codeclimate.com/github/goby-lang/api.doc)

This is the API documentation for goby language. 

The doc is hosted on: https://goby-lang.github.io/api.doc/

## Contributing to Parser

Install dependency:

```plain
go get github.com/russross/blackfriday
```

Run in project root:

```plain
go run *.go
```

That should generate new docs in `/docs` directory. Commit and push them.

## Documenting Code

We use Markdown to document. All markdown syntax are supported.

### Documenting Classes

All classes can be found in `/vm` directory. Each class is represented by a file. For example, you can find `Integer` class in `integer.go` file.

The class definition can be found with the `type` named as the filename or suffixed with "Object". The documentation for this class is right above this `type`. For example, `Integer` class is documented as:

```go
// Integer represents number objects which can bring into mathematical calculations.
//
// ```
// 1 + 1 # => 2
// 2 * 2 # => 4
// ```
type IntegerObject struct {
  Class *RInteger
  Value int
}
```

### Documenting Methods

Methods of a class are stored in an array, with its name prefixed with `builtin` and suffixed with `Methods`. For example, the name for `Integer` class is `builtinIntegerMethods`.

Each method comes with two keys, `Fn` and `Name`. The document of a method comes right above `Fn`. For example:

```go
var builtinIntegerMethods = []*BuiltInMethod{
  {
    // Returns the sum of self and another Integer.
    //
    // ```
    // 1 + 1
    // ```
    Fn: func(receiver Object) builtinMethodBody {
      // ...
    },
    Name: "+",
  },
}
```

Remember to leave a space after `//`.

### Others

Currently the documentation only supports `//` comments, not `/* */`. The following style will not be parsed:

```go
/*
  Integer represents number objects which can bring into mathematical calculations.
*/
type IntegerObject struct {
  Class *RInteger
  Value int
}
```
