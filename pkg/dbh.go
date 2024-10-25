package pkg

import (
    "errors"

	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DATABASE MODELS




func getAllPosts() []Post {

	db, err := sql.Open("sqlite3", "data/crud.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	result, err := db.Query("SELECT * FROM posts ORDER BY createdOn")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()
	var posts []Post

	
	for result.Next() {
		var temp Post
		
		if err := result.Scan(&temp.Id, &temp.Text, &tempBook.Created, &tempBook.Updated); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, temp)
	}

    return posts
}



func addPost(text string) bool {
	db, err := sql.Open("sqlite3", "data/crud.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	stmt, err := db.Prepare("INSERT INTO posts(text, createdOn) VALUES(?, ?)")
	_, err := stmt.Query(text, now.UnixMilli())

	if err != nil {
		log.Fatal(err)
	}


    return true
}

func deletePost(id int64) {

}


func updatePostText(id int64, newText string) {

}


func getPostById(id int) (Post, error) {
    
}