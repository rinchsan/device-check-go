# device-check-go

![golang](https://img.shields.io/badge/golang-1.11-blue.svg?style=flat)

iOS DeviceCheck SDK for Go - query and modify the per-device bits

## Installation

```bash
go get github.com/snowman-mh/device-check-go
```

## Get started

### Query two bits

```go
import "github.com/snowman-mh/device-check-go"

cred := devicecheck.NewCredentialFile("/path/to/private/key/file")
cfg := devicecheck.NewConfig("ISSUER", "KEY_ID", devicecheck.Development)
client := devicecheck.New(cred, cfg)

result := devicecheck.QueryTwoBitsResult{}
err := client.QueryTwoBits("DEVICE_TOKEN_FROM_CLIENT", &result)
```

### Update two bits

```go
import "github.com/snowman-mh/device-check-go"

cred := devicecheck.NewCredentialFile("/path/to/private/key/file")
cfg := devicecheck.NewConfig("ISSUER", "KEY_ID", devicecheck.Development)
client := devicecheck.New(cred, cfg)

err := client.UpdateTwoBits("DEVICE_TOKEN_FROM_CLIENT", true, true)
```

### Validate device token

```go
import "github.com/snowman-mh/device-check-go"

cred := devicecheck.NewCredentialFile("/path/to/private/key/file")
cfg := devicecheck.NewConfig("ISSUER", "KEY_ID", devicecheck.Development)
client := devicecheck.New(cred, cfg)

err := client.ValidateDeviceToken("DEVICE_TOKEN_FROM_CLIENT")
```

## Apple documentation

- [iOS DeviceCheck API for Swift](https://developer.apple.com/documentation/devicecheck)
- [HTTP commands to query and modify the per-device bits](https://developer.apple.com/documentation/devicecheck/accessing_and_modifying_per-device_data)
