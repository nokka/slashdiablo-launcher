# Slashdiablo launcher
[![Go Report Card](https://goreportcard.com/badge/github.com/nokka/slashdiablo-launcher)](https://goreportcard.com/report/github.com/nokka/slashdiablo-launcher)
[![GoDoc](https://godoc.org/github.com/nokka/slashdiablo-launcher?status.svg)](https://godoc.org/github.com/nokka/slashdiablo-launcher)

Work in progress, gets updated frequently.

## Deploying
Deploying to a target can be done from any host OS if there's a docker image available,
otherwise the target OS and the host must be the same.

### Windows

```bash
$ docker pull therecipe/qt:windows_64_static
$ qtdeploy -docker build windows_64_static
```

### MacOS (from MacOS only)

```bash
$ qtdeploy build darwin github.com/nokka/slashdiablo-launcher
```
