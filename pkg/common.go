package pkg

import (
    "fmt"
    "time"
    "net/http"
)



// COMMON STRUCTS

type Page struct {
    PageTitle string
    Posts []Post
}

type Post struct {
	Id int64
	Text string
	Created int64
	Updated int64
}

type HomePage struct {
    Page
}

type ProjectPage struct {
    Page
    Project
}



// COMMON FUNCTIONS

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkToDisplay(e error, w http.ResponseWriter) bool {
    if e != nil {
        fmt.Fprintf(w, e.Error())
    }
    return e != nil
}

func formatDate(epochTime int64) string {
    return time.Unix(epochTime, 0).Format("01/02/2006")
}