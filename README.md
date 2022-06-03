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
    err := envloader.LoadEnv()
    if err != nil {
        panic(err)
    }

    youtubeApiToken := os.Getenv("YOUTUBE_API")
    production := os.Getenv("PRODUCTION")

    fmt.Println("Youtube API Token:", youtubeApiToken)
    fmt.Println("Production:", production)
}
```

## Related

[godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.)

[python-dotenv](https://pypi.org/project/python-dotenv/) - Read key-value pairs from a .env file and set them as environment variables

[dotenv](https://www.npmjs.com/package/dotenv) - Loads environment variables from .env file

## License
[MIT](https://github.com/davipatricio/go-envloader/blob/master/LICENSE)
