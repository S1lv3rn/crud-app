package pkg

import (
    "net/http"
)




// ROUTER
func Start() { 
    fs := http.FileServer(http.Dir("assets/"))
    r := http.NewServeMux()
    r.Handle("/static/", http.StripPrefix("/static/", fs))



    r.HandleFunc("/", getHome)

    
    
    http.ListenAndServe(":8080", r)
}