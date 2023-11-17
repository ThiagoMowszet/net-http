package server

import (
    "net/http"
    "fmt"
)

func initRoutes(){
    http.HandleFunc("/", index)

    http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {

        case http.MethodGet:
            getCountries(w,r)

        case http.MethodPost:
            addCountries(w, r)

        default:
            if r.Method != http.MethodGet {
                w.WriteHeader(http.StatusMethodNotAllowed)
                fmt.Fprintf(w, "Method not allowed")
            } else {
                fmt.Fprintf(w, "Hello There %s", "visitor")
            }
        }
    })

}
