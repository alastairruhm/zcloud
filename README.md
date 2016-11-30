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

```sh
Usage:
  zcloud [command]

Available Commands:
  server      Echo anything to the screen
  version     Echo anything to the screen

Flags:
      --host string       openstack auth service host
      --password string   openstack use's name
      --username string   openstack use's name
  -v, --verbose           verbose output

Use "zcloud [command] --help" for more information about a command.
```


## License

`zcloud` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).