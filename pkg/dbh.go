package pkg

import (
    //"errors"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DATABASE MODELS




func getAllPosts() []Post {

	db, err := sql.Open("sqlite3", "data/crud.db")
	if err != nil {
		log.Print(err)
		return nil
	}
	log.Print("Connected to db")

	defer db.Close()


	rows, err := db.Query("SELECT * FROM main.posts ORDER BY createdOn")
	if err != nil {
		log.Print(err)
		return nil
	}

	log.Print("Got the results")

	if (!rows.Next()) { 
		log.Print("Empty List!")
		return nil
	}

	defer rows.Close()
	var posts []Post

	
	for rows.Next() {
		var temp Post
		
		if err := rows.Scan(&temp.Id, &temp.Text, &temp.Created, &temp.Updated); err != nil {
			log.Print(err)
			return nil
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
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Query(text, time.Now().UnixMilli())

	if err != nil {
		log.Fatal(err)
	}


    return result != nil
}

func deletePost(id int64) {

}


func updatePostText(id int64, newText string) {

}


func getPostById(id int) {
    
}