package clients

import (
    "net/http"
    "encoding/json"
    "app/types"
    "app/logging"
    "go.uber.org/zap"
    "io"
    "strings"
    "fmt"
)



    // A separate handler function with the required signature
func ApiResourceHandler(w http.ResponseWriter, r *http.Request) {
    var logger = logging.NewLogger()
    resourceType := strings.TrimPrefix(r.URL.Path, "/api/v1/resources/")

    body, err := requestApiResources(resourceType, logger)

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

func requestApiResources(resourceType string, logger *zap.Logger) (response types.KubeResponse, err error) {
     request := string("http://127.0.0.1:8080/api/v1/" + resourceType)
     resp, err := http.Get(request)

     if err != nil {
        logger.Error("Error fetching Kubernetes API",
            zap.String("request", request),
            zap.String("status", resp.Status),
            zap.Error(err))
     }


     body, err := io.ReadAll(resp.Body)
     if err != nil {
        logger.Error("Error reading Kubernetes API response",
            zap.String("request", request),
            zap.Error(err))
     }

    var data types.KubeResponse

    if err := json.Unmarshal(body, &data); err != nil {
        panic(err)
    }

    return data, err
}
