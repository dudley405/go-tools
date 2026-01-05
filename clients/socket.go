package main

import (
	"fmt"
    "context"
	"net"
	"net/http"
	"io"
)

// kubectl proxy --unix-socket=./proxy.sock --accept-hosts='^.*$'
var (
    socketPath = "/home/dudley/proxy.sock"
    // Creating a new HTTP client that is configured to make HTTP requests over a Unix domain socket.
    httpClient = http.Client{
        Transport: &http.Transport{
            // Set the DialContext field to a function that creates
            // a new network connection to a Unix domain socket
            DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
                return net.Dial("unix", socketPath)
            },
        },
    }
)


func main () {
       resp, err := httpClient.Get("http://unix/api/v1/pods")
       if err != nil {
            panic(err)
       }
       defer resp.Body.Close()

       b, err := io.ReadAll(resp.Body)
          if err != nil {
              fmt.Printf("%s", err)
          }

       fmt.Printf("%s", resp)
       fmt.Printf("%s", b)


}