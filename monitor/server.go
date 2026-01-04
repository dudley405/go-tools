package main

import (
	"io"
    "net/http"
    "encoding/json"
    "app/types"
    "app/logging"
    "go.uber.org/zap"
)



/* kubectl proxy --port=8080
 * http://127.0.0.1:8080/api/v1/namespaces/default/pods
 *for prod curl --cacert ~/.minikube/ca.crt --cert ~/.minikube/client.crt --key ~/.minikube/client.key https://`minikube ip`:8443/api/
*/

func main() {
    var logger = logging.NewLogger()

    resourceType := "services"
    body, err := requestApiResources(resourceType, logger)

    results := make([]string, 0)

    for _, item := range body.Items {
        results = append(results, item.Metadata.Name)
    }

    response := make(map[string]interface{})

    response[resourceType] = results

    json, err := json.Marshal(response)
        if err != nil {
            logger.Error("Failed marshalling json")
        }

    buffer := string(json)
    logger.Info(buffer)
}

func requestApiResources(resourceType string, logger *zap.Logger) (response types.KubeResponse, err error) {
     request := string("http://127.0.0.1:8080/api/v1/" + resourceType)
     resp, err := http.Get(request)

     if err != nil {
        logger.Error("Error fetching Kubernetes API",
            zap.String("request", request),
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