<div align="center">

# 👏 gopicker

*A command line [pkg.go.dev](https://pkg.go.dev) searcher and `go get` helper*

[![MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)](https://github.com/edifierx666/gopicker/blob/main/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/edifierx666/gopicker?style=flat-square)](https://github.com/edifierx666/gopicker/releases/latest)

</div>

## Features

- [x] Quickly search Go packages 
- [x] Easily get package that you selected

<img width="1078" alt="pic" src="https://user-images.githubusercontent.com/39356752/232316344-60ad8492-1f7f-4dd3-b4d5-2c96e7e736c5.png">


## Usage

```
Usage:
  gopicker [OPTIONS] QUERY...

Application Options:
  -l, --limit=   Number of search result items limit (default: 20)

Help Options:
  -h, --help     Show this help message
```


> **NOTE**:
> To see examples of keywords to search for, check [search-help](https://pkg.go.dev/search-help) on pkg.go.dev.

## Installation

You can download the executable binaries from the latest page.

> [![Latest Release](https://img.shields.io/github/v/release/edifierx666/gopicker?style=flat-square)](https://github.com/edifierx666/gopicker/releases/latest)

To build from source, clone or download this repository then run `go install`, or run below:

```sh
go install github.com/edifierx666/gopicker@latest
```

## License

[MIT](./LICENSE)
