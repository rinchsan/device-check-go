# device-check-go

[![GitHub Actions Status](https://github.com/rinchsan/device-check-go/workflows/build/badge.svg)](https://github.com/rinchsan/device-check-go/actions)
![GitHub release](https://img.shields.io/github/release/rinchsan/device-check-go.svg?colorB=7E7E7E)
[![Go Reference](https://pkg.go.dev/badge/github.com/rinchsan/device-check-go.svg)](https://pkg.go.dev/github.com/rinchsan/device-check-go)
[![codecov.io](https://codecov.io/github/rinchsan/device-check-go/coverage.svg?branch=master)](https://codecov.io/github/rinchsan/device-check-go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rinchsan/device-check-go)](https://goreportcard.com/report/github.com/rinchsan/device-check-go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

iOS DeviceCheck SDK for Go - query and modify the per-device bits

## Installation

```bash
go get github.com/rinchsan/device-check-go
```

## Get started

### Initialize SDK

```go
import "github.com/rinchsan/device-check-go"

cred := devicecheck.NewCredentialFile("/path/to/private/key/file") // You can create credential also from raw string/bytes
cfg := devicecheck.NewConfig("ISSUER", "KEY_ID", devicecheck.Development)
client := devicecheck.New(cred, cfg)
````

### Use DeviceCheck API

#### Query two bits

```go
var result devicecheck.QueryTwoBitsResult
if err := client.QueryTwoBits("DEVICE_TOKEN_FROM_CLIENT", &result); err != nil {
	// error handling
	// Note that SDK returns ErrBitStateNotFound error if no bits found
}
```

#### Update two bits

```go
if err := client.UpdateTwoBits("DEVICE_TOKEN_FROM_CLIENT", true, true); err != nil {
	// error handling
}
```

#### Validate device token

```go
if err := client.ValidateDeviceToken("DEVICE_TOKEN_FROM_CLIENT"); err != nil {
	// error handling
}
```

## Apple documentation

- [iOS DeviceCheck API for Swift](https://developer.apple.com/documentation/devicecheck)
- [HTTP commands to query and modify the per-device bits](https://developer.apple.com/documentation/devicecheck/accessing_and_modifying_per-device_data)
