
# Slashdiablo launcher

[![Go Report Card](https://goreportcard.com/badge/github.com/nokka/slashdiablo-launcher)](https://goreportcard.com/report/github.com/nokka/slashdiablo-launcher)
[![GoDoc](https://godoc.org/github.com/nokka/slashdiablo-launcher?status.svg)](https://godoc.org/github.com/nokka/slashdiablo-launcher)

![launcher screenshot](/docs/launcher.png)

## About the project

Slashdiablo launcher is a cross platform game launcher for Diablo II and specifically the [Slashdiablo](https://old.reddit.com/r/slashdiablo/) community. It was built to help new players install patches, updating gateways and help with other technical issues to lower the barrier of entry into the Slashdiablo community, while also assisting more experienced players with more advanced settings such as HD mods and launching multiple boxes.
  
## Features

- [x] Patching Diablo II up to 1.13c from previous game versions
- [x] Applying Slashdiablo patches automatically
- [x] Allows for multiple installs of Diablo II with different settings (such as Maphack & HD)
- [x] Automatically installs and updates Maphack & HD mod
- [x] Launch multiple Diablo II boxes from multiple installs
- [x] Help with OS specific configuration such as DEP issues
- [x] View ingame top ladder
- [ ] Patch Diablo II from 1.14+ down to 1.13c


### Full OS support
- [x] Windows
- [ ] OSX (missing some D2 specific features)
- [ ] Linux (missing some D2 specific features)


## Development

### Go
Install Go 1.12 or higher by following  [install instructions](http://golang.org/doc/install.html) for your OS.

### Qt bindings for Go
Before you can build you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

###  Install Qt5

#### OSX
On OSX using brew is by far the most simple way of installing Qt5.
```bash
$ brew install qt
```

#### Windows
Use the [installer](https://download.qt.io/official_releases/qt/5.13/5.13.0/qt-opensource-windows-x86-5.13.0.exe) provided by Qt (Make sure you install the MinGW build of Qt).

#### Building Slashdiablo launcher

```bash
# Get binding source
$ go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...

# Download the repository with dependencies
$ go get -d -u -v github.com/nokka/slashdiablo-launcher

# Build the launcher
$ cd $(go env GOPATH)/src/github.com/nokka/slashdiablo-launcher
$ qtdeploy build

# Start launcher (different depending on OS)
$ ./deploy/darwin/slashdiablo-launcher.app/Contents/MacOS/slashdiablo-launcher
```

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
