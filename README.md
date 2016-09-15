slacker
=========

[![Build Status](https://img.shields.io/circleci/project/amylum/slacker.svg)](https://circleci.com/gh/amylum/slacker)
[![GitHub release](https://img.shields.io/github/release/akerl/slacker.svg)](https://github.com/akerl/slacker/releases)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Simple binary for converting a MAC address to an EUI-64 address. I wrote this for use in finding the SLAAC address of Docker containers.

## Usage

```
./slacker 02:42:ac:10:00:02
```

## Installation

```
git clone git://github.com/akerl/slacker
cd slacker
go build
# Optionally, move ./slacker into your $PATH somewhere, like /usr/local/bin
```

## License

slacker is released under the MIT License. See the bundled LICENSE file for details.

