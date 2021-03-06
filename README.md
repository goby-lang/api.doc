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

Go to [Goby wiki](https://github.com/goby-lang/goby/wiki/Documenting-Goby-Code).
