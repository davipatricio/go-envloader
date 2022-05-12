# go-envloader

This project is a simple [Go](https://golang.org) package that will load environment variables from the `.env` file.

## Installation

**As a library**: `go get github.com/davipatricio/go-envloader`

## Example

```go
package main

import (
    "fmt"
    "github.com/davipatricio/go-envloader"
)

func main() {
    // Load environment variables from the `.env` file and set them to the `os.Environ` variable.
    err := go-envloader.Load()
    if err != nil {
        panic(err)
    }

    youtubeApiToken := os.Getenv("YOUTUBE_API")
    production := os.Getenv("PRODUCTION")

    fmt.Println("Youtube API Token:", youtubeApiToken)
    fmt.Println("Production:", production)
}
```

## License
[MIT](https://github.com/davipatricio/go-envloader/blob/master/LICENSE)
