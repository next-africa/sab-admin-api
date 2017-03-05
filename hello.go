package hello

import (
    "fmt"
    "net/http"
    "os"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/messenger", handleMessengerHook)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func handleMessengerHook(responseWriter http.ResponseWriter, request *http.Request) {
    query := request.URL.Query()

    if query.Get("hub.mode") == "subscribe" && query.Get("hub.verify_token") == os.Getenv("HUB_VERIFY_TOKEN") {
        responseWriter.WriteHeader(http.StatusOK)
        fmt.Fprint(responseWriter, query.Get("hub.challenge"))
    } else {
        fmt.Print(responseWriter, query)
        responseWriter.WriteHeader(http.StatusForbidden)
    }
    return
}
