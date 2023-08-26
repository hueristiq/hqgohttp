# hqgohttp/status

`hqgohttp/status` lists all the HTTP status codes as constants.

Categories of the HTTP Status Codes:
* **Informational responses (100–199)**: These are informational responses indicating that the request was received and understood. It is issued on a provisional, intermediate basis while the request continues to be processed.
* **Successful responses (200–299)**: These indicate that the client's request was successfully received, understood, and accepted.
* **Redirection messages (300–399)**: These are used to indicate that the client must take some additional action in order to complete their request.
* **Client error responses (400–499)**: These are used to indicate that there was a problem with the request made by the client.
* **Server error responses (500–599)**: These indicate that the server encountered a problem and was unable to complete the request.

## Resource

* [Usage](#usage)
* [Contributing](#contributing)
* [Licensing](#licensing)

## Usage

To use this package in your Go application, you need to import it like any other Go package.

```go
package main

import (
    "fmt"
    "github.com/hueristiq/hqgohttp/status"
)

func main() {
    fmt.Println("Status OK:", status.OK)
    fmt.Println("Status NotFound:", status.NotFound)
}
```

## Contributing

[Issues](https://github.com/hueristiq/hqgohttp/issues) and [Pull Requests](https://github.com/hueristiq/hqgohttp/pulls) are welcome! **Check out the [contribution guidelines](https://github.com/hueristiq/hqgohttp/blob/master/CONTRIBUTING.md).**

## Licensing

This utility is distributed under the [MIT license](https://github.com/hueristiq/hqgohttp/blob/master/LICENSE).