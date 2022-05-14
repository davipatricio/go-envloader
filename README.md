# go-envloader

This project is a simple [Go](https://golang.org) package that will load environment variables from the `.env` file.

## Installation

**As a library**: `go get github.com/denkylabs/go-envloader`

## Example

```go
package main

import (
    "fmt"
    "github.com/denkylabs/go-envloader"
)

func main() {
    // Load environment variables from the `.env` file and set them to the `os.Environ` variable.
    err := envloader.Load()
    if err != nil {
        panic(err)
    }

    youtubeApiToken := os.Getenv("YOUTUBE_API")
    production := os.Getenv("PRODUCTION")

    fmt.Println("Youtube API Token:", youtubeApiToken)
    fmt.Println("Production:", production)
}
```

## Similar projects

[godotenv](https://github.com/joho/godotenv)

[python-dotenv](https://pypi.org/project/python-dotenv/)

[dotenv](https://www.npmjs.com/package/dotenv)

## License
[MIT](https://github.com/davipatricio/go-envloader/blob/master/LICENSE)
