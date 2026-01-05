package main

import (

    "app/clients"

    "net/http"
    "fmt"
)



/* kubectl proxy --port=8080
 * http://127.0.0.1:8080/api/v1/namespaces/default/pods
 *for prod curl --cacert ~/.minikube/ca.crt --cert ~/.minikube/client.crt --key ~/.minikube/client.key https://`minikube ip`:8443/api/
*/

func main() {
    startServer()
}

func startServer() {

    http.HandleFunc("/api/v1/resources/", clients.ApiResourceHandler)

    fmt.Println("Server starting on port 8081...")

    if err := http.ListenAndServe(":8081", nil); err != nil {
    	fmt.Printf("Server failed: %v\n", err)
    }
}

