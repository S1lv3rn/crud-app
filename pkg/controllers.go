package pkg

// CONTROLLERS
/*
Controllers:
 - gets request
 - extracts post/get variables and vaildates
 - requests model data
 - put model data in view

*/

import (
    "net/http"
    "html/template"
    "strconv"
)

var funcMap = template.FuncMap{
    "formatDate": formatDate,
}


func getHome(w http.ResponseWriter, r *http.Request) {
    data := getHomeData()

    tmpl := template.Must(template.New("home.html").Funcs(funcMap).ParseFiles(
        "views/home.html",
        "views/common.html",
    ))
    tmpl.Execute(w, data)
}