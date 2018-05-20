# Synology DSM API Client
Client library for calling the Synology DSM APIs. Synology DSM (DiskStation Manager) is the Operating System: <https://www.synology.com/en-us/dsm/6.1>

## Getting Started
Instructions for using the library.

### Getting the library
```
go get github.com/garrettdieckmann/synologyapi
```
### Examples
#### Get 1 Minute Load
```
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/garrettdieckmann/synologyapi"
)

func main() {
	var synas synologyapi.SynologyConnection
	if err := sapi.NewConnection("192.168.x.x", "5000", "account", "password"); err != nil {
		log.Fatal(err)
	}
	sysresp, err := synas.GetSystemInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("1 minute load: %v.", sysresp.CPU.OneMinLoad))
}
```

## Endpoints implemented
| Function Name | Synology API | Notes |
| ------------- | ------------ | ----- |
| getSIDToken | SYNO.API.Auth | Not an exported function |
| GetSystemInfo | SYNO.Core.System.Utilization | |
| GetShareInfo | SYNO.Core.Share | |
| GetStorageInfo | SYNO.Storage.CGI.Storage | |

## Acknowledgements
* Thomas Theunenh's blog post was incredibly helpful for discovering the Synology APIs behind the web-based DSM http://www.thomastheunen.eu/2015/06/the-synology-api-not-much-documentation.html