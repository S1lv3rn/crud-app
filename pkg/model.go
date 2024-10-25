package pkg

import (
    "net/http"
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

func getProjectData(id int, w http.ResponseWriter) (ProjectPage, error) {
    p, err := getProjectById(id)

    return ProjectPage {
        Page: getPageData(p.Title),
        Project: p,
    }, err
}