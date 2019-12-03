
# Slashdiablo launcher

[![Go Report Card](https://goreportcard.com/badge/github.com/nokka/slashdiablo-launcher)](https://goreportcard.com/report/github.com/nokka/slashdiablo-launcher)

[![GoDoc](https://godoc.org/github.com/nokka/slashdiablo-launcher?status.svg)](https://godoc.org/github.com/nokka/slashdiablo-launcher)

  

## About
Slashdiablo launcher is a cross platform game launcher for Diablo II and specfically the [Slashdiablo]([https://old.reddit.com/r/slashdiablo/](https://old.reddit.com/r/slashdiablo/)) private server community. It was built to help new players install patches, updating gateways and help with other technical issues to lower the barrier of entry into the Slashdiablo community, while also assisting more experienced players with more advanced settings such as HD mods and launching multiple boxes.


  
## Features

- [x] Patching Diablo II up to 1.13c from previous game versions
- [x] Applying Slashdiablo patches automatically
- [x] Allows for multiple installs of Diablo II with different settings (such as Maphack & HD)
- [x] Automatically installs and updates Maphack & HD mod
- [x] Launch multiple Diablo II boxes from multiple installs
- [x] Help with OS specific configuration such as DEP issues
- [x] View ingame top ladder
- [ ] Patch Diablo II from 1.14 down to 1.13c


### Full OS support
- [x] Windows
- [ ] OSX (missing some D2 specific features)
- [ ] Linux (missing some D2 specific features)

## Installation

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
