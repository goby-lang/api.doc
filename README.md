# Goby API Documentation

[![Code Climate](https://codeclimate.com/github/goby-lang/api.doc/badges/gpa.svg)](https://codeclimate.com/github/goby-lang/api.doc)

This is the API documentation for goby language. 

The doc is hosted on: https://goby-lang.github.io/api.doc/

## Setup & Update Goby Documentation

Follow the steps to setup a running parser.

- Fork & Clone (or sync if you already have a copy [by this instruction](https://help.github.com/articles/syncing-a-fork/))
- Clone [Goby language project](https://github.com/goby-lang/goby). The best way to do it is using Go's `go get` command. By default, this parser looks for `[GOPATH]/src/github.com/goby-lang/goby/vm` directory. If you're new to Golang, make sure you've setup `GOPATH` and clone the project following Go's convention.
- Install dependency:

```plain
go get github.com/russross/blackfriday
```

- Update [settings.yml](https://github.com/goby-lang/api.doc/blob/master/settings.yml) if necessary.
- Make sure the following command runs without any error:

```plain
go run *.go
```

This should generate (or overwrite) new docs in `/docs` directory. 

If you want to update the documentation, make sure you checkout to the desired branch for Goby project.

After that, commit and push them for update. GitHub will handle the rest.

## Documenting Goby Code

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

Each method comes with two keys, `Fn` and `Name`. The document of a method comes right above `Name`. For example:

```go
var builtinIntegerMethods = []*BuiltInMethod{
  {
    // Returns the sum of self and another Integer.
    //
    // ```
    // 1 + 1
    // ```
    Name: "+",
    Fn: func(receiver Object) builtinMethodBody {
      // ...
    },
  },
}
```

Remember to leave one space after `//`.

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
