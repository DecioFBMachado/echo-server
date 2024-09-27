# Simple HTTP/HTTPS Server with Custom Response

This Go application is a simple HTTP/HTTPS server that allows you to specify various configuration options such as port, status code, response body, and SSL certificates. It logs incoming requests and can serve a static response body from a file or from the request body itself.

## Features

- Serve HTTP or HTTPS requests.
- Set custom status code for responses.
- Serve a response body from a file or echo the request body.
- Log incoming request details including headers and body.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository or download the source code.
2. Make sure Go is installed on your system.
3. Build the binary

## Options

- **`-port`**: Defines the port to be used by the server (default is `8080`).
- **`-c`**: Path to the SSL certificate file. If both `-c` (certificate) and `-key` (key file) are provided, the server will start in HTTPS mode.
- **`-key`**: Path to the SSL key file (used with `-c` for HTTPS).
- **`-status`**: Set the HTTP status code for the response (default is `200 OK`).
- **`-rf`**: Path to the file containing the response body. If specified, the content of this file will be sent as the HTTP response body.
- **`-h`**: Show available options.

