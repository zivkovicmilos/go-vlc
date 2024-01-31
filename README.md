## Overview

`go-vlc` is a Go library for interacting with running VLC instances over the VLC HTTP API. This library acts as a client
implementation, offering seamless integration with officially documented VLC API endpoints.

The library encompasses a range of endpoints, as outlined in the sections below:

- For insights into the Official VLC Lua implementation, visit
  the [official VLC Lua implementation](https://github.com/videolan/vlc/blob/f7bb59d9f51cc10b25ff86d34a3eff744e60c46e/share/lua/http/requests/README.txt)
- To learn more about VLC HTTP Requests, refer to
  the [VLC HTTP Requests documentation](https://wiki.videolan.org/VLC_HTTP_requests/)

## Installation

Integrating go-vlc into your project is straightforward, execute the following command:

```shell
go get -u github.com/zivkovicmilos/go-vlc
```

This command ensures that you have the latest version of go-vlc in your project.

## Running the VLC HTTP server

To run the VLC HTTP server, using the CLI interface, execute the following command:

```shell
vlc --extraintf http --http-port 8080 --http-password 1234
```

Remember, the password used here is crucial for initializing and authenticating the VLC client within the `go-vlc`
library.

For GUI-based server initialization, please
consult [the official VLC documentation](https://wiki.videolan.org/Control_VLC_via_a_browser/).

## Example usage

Below is a sample Go program demonstrating how to utilize the `go-vlc` library:

```go
package main

import (
	"fmt"

	"github.com/zivkovicmilos/go-vlc/client/http"
	"github.com/zivkovicmilos/go-vlc"
)

func main() {
	// Create the VLC credentials
	baseURL := "http://127.0.0.1:8080" // server IP, and the http-port value

	auth := http.RequestAuth{
		Username: "",     // the username; default empty
		Password: "1234", // the http-password value
	}

	// Create the VLC client
	client := http.NewClient(baseURL, auth)

	// Create the VLC instance
	v := vlc.NewVLC(client)

	// Call any VLC method
	status, err := v.GetStatus()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", status)
}
```

