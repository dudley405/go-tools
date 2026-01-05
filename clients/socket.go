package clients

import (
	"fmt"
    "context"
	"net"
	"net/http"
	"io"
	"app/types"
	"app/logging"
	"encoding/json"
	"go.uber.org/zap"
	"strings"
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


func SocketResourceHandler(w http.ResponseWriter, r *http.Request) {
    var logger = logging.NewLogger()
    resourceType := strings.TrimPrefix(r.URL.Path, "/api/v1/resources/")

    body, err := requestSocketResource(resourceType, logger)

    // Hold list of results
    results := make([]string, 0)

    for _, item := range body.Items {
        results = append(results, item.Metadata.Name)
    }

    /* Create object json response - I.e if 'services' is resource type:
     *
     *  {"services": [results_here]}
     */
    response := make(map[string]interface{})
    response[resourceType] = results

    json, err := json.Marshal(response)
    if err != nil {
        logger.Error("Failed marshalling json")
    }

    buffer := string(json)
    logger.Info(buffer)
    fmt.Fprintf(w, buffer)
 }


func requestSocketResource(resourceType string, logger *zap.Logger) (response types.KubeResponse, err error) {
       resp, err := httpClient.Get("http://unix/api/v1/" + resourceType)
       if err != nil {
            panic(err)
       }
       defer resp.Body.Close()

       body, err := io.ReadAll(resp.Body)
          if err != nil {
              fmt.Printf("%s", err)
          }


       var data types.KubeResponse

       if err := json.Unmarshal(body, &data); err != nil {
           panic(err)
       }

       return data, err
}