# FileDAG SDK

[![](https://img.shields.io/github/go-mod/go-version/filedrive-team/go-filedag-sdk)]()
[![](https://goreportcard.com/badge/github.com/filedrive-team/go-filedag-sdk)](https://goreportcard.com/report/github.com/filedrive-team/go-filedag-sdk)
[![](https://img.shields.io/github/license/filedrive-team/go-filedag-sdk)](https://github.com/filedrive-team/go-filedag-sdk/blob/main/LICENSE)

Official GO SDK for [FileDAG](https://filedag.cloud)

## Overview

The FileDAG GO SDK provides the quickest / easiest path for interacting with the [FileDAG API](https://docs.filedag.cloud).

## Installation

Use go get.

	go get github.com/filedrive-team/go-filedag-sdk

Then import the sdk package into your own code.

	import "github.com/filedrive-team/go-filedag-sdk/client"

## Setup

To start, simply require the FileDAG SDK and set up an instance with your FileDAG API Keys. Don't know what your keys are? Check out your [API Keys Page](https://filedag.cloud/cn/dashboard/API%E7%A7%98%E9%92%A5).

```go
cli := client.NewWithJwtToken("https://api.filedag.cloud", "YOUR_JWT")
```
or

```go
cli := client.NewWithKeySecret("https://api.filedag.cloud", "YOUR_API_KEY", "YOUR_API_SECRET")
```

## Usage

Once you've set up your instance, using the FileDAG SDK is easy. Simply call your desired function and handle the results.

```go
resp, err := cli.PinnedDataTotal()
```

## License

Distributed under MIT License, please see license file within the code for more details.