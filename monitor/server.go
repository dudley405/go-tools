package main

import (
	"io"
    "log"
    "net/http"
    "encoding/json"
    "app/types"
    "app/logging"
)



/* kubectl proxy --port=8080
 * http://127.0.0.1:8080/api/v1/namespaces/default/pods
 *for prod curl --cacert ~/.minikube/ca.crt --cert ~/.minikube/client.crt --key ~/.minikube/client.key https://`minikube ip`:8443/api/
*/

func main() {
    var logger = logging.NewLogger()

    resource_type := "pods"
    resp, err := http.Get("http://127.0.0.1:8080/api/v1/" + resource_type)
    if err != nil {
        log.Fatalln(err)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }


    var data types.KubeResponse

    if err := json.Unmarshal(body, &data); err != nil {
        panic(err)
    }

    results := make([]string, 0)

    for _, item := range data.Items {
        results = append(results, item.Metadata.Name)
    }


    response := make(map[string]interface{})

    response[resource_type] = results

    json, err := json.Marshal(response)
        if err != nil {
            logger.Error("Failed marshalling json")
        }

    buffer := string(json)
    logger.Info(buffer)
}