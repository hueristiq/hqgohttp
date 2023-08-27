# hqgohttp/methods

`hqgohttp/methods` lists the HTTP methods as constants.

HTTP Methods:
* **GET**: The GET method is used to request data from a specified resource.
* **HEAD**: The HEAD method is almost identical to GET, but without the response body.
* **POST**: The POST method is used to send data to a server to create/update a resource.
* **PUT**: The PUT method is used to update a current resource with new data.
* **PATCH**: The PATCH method is used to apply partial modifications to a resource.
* **DELETE**: The DELETE method is used to delete the specified resource.
* **CONNECT**: The CONNECT method establishes a tunnel to the server identified by the target resource.
* **OPTIONS**: The OPTIONS method is used to describe the communication options for the target resource.
* **TRACE**: The TRACE method performs a message loop-back test along the path to the target resource.

## Usage

To use this package in your Go application, you need to import it like any other Go package.

```go
package main

import (
    "fmt"
    "github.com/hueristiq/hqgohttp/methods"
)

func main() {
    fmt.Println("GET Method:", methods.Get)
    fmt.Println("POST Method:", methods.Post)
}
```

## Contributing

[Issues](https://github.com/hueristiq/hqgohttp/issues) and [Pull Requests](https://github.com/hueristiq/hqgohttp/pulls) are welcome! **Check out the [contribution guidelines](https://github.com/hueristiq/hqgohttp/blob/master/CONTRIBUTING.md).**

## Licensing

This utility is distributed under the [MIT license](https://github.com/hueristiq/hqgohttp/blob/master/LICENSE).