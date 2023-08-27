# hqgohttp/headers

`hqgohttp/headers` lists the HTTP header field names as constants.

Here is a brief overview of the categories of header fields included in this package:
* **Authentication:** Related to authentication and authorization.
* **Caching:** Related to caching policy.
* **Client hints:** Used for content adaptation.
* **Conditionals:** Used for conditional requests.
* **Connection management:** Related to connection management.
* **Content negotiation:** Used for content negotiation.
* **Controls:** Related to general controls.
* **CORS:** Related to Cross-Origin Resource Sharing.
* **Do Not Track:** Related to user tracking preference.
* **Downloads:** Related to content disposition.
* **Message body information:** Related to the message body.
* **Proxies:** Related to proxy servers.
* **Redirects:** Related to HTTP redirection.
* **Request context:** Related to the request context.
* **Response context:** Related to the response context.
* **Range requests:** Related to partial requests and responses.
* **Security:** Related to security.
* **Server-sent event:** Related to server-sent events.
* **Transfer coding:** Related to transfer coding.
* **WebSockets:** Related to the WebSocket protocol.
* **Other:** Do not fall into the above categories.

## Usage

To use this package in your Go application, you need to import it like any other Go package.

```go
package main

import (
    "fmt"
    "github.com/hueristiq/hqgohttp/headers"
)

func main() {
    fmt.Println("Authorization Header:", headers.Authorization)
    fmt.Println("Connection Header:", headers.Connection)
}
```

## Contributing

[Issues](https://github.com/hueristiq/hqgohttp/issues) and [Pull Requests](https://github.com/hueristiq/hqgohttp/pulls) are welcome! **Check out the [contribution guidelines](https://github.com/hueristiq/hqgohttp/blob/master/CONTRIBUTING.md).**

## Licensing

This utility is distributed under the [MIT license](https://github.com/hueristiq/hqgohttp/blob/master/LICENSE).