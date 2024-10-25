package pkg

import (
    //"net/http"
)

// BUISNESS MODELS
// viewProject, editProjectPage, saveProject

func getPageData(title string) Page {
    return Page {
        PageTitle: title,
        Posts: getAllPosts(),
    }
}

func getHomeData() HomePage {
    return HomePage {
        Page: getPageData("Skye Chat"),
    }
}

