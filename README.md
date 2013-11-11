# BTSync API

Golang client for the Bittorrent Sync API.

## Example

```go
package main

import (
  btsync "btsync-api"
  "fmt"
  "log"
)

func main() {
  // Create an instance of the client. Needs the configured
  // login, password, and port of the Sync API. The last 
  // argument is to enable debug logging.
  api := btsync.New("login", "password", 8080, true)

  // Get a list of Sync folders. All API methods return
  // a response and an error object.
  folders, err := api.GetFolders()
  if err != nil {
    log.Fatalf("Error! %s", err)
  }
  
  // Response objects map directly to the documented JSON
  // responses. However they are Golang style (caps/camelcase).
  for _, folder := range *folders {
    fmt.Printf("Sync folder %s has %d files\n", folder.Dir, folder.Files)
  }
  
  // Get Sync's current upload/download speed.
  speed, _ := api.GetSpeed()
  fmt.Printf("Speed: upload=%d, download=%d", speed.Upload, speed.Download)
}

```

## Documentation

http://godoc.org/github.com/vole/btsync-api
