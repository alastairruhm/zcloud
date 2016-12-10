[![Build Status](https://travis-ci.org/alastairruhm/zcloud.svg?branch=master)](https://travis-ci.org/alastairruhm/zcloud) 
[![Go Report Card](https://goreportcard.com/badge/github.com/alastairruhm/zcloud)](https://goreportcard.com/report/github.com/alastairruhm/zcloud) 
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/alastairruhm/zcloud/blob/master/LICENSE) 
 
# zcloud

command line application for openstack client based on  gophercloud SDK for OpenStack

## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/alastairruhm/zcloud
```

## Usage

From the commandline, `zcloud` provides the functions as follows.By default, it prints its output to `stdout`.

### overview

```sh
$ zcloud --help
zcloud is a CLI app for openstack client.
This application is created for speeding up the daily operation on cloud servers on openstack

Usage:
  zcloud [command]

Available Commands:
  server      operation about server
  version     show zcloud version

Flags:
      --host string       openstack auth service host
      --password string   openstack user's name
      --username string   openstack user's name
  -v, --verbose           verbose output

Use "zcloud [command] --help" for more information about a command.
```

### server 

```sh
$ zcloud help server
echo is for echoing anything back.
    Echo echo’s.

Usage:
  zcloud server [command]

Available Commands:
  create      Echo anything to the screen
  delete      delete a specific server
  list        Echo anything to the screen

Flags:
  -n, --name string      server name
  -p, --project string   project name

Global Flags:
      --host string       openstack auth service host
      --password string   openstack user's name
      --username string   openstack user's name
  -v, --verbose           verbose output

Use "zcloud server [command] --help" for more information about a command.

```

### version

```sh
$ zcloud version
zcloud client tool version 0.0.1
```

## License

`zcloud` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).