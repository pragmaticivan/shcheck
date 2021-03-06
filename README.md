# shcheck

A tool to easily bind shellcheck and shfmt to your CI, so your
shellscripts can be checked.

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.
We appreciate your contribution. Please refer to our [contributing guidelines](CONTRIBUTING.md).

[![Release](https://img.shields.io/github/release/caarlos0/shcheck.svg?style=flat-square)](https://github.com/caarlos0/shcheck/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Travis](https://img.shields.io/travis/caarlos0/shcheck.svg?style=flat-square)](https://travis-ci.org/caarlos0/shcheck)
[![Coverage Status](https://img.shields.io/codecov/c/github/caarlos0/shcheck/master.svg?style=flat-square)](https://codecov.io/gh/caarlos0/shcheck)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/caarlos0/shcheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/caarlos0/shcheck?style=flat-square)](https://goreportcard.com/report/github.com/caarlos0/shcheck)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)


## Usage

Just add it to your travis.yml:

```yaml
script:
  - curl -sL https://git.io/shcheck | bash -s
```

You can also pass files and folders to be ignored, like this:

```yaml
script:
  - curl -sL https://git.io/shcheck | bash -s -f -- --ignore=somefile.sh --ignore='folder/**/*'
```
