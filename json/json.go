package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Post     Post     `json:"post"`
	Category Category `json:"category"`
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Category struct {
	PostID       int      `json:"post_id"`
	CategoryName []string `json:"name_cat"`
}

func main() {
	//jso, err := os.ReadFile("./json/asd.json")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(err, jso)

	file, _ := os.Open("./json/asd.json")
	defer file.Close()

	var res Response
	json.NewDecoder(file).Decode(&res)
	fmt.Println(res)
}

func CreatePost(post Post) (int, error) {
	return 1, nil
}

//func InsertCategory(category Category) {
//	for _, cat := range category.CategoryName {
//		db.Exec(`insert into post_category(post_id, catgory_name) values ($1, $2)`, category.PostID, cat)
//	}
//}
